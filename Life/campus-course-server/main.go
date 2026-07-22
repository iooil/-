package main

import (
	"net/http"

	"campus-course-server/config"
	"campus-course-server/controller"
	"campus-course-server/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDatabase()

	router := gin.Default()

	router.GET("/api/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "校园选课评价系统后端连接成功",
			"data": gin.H{
				"system":   "校园选课评价系统",
				"backend":  "Go + Gin",
				"database": "MySQL",
			},
		})
	})

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", controller.Register)
			auth.POST("/login", controller.Login)
		}

		user := api.Group("/user")
		user.Use(middleware.AuthMiddleware())
		{
			user.GET("/profile", controller.GetProfile)
			user.GET("/dashboard", controller.GetStudentDashboard)
		}

		course := api.Group("/courses")
		course.Use(middleware.AuthMiddleware())
		{
			course.GET("", controller.GetCourseList)
			course.GET("/filters", controller.GetCourseFilters)
			course.POST("/:id/select", controller.SelectCourse)
		}

		review := api.Group("/reviews")

		review.Use(
			middleware.AuthMiddleware(),
		)

		{

			review.POST(
				"/course",
				controller.CreateCourseReview,
			)

			review.POST(
				"/teacher",
				controller.CreateTeacherReview,
			)

		}

		teacher := api.Group("/teachers")
		teacher.Use(middleware.AuthMiddleware())
		{
			teacher.GET("", controller.GetTeacherList)
			teacher.GET("/filters", controller.GetTeacherFilters)
			teacher.GET("/:id", controller.GetTeacherDetail)
		}

		my := api.Group("/my")

		my.Use(
			middleware.AuthMiddleware(),
		)

		my.GET(
			"/reviews",
			controller.GetMyReviews,
		)

		{

			my.GET(
				"/courses",
				controller.GetMyCourses,
			)

			my.POST(
				"/courses/:id/drop",
				controller.DropCourse,
			)

		}
	}

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
