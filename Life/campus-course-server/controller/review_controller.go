package controller

import (
	"fmt"

	"campus-course-server/config"
	"campus-course-server/model"
	"campus-course-server/utils"

	"github.com/gin-gonic/gin"
)

type ReviewRequest struct {
	CourseID uint64 `json:"course_id"`

	Score int `json:"score"`

	Content string `json:"content"`
}

func CreateCourseReview(context *gin.Context) {

	userID, _ := context.Get("userID")

	var req ReviewRequest

	if err := context.ShouldBindJSON(&req); err != nil {

		utils.Fail(
			context,
			400,
			"参数错误",
		)

		return
	}

	if req.Score < 1 || req.Score > 5 {

		utils.Fail(
			context,
			400,
			"评分范围1-5",
		)

		return

	}

	uid := userID.(uint64)

	// 判断是否选过该课程

	var count int64

	config.DB.
		Table("course_selections").
		Where(
			"user_id=? AND course_id=? AND status='selected'",
			uid,
			req.CourseID,
		).
		Count(&count)

	if count == 0 {

		utils.Fail(
			context,
			400,
			"未选择该课程，不能评价",
		)

		return

	}

	// 判断重复评价

	var exist int64

	config.DB.
		Table("course_reviews").
		Where(
			"user_id=? AND course_id=?",
			uid,
			req.CourseID,
		).
		Count(&exist)

	if exist > 0 {

		utils.Fail(
			context,
			400,
			"已经评价过该课程",
		)

		return

	}

	review := model.CourseReview{

		UserID: uid,

		CourseID: req.CourseID,

		Score: req.Score,

		Content: req.Content,
	}

	if err := config.DB.Create(&review).Error; err != nil {

		utils.Fail(
			context,
			500,
			"评价失败",
		)

		return

	}

	updateCourseScore(req.CourseID)

	utils.Success(
		context,
		"课程评价成功",
		review,
	)

}

func updateCourseScore(courseID uint64) {

	var avg float64

	var count int64

	config.DB.
		Table("course_reviews").
		Where(
			"course_id=?",
			courseID,
		).
		Select(
			"AVG(score)",
		).
		Scan(&avg)

	config.DB.
		Table("course_reviews").
		Where(
			"course_id=?",
			courseID,
		).
		Count(&count)

	config.DB.
		Table("courses").
		Where(
			"id=?",
			courseID,
		).
		Updates(map[string]interface{}{

			"average_score": avg,

			"review_count": count,
		})

}

func CreateTeacherReview(context *gin.Context) {

	userIDValue, _ := context.Get("userID")

	userID := userIDValue.(uint64)

	var req struct {
		TeacherID uint64 `json:"teacher_id"`

		CourseID uint64 `json:"course_id"`

		Score int `json:"score"`

		Content string `json:"content"`
	}

	if err := context.ShouldBindJSON(&req); err != nil {

		utils.Fail(
			context,
			400,
			"参数错误",
		)

		return
	}

	if req.Score < 1 || req.Score > 5 {

		utils.Fail(
			context,
			400,
			"评分范围1-5",
		)

		return
	}

	// 判断是否选择过该教师课程

	var count int64

	config.DB.
		Table("course_selections cs").
		Joins(
			"JOIN courses c ON cs.course_id=c.id",
		).
		Where(
			"cs.user_id=? AND c.teacher_id=? AND c.id=? AND cs.status='selected'",
			userID,
			req.TeacherID,
			req.CourseID,
		).
		Count(&count)

	if count == 0 {

		utils.Fail(
			context,
			400,
			"未选择该教师课程，不能评价",
		)

		return
	}

	// 防止重复评价

	var exist int64

	config.DB.
		Table("teacher_reviews").
		Where(
			"user_id=? AND teacher_id=? AND course_id=?",
			userID,
			req.TeacherID,
			req.CourseID,
		).
		Count(&exist)

	if exist > 0 {

		utils.Fail(
			context,
			400,
			"已经评价过该教师",
		)

		return

	}

	review := model.TeacherReview{

		UserID: userID,

		TeacherID: req.TeacherID,

		CourseID: req.CourseID,

		Score: req.Score,

		Content: req.Content,
	}

	if err := config.DB.Create(&review).Error; err != nil {

		utils.Fail(
			context,
			500,
			"教师评价失败",
		)

		return

	}

	updateTeacherScore(
		req.TeacherID,
	)

	utils.Success(
		context,
		"教师评价成功",
		review,
	)

}

func updateTeacherScore(teacherID uint64) {

	var avg float64

	var count int64

	// 查询平均分
	config.DB.
		Table("teacher_reviews").
		Where(
			"teacher_id = ?",
			teacherID,
		).
		Select(
			"COALESCE(AVG(score),0)",
		).
		Scan(&avg)

	// 查询评价数量
	config.DB.
		Table("teacher_reviews").
		Where(
			"teacher_id = ?",
			teacherID,
		).
		Count(&count)

	// 更新教师表
	result := config.DB.
		Table("teachers").
		Where(
			"id = ?",
			teacherID,
		).
		Updates(map[string]interface{}{

			"average_score": avg,

			"review_count": count,
		})

	if result.Error != nil {

		fmt.Println(
			"更新教师评分失败:",
			result.Error,
		)

	}

	fmt.Println(
		"教师评分更新:",
		teacherID,
		avg,
		count,
	)

}
