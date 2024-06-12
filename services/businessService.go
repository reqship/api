package services

import (
	"reqship-api/helpers/http"
	"reqship-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBusiness(ctx *gin.Context) {
	business := models.Business{}
	if err := ctx.BindJSON(&business); err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	user_id := ctx.GetInt64("user_id")

	business.UserID = user_id

	business.Create()
	http.Ok_no_body(ctx, "Business created successfully")
}

func GetAllBusinesses(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	count := ctx.DefaultQuery("count", "20")

	page_int, err := strconv.Atoi(page)

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	count_int, err := strconv.Atoi(count)
	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	businesses, err := models.GetPaginatedBusinesses(page_int, count_int)
	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	http.Ok(ctx, businesses, "Successfully grabbed all businesses")
}
