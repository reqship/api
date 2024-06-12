package middleware

import (
	"reqship-api/helpers/auth"
	"reqship-api/helpers/http"

	"github.com/gin-gonic/gin"
)

func CheckUserAccess(ctx *gin.Context) {
	var username string

	token := ctx.GetHeader("Authorization")

	err := auth.CheckJWT(token, &username)
	if err != nil {
		http.Forbidden(ctx, "Invalid token")
		ctx.Abort()
	} else {
		ctx.Set("username", username)
		ctx.Next()
	}
}
