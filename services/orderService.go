package services

import (
	"reqship-api/helpers/http"
	"reqship-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// order router
func NewOrder(ctx *gin.Context) {
	order := models.FullOrder{}
	user_id := ctx.GetInt64("user_id") // comes from middleware

	if err := ctx.BindJSON(&order); err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	o := models.Order{
		ID:         order.ID,
		BusinessId: order.BusinessId,
		UserId:     user_id,
	}

	err := o.CreateOrderWithItems(order.Items)
	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	http.Ok_no_body(ctx, "Successfully created order")
}

// order router
func GetOrderContent(ctx *gin.Context) {
	order_id := ctx.Param("order_id")
	order_id_int, err := strconv.Atoi(order_id)
	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	user_id := ctx.GetInt64("user_id") // user id from middleware

	order, err := models.GetOrderById(int64(order_id_int))

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	if order.UserId != user_id {
		http.Forbidden(ctx, "user does not own this order")
		return
	}

	fullOrder, err := order.GetOrderWithItems()

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	http.Ok(ctx, fullOrder, "successfully retreived full order")
}

// business router
func GetOrdersByBusinessId(ctx *gin.Context) {
	business_id := ctx.Param("business_id")
	business_id_int, err := strconv.Atoi(business_id)

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	orders, err := models.GetOrdersByBusinessId(int64(business_id_int))
	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	http.Ok(ctx, orders, "successfully retreived orders")
}

// user router
func GetOrdersByUserId(ctx *gin.Context) {
	order_id := ctx.Param("order_id")
	order_id_int, err := strconv.Atoi(order_id)
	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	orders, err := models.GetOrdersByUserId(int64(order_id_int))
	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	http.Ok(ctx, orders, "successfully retreived orders")

}

// order router
func CompleteOrder(ctx *gin.Context) {
	order_id := ctx.Param("order_id")
	order_id_int, err := strconv.Atoi(order_id)
	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	orderOwnedByUser := false
	business_ids := ctx.GetStringSlice("business_ids") // comes from middleware

	// check order in businesses list
	// check business has order
	order, err := models.GetOrderById(int64(order_id_int))
	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	for _, business_id := range business_ids {
		business_id_int, err := strconv.Atoi(business_id)
		if err != nil {
			http.Bad(ctx, err.Error())
			return
		}

		if order.BusinessId == int64(business_id_int) {
			orderOwnedByUser = true
			break
		}
	}

	if !orderOwnedByUser {
		http.Forbidden(ctx, "you cannot finalize this order")
		return
	}

	err = order.CompleteOrder()
	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}
	http.Ok_no_body(ctx, "successfully completed order")

}
