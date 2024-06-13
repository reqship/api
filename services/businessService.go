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

func UpdateBusiness(ctx *gin.Context) {
	b_id := ctx.Param("business_id")
	b_id_int, err := strconv.Atoi(b_id)

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	business := models.Business{}
	if err := ctx.BindJSON(&business); err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	if int64(b_id_int) != business.ID {
		http.Forbidden(ctx, "business id's must match to make any changes")
		return
	}

	user_id := ctx.GetInt64("user_id")

	business.UserID = user_id

	err = business.Update()

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	http.Ok_no_body(ctx, "Business updated successfully")
}

func DeleteBusiness(ctx *gin.Context) {
	business := models.Business{}
	auto_business_id := ctx.GetStringSlice("business_ids")

	if err := ctx.BindJSON(&business); err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	b_id := ctx.Param("business_id")

	b_id_int, err := strconv.Atoi(b_id)

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	isInUsersBusinessList := false
	for _, business := range auto_business_id {
		bid, err := strconv.Atoi(business)
		if err != nil {
			http.Bad(ctx, err.Error())
			return
		}
		if int64(bid) == int64(b_id_int) {
			isInUsersBusinessList = true
		}

	}

	if !isInUsersBusinessList {
		http.Forbidden(ctx, "Authentication failed")
		return

	}

	user_id := ctx.GetInt64("user_id")

	business.UserID = user_id

	err = business.Delete()

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	http.Ok_no_body(ctx, "Successfully deleted business")
}

func GetAllBusinesses(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	count := ctx.DefaultQuery("count", "20")
	searchQuery := ctx.DefaultQuery("searchQuery", "")

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
	businesses, err := models.GetPaginatedBusinesses(page_int, count_int, searchQuery)
	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	http.Ok(ctx, businesses, "Successfully grabbed all businesses")
}

func GetBusinessById(ctx *gin.Context) {
	id := ctx.Param("business_id")

	id_int, err := strconv.Atoi(id)

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	business, err := models.GetBusinessByID(int64(id_int))

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	http.Ok(ctx, business, "successfully retrieved business")
}
