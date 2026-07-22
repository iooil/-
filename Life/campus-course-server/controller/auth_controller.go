package controller

import (
	"net/http"
	"strings"

	"campus-course-server/config"
	"campus-course-server/model"
	"campus-course-server/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	StudentNo string `json:"student_no" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Major     string `json:"major" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
}

func Register(context *gin.Context) {
	var request RegisterRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		utils.Fail(context, http.StatusBadRequest, "注册信息不完整或格式错误")
		return
	}

	request.StudentNo = strings.TrimSpace(request.StudentNo)
	request.Name = strings.TrimSpace(request.Name)
	request.Major = strings.TrimSpace(request.Major)

	if request.StudentNo == "" || request.Name == "" || request.Major == "" {
		utils.Fail(context, http.StatusBadRequest, "注册信息不能为空")
		return
	}

	var existingUser model.User

	result := config.DB.
		Where("student_no = ?", request.StudentNo).
		First(&existingUser)

	if result.Error == nil {
		utils.Fail(context, http.StatusConflict, "该学号已注册")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		utils.Fail(context, http.StatusInternalServerError, "密码加密失败")
		return
	}

	user := model.User{
		StudentNo: request.StudentNo,
		Name:      request.Name,
		Major:     request.Major,
		Password:  string(hashedPassword),
		Role:      "student",
		Status:    1,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.Fail(context, http.StatusInternalServerError, "注册失败")
		return
	}

	utils.Success(context, "注册成功", gin.H{
		"id":         user.ID,
		"student_no": user.StudentNo,
		"name":       user.Name,
		"major":      user.Major,
		"role":       user.Role,
	})
}

type LoginRequest struct {
	StudentNo string `json:"student_no" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func Login(context *gin.Context) {
	var request LoginRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		utils.Fail(context, http.StatusBadRequest, "请输入学号和密码")
		return
	}

	request.StudentNo = strings.TrimSpace(request.StudentNo)

	var user model.User

	result := config.DB.
		Where("student_no = ?", request.StudentNo).
		First(&user)

	if result.Error != nil {
		utils.Fail(context, http.StatusUnauthorized, "学号或密码错误")
		return
	}

	if user.Status != 1 {
		utils.Fail(context, http.StatusForbidden, "该账号已被禁用")
		return
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password),
	)

	if err != nil {
		utils.Fail(context, http.StatusUnauthorized, "学号或密码错误")
		return
	}

	token, err := utils.GenerateToken(
		user.ID,
		user.StudentNo,
		user.Role,
	)

	if err != nil {
		utils.Fail(context, http.StatusInternalServerError, "登录凭证生成失败")
		return
	}

	utils.Success(context, "登录成功", gin.H{
		"token": token,
		"user": gin.H{
			"id":         user.ID,
			"student_no": user.StudentNo,
			"name":       user.Name,
			"major":      user.Major,
			"role":       user.Role,
		},
	})
}
