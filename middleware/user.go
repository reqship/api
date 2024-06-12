package middleware

import (
	"context"
	"reqship-api/helpers/auth"
	"reqship-api/helpers/db"
	"reqship-api/helpers/http"
	"reqship-api/models"

	"github.com/gin-gonic/gin"
)

func GetUserDetails(ctx *gin.Context) {
	var username string

	token := ctx.GetHeader("Authorization")

	err := auth.CheckJWT(token, &username)

	if err != nil {
		http.Forbidden(ctx, "Invalid token")
		ctx.Abort()
	} else {
		db := db.Init()
		context := context.Background()
		defer db.Close()

		users := []models.User{}
		err := db.NewSelect().Model(&users).Where("username = ?", username).Scan(context)
		if err != nil {
			return
		}

		user := users[0]

		ctx.Set("user_id", user.ID)
		ctx.Set("user_username", user.Username)
		ctx.Set("user_email", user.Email)
		ctx.Next()
	}
}
