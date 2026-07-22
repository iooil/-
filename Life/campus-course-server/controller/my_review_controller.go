package controller

import (
	"campus-course-server/config"
	"campus-course-server/model"

	"campus-course-server/utils"

	"github.com/gin-gonic/gin"
)

func GetMyReviews(context *gin.Context) {

	userIDValue, _ := context.Get("userID")

	userID := userIDValue.(uint64)

	var courseReviews []model.MyCourseReview

	config.DB.
		Table("course_reviews cr").
		Select(`
	cr.id,
	cr.score,
	cr.content,
	cr.created_at,
	c.course_name
`).
		Joins(
			"LEFT JOIN courses c ON cr.course_id=c.id",
		).
		Where(
			"cr.user_id=?",
			userID,
		).
		Order(
			"cr.created_at DESC",
		).
		Scan(&courseReviews)

	var teacherReviews []model.MyTeacherReview

	config.DB.
		Table("teacher_reviews tr").
		Select(`
	tr.id,
	tr.score,
	tr.content,
	tr.created_at,
	t.name AS teacher_name
`).
		Joins(
			"LEFT JOIN teachers t ON tr.teacher_id=t.id",
		).
		Where(
			"tr.user_id=?",
			userID,
		).
		Order(
			"tr.created_at DESC",
		).
		Scan(&teacherReviews)

	utils.Success(
		context,
		"获取我的评价成功",
		gin.H{

			"course_reviews": courseReviews,

			"teacher_reviews": teacherReviews,
		},
	)

}
