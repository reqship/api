package middleware

import (
	"context"
	"reqship-api/helpers/auth"
	"reqship-api/helpers/db"
	"reqship-api/helpers/http"
	"reqship-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBusinessDetails(ctx *gin.Context) {
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

		businesses := []models.Business{}
		err = db.NewSelect().Model(&businesses).Where("user_id = ?", user.ID).Scan(context)
		if err != nil {
			return
		}

		business_ids := []string{}
		business_names := []string{}
		for _, business := range businesses {
			business_id := strconv.Itoa(int(business.ID))
			business_ids = append(business_ids, business_id)
			business_names = append(business_names, business.Name)
		}

		ctx.Set("business_ids", business_ids)
		ctx.Set("business_names", business_names)

		ctx.Next()
	}
}
