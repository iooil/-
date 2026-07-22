package controller

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"campus-course-server/config"
	"campus-course-server/model"
	"campus-course-server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CourseListItem struct {
	ID            uint64  `json:"id"`
	CourseNo      string  `json:"course_no"`
	CourseName    string  `json:"course_name"`
	TeacherID     uint64  `json:"teacher_id"`
	TeacherName   string  `json:"teacher_name"`
	TeacherTitle  string  `json:"teacher_title"`
	College       string  `json:"college"`
	Credit        float64 `json:"credit"`
	CourseType    string  `json:"course_type"`
	Description   string  `json:"description"`
	Capacity      int     `json:"capacity"`
	SelectedCount int     `json:"selected_count"`
	Remaining     int     `json:"remaining"`
	AverageScore  float64 `json:"average_score"`
	ReviewCount   int     `json:"review_count"`
	Status        int8    `json:"status"`
	IsSelected    bool    `json:"is_selected"`
}

func GetCourseList(context *gin.Context) {
	keyword := strings.TrimSpace(context.Query("keyword"))
	college := strings.TrimSpace(context.Query("college"))
	courseType := strings.TrimSpace(context.Query("course_type"))

	userIDValue, exists := context.Get("userID")
	if !exists {
		utils.Fail(context, http.StatusUnauthorized, "请先登录")
		return
	}

	userID, ok := userIDValue.(uint64)
	if !ok {
		utils.Fail(context, http.StatusInternalServerError, "用户身份解析失败")
		return
	}

	query := config.DB.
		Table("courses AS c").
		Select(`
			c.id,
			c.course_no,
			c.course_name,
			c.teacher_id,
			t.name AS teacher_name,
			t.title AS teacher_title,
			c.college,
			c.credit,
			c.course_type,
			c.description,
			c.capacity,
			c.selected_count,
			c.average_score,
			c.review_count,
			c.status
		`).
		Joins("LEFT JOIN teachers AS t ON t.id = c.teacher_id").
		Where("c.status = ?", 1)

	if keyword != "" {
		searchKeyword := "%" + keyword + "%"
		query = query.Where(
			"(c.course_name LIKE ? OR c.course_no LIKE ? OR t.name LIKE ?)",
			searchKeyword,
			searchKeyword,
			searchKeyword,
		)
	}

	if college != "" {
		query = query.Where("c.college = ?", college)
	}

	if courseType != "" {
		query = query.Where("c.course_type = ?", courseType)
	}

	var courses []CourseListItem

	if err := query.Order("c.id ASC").Scan(&courses).Error; err != nil {
		utils.Fail(context, http.StatusInternalServerError, "课程查询失败")
		return
	}

	var selectedCourseIDs []uint64

	config.DB.
		Table("course_selections").
		Where("user_id = ? AND status = ?", userID, "selected").
		Pluck("course_id", &selectedCourseIDs)

	selectedMap := make(map[uint64]bool)

	for _, courseID := range selectedCourseIDs {
		selectedMap[courseID] = true
	}

	for index := range courses {
		courses[index].Remaining =
			courses[index].Capacity - courses[index].SelectedCount

		if courses[index].Remaining < 0 {
			courses[index].Remaining = 0
		}

		courses[index].IsSelected = selectedMap[courses[index].ID]
	}

	utils.Success(context, "获取课程列表成功", courses)
}

func GetCourseFilters(context *gin.Context) {
	var colleges []string
	var courseTypes []string

	config.DB.
		Table("courses").
		Where("status = ?", 1).
		Distinct().
		Order("college ASC").
		Pluck("college", &colleges)

	config.DB.
		Table("courses").
		Where("status = ?", 1).
		Distinct().
		Order("course_type ASC").
		Pluck("course_type", &courseTypes)

	utils.Success(context, "获取课程筛选条件成功", gin.H{
		"colleges":     colleges,
		"course_types": courseTypes,
	})
}

func SelectCourse(context *gin.Context) {
	userIDValue, exists := context.Get("userID")
	if !exists {
		utils.Fail(context, http.StatusUnauthorized, "请先登录")
		return
	}

	userID, ok := userIDValue.(uint64)
	if !ok {
		utils.Fail(context, http.StatusInternalServerError, "用户身份解析失败")
		return
	}

	courseID, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil || courseID == 0 {
		utils.Fail(context, http.StatusBadRequest, "课程ID错误")
		return
	}

	err = config.DB.Transaction(func(transaction *gorm.DB) error {
		var course model.Course

		result := transaction.First(&course, courseID)

		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return errors.New("课程不存在")
			}

			return errors.New("课程查询失败")
		}

		if course.Status != 1 {
			return errors.New("该课程暂未开放选课")
		}

		var selection model.CourseSelection

		queryResult := transaction.
			Where("user_id = ? AND course_id = ?", userID, courseID).
			First(&selection)

		if queryResult.Error == nil && selection.Status == "selected" {
			return errors.New("你已经选择了该课程")
		}

		if course.SelectedCount >= course.Capacity {
			return errors.New("该课程人数已满")
		}

		now := time.Now()

		if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
			selection = model.CourseSelection{
				UserID:     userID,
				CourseID:   courseID,
				Status:     "selected",
				SelectedAt: now,
			}

			if err := transaction.Create(&selection).Error; err != nil {
				return errors.New("创建选课记录失败")
			}
		} else {
			if err := transaction.Model(&selection).Updates(map[string]interface{}{
				"status":      "selected",
				"selected_at": now,
				"dropped_at":  nil,
			}).Error; err != nil {
				return errors.New("恢复选课记录失败")
			}
		}

		if err := transaction.
			Model(&model.Course{}).
			Where("id = ? AND selected_count < capacity", courseID).
			UpdateColumn("selected_count", gorm.Expr("selected_count + 1")).
			Error; err != nil {
			return errors.New("更新课程人数失败")
		}

		return nil
	})

	if err != nil {
		utils.Fail(context, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(context, "选课成功", nil)
}
