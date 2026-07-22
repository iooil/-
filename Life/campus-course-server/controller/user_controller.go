package controller

import (
	"net/http"

	"campus-course-server/config"
	"campus-course-server/model"
	"campus-course-server/utils"

	"github.com/gin-gonic/gin"
)

func GetProfile(context *gin.Context) {
	userIDValue, exists := context.Get("userID")

	if !exists {
		utils.Fail(context, http.StatusUnauthorized, "无法获取登录用户")
		return
	}

	userID, ok := userIDValue.(uint64)

	if !ok {
		utils.Fail(context, http.StatusInternalServerError, "用户身份解析失败")
		return
	}

	var user model.User

	result := config.DB.First(&user, userID)

	if result.Error != nil {
		utils.Fail(context, http.StatusNotFound, "用户不存在")
		return
	}

	utils.Success(context, "获取个人信息成功", gin.H{
		"id":         user.ID,
		"student_no": user.StudentNo,
		"name":       user.Name,
		"major":      user.Major,
		"role":       user.Role,
		"status":     user.Status,
		"created_at": user.CreatedAt,
	})
}

func GetStudentDashboard(context *gin.Context) {
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

	var selectedCourseCount int64
	var courseReviewCount int64
	var teacherReviewCount int64

	config.DB.
		Table("course_selections").
		Where("user_id = ? AND status = ?", userID, "selected").
		Count(&selectedCourseCount)

	config.DB.
		Table("course_reviews").
		Where("user_id = ?", userID).
		Count(&courseReviewCount)

	config.DB.
		Table("teacher_reviews").
		Where("user_id = ?", userID).
		Count(&teacherReviewCount)

	utils.Success(context, "获取首页数据成功", gin.H{
		"selected_course_count": selectedCourseCount,
		"course_review_count":   courseReviewCount,
		"teacher_review_count":  teacherReviewCount,
	})
}
