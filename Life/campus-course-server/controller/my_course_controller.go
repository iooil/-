package controller

import (
	"fmt"
	"net/http"

	"campus-course-server/config"
	"campus-course-server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MyCourseItem struct {
	ID uint64 `json:"id"`

	TeacherID uint64 `json:"teacher_id"`
	
	CourseName string `json:"course_name"`

	CourseNo string `json:"course_no"`

	Credit float64 `json:"credit"`

	College string `json:"college"`

	CourseType string `json:"course_type"`

	Description string `json:"description"`

	TeacherName string `json:"teacher_name"`

	Status string `json:"status"`

	SelectedAt string `json:"selected_at"`
}

func GetMyCourses(context *gin.Context) {

	userIDValue, exists := context.Get("userID")

	if !exists {

		utils.Fail(
			context,
			http.StatusUnauthorized,
			"用户未登录",
		)

		return
	}

	userID := userIDValue.(uint64)

	var courses []MyCourseItem

	err := config.DB.
		Table("course_selections AS cs").
		Select(`
		c.id,
		c.course_name,
		c.teacher_id,
		c.course_no,
		c.credit,
		c.college,
		c.course_type,
		c.description,
		t.name AS teacher_name,
		cs.status,
		cs.selected_at
	`).
		Joins(`
		LEFT JOIN courses AS c
		ON cs.course_id = c.id
	`).
		Joins(`
		LEFT JOIN teachers AS t
		ON c.teacher_id = t.id
	`).
		Where(
			"cs.user_id=? AND cs.status=?",
			userID,
			"selected",
		).
		Order(
			"cs.selected_at DESC",
		).
		Scan(&courses).
		Error

	if err != nil {

		utils.Fail(
			context,
			500,
			"查询我的课程失败",
		)

		return
	}

	utils.Success(
		context,
		"获取我的课程成功",
		courses,
	)

}

func DropCourse(context *gin.Context) {

	userIDValue, _ := context.Get("userID")

	userID := userIDValue.(uint64)

	courseID := context.Param("id")

	err := config.DB.Transaction(
		func(tx *gorm.DB) error {

			result := tx.Table(
				"course_selections",
			).
				Where(
					"user_id=? AND course_id=?",
					userID,
					courseID,
				).
				Update(
					"status",
					"dropped",
				)

			if result.Error != nil {
				return result.Error
			}

			tx.Table(
				"courses",
			).
				Where(
					"id=? AND selected_count>0",
					courseID,
				).
				UpdateColumn(
					"selected_count",
					gorm.Expr(
						"selected_count - 1",
					),
				)

			return nil

		})

	if err != nil {

		fmt.Println("查询我的课程错误:", err.Error())

		utils.Fail(
			context,
			500,
			err.Error(),
		)

		return
	}

	utils.Success(
		context,
		"退课成功",
		nil,
	)

}
