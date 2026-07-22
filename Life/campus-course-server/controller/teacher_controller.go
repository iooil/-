package controller

import (
	"net/http"
	"strings"

	"campus-course-server/config"
	"campus-course-server/utils"

	"github.com/gin-gonic/gin"
)

type TeacherListItem struct {
	ID           uint64  `json:"id"`
	TeacherNo    string  `json:"teacher_no"`
	Name         string  `json:"name"`
	College      string  `json:"college"`
	Title        string  `json:"title"`
	Introduction string  `json:"introduction"`
	AverageScore float64 `json:"average_score"`
	ReviewCount  int     `json:"review_count"`
	CourseCount  int64   `json:"course_count"`
}

func GetTeacherList(context *gin.Context) {
	keyword := strings.TrimSpace(context.Query("keyword"))
	college := strings.TrimSpace(context.Query("college"))

	query := config.DB.
		Table("teachers AS t").
		Select(`
			t.id,
			t.teacher_no,
			t.name,
			t.college,
			t.title,
			t.introduction,
			t.average_score,
			t.review_count,
			COUNT(c.id) AS course_count
		`).
		Joins("LEFT JOIN courses AS c ON c.teacher_id = t.id").
		Group(`
			t.id,
			t.teacher_no,
			t.name,
			t.college,
			t.title,
			t.introduction,
			t.average_score,
			t.review_count
		`)

	if keyword != "" {
		searchKeyword := "%" + keyword + "%"

		query = query.Where(
			"(t.name LIKE ? OR t.teacher_no LIKE ?)",
			searchKeyword,
			searchKeyword,
		)
	}

	if college != "" {
		query = query.Where("t.college = ?", college)
	}

	var teachers []TeacherListItem

	if err := query.Order("t.id ASC").Scan(&teachers).Error; err != nil {
		utils.Fail(
			context,
			http.StatusInternalServerError,
			"教师列表查询失败",
		)
		return
	}

	utils.Success(context, "获取教师列表成功", teachers)
}

func GetTeacherFilters(context *gin.Context) {
	var colleges []string

	if err := config.DB.
		Table("teachers").
		Distinct().
		Order("college ASC").
		Pluck("college", &colleges).
		Error; err != nil {
		utils.Fail(
			context,
			http.StatusInternalServerError,
			"教师筛选条件查询失败",
		)
		return
	}

	utils.Success(context, "获取教师筛选条件成功", gin.H{
		"colleges": colleges,
	})
}

func GetTeacherDetail(context *gin.Context) {

	id := context.Param("id")

	var teacher TeacherDetail

	err := config.DB.
		Table("teachers AS t").
		Select(`
			t.id,
			t.teacher_no,
			t.name,
			t.college,
			t.title,
			t.introduction,
			t.average_score,
			t.review_count
		`).
		Where("t.id = ?", id).
		Scan(&teacher).Error

	if err != nil {
		utils.Fail(
			context,
			500,
			"教师信息查询失败",
		)
		return
	}

	var courses []gin.H

	config.DB.
		Table("courses").
		Where("teacher_id = ?", id).
		Find(&courses)

	var reviews []gin.H

	config.DB.
		Table(`
			teacher_reviews tr
			JOIN users u
			ON tr.user_id=u.id
		`).
		Select(`
			tr.score,
			tr.content,
			tr.created_at,
			u.name
		`).
		Where(
			"tr.teacher_id=?",
			id,
		).
		Order(
			"tr.created_at DESC",
		).
		Limit(20).
		Scan(&reviews)

	utils.Success(
		context,
		"获取教师详情成功",
		gin.H{
			"teacher": teacher,
			"courses": courses,
			"reviews": reviews,
		},
	)
}

type TeacherDetail struct {
	ID uint64 `json:"id"`

	TeacherNo string `json:"teacher_no"`

	Name string `json:"name"`

	College string `json:"college"`

	Title string `json:"title"`

	Introduction string `json:"introduction"`

	AverageScore float64 `json:"average_score"`

	ReviewCount int `json:"review_count"`
}
