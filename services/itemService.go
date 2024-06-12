package services

import (
	"fmt"
	"reqship-api/helpers/http"
	"reqship-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateItem(ctx *gin.Context) {
	item := models.Item{}
	business_ids := ctx.GetStringSlice("business_ids")
	if err := ctx.BindJSON(&item); err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	user_owns_business := false
	user_provided_business_id := strconv.Itoa(int(item.BusinessID))
	for _, business := range business_ids {
		fmt.Println(business, user_provided_business_id)

		if business == user_provided_business_id {

			user_owns_business = true
			break
		}
	}

	if !user_owns_business {
		http.Forbidden(ctx, "you cannot add an item to a business you do not own")
		return
	}
	err := item.Create()

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	http.Ok_no_body(ctx, "successfully created item")
}

func GetItemsByBusiness(ctx *gin.Context) {
	business_id_str := ctx.Param("business_id")

	business_id, err := strconv.Atoi(business_id_str)

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	business, err := models.GetBusinessByID(int64(business_id))

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	items, err := business.GetItems()

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	http.Ok(ctx, items, "successfully retrieved items")
}
