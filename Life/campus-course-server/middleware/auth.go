package middleware

import (
	"net/http"
	"strings"

	"campus-course-server/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authorization := context.GetHeader("Authorization")

		if authorization == "" {
			utils.Fail(context, http.StatusUnauthorized, "请先登录")
			context.Abort()
			return
		}

		parts := strings.SplitN(authorization, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Fail(context, http.StatusUnauthorized, "Token格式错误")
			context.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])

		if err != nil {
			utils.Fail(context, http.StatusUnauthorized, "登录状态已失效")
			context.Abort()
			return
		}

		context.Set("userID", claims.UserID)
		context.Set("studentNo", claims.StudentNo)
		context.Set("role", claims.Role)

		context.Next()
	}
}
