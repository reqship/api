package models

import (
	"context"
	"errors"
	"reqship-api/helpers/db"

	"github.com/uptrace/bun"
)

type Order struct {
	bun.BaseModel `bun:"table:orders,alias:orders"`

	ID            int64 `bun:",pk,autoincrement"`
	UserId        int64
	BusinessId    int64
	NumberOfItems int64
	TotalPrice    float32
	Completed     bool
}

type FullOrder struct {
	ID            int64
	UserId        int64
	BusinessId    int64
	NumberOfItems int64
	TotalPrice    float32
	Completed     bool
	Items         []OrderItem
}

func (o *Order) updateOrder() (err error) {
	db := db.Init()
	defer db.Close()
	ctx := context.Background()

	_, err = db.NewUpdate().Model(&o).WherePK().Exec(ctx)
	return
}

func (o *Order) GetOrderWithItems() (fullOrder FullOrder, err error) {
	db := db.Init()
	defer db.Close()
	ctx := context.Background()
	fullOrder = FullOrder{
		ID:            o.ID,
		BusinessId:    o.BusinessId,
		TotalPrice:    o.TotalPrice,
		NumberOfItems: o.NumberOfItems,
		Completed:     o.Completed,
		UserId:        o.UserId,
	}

	err = db.NewSelect().Model(&fullOrder.Items).Where("order_id = ?", o.ID).Scan(ctx)
	return
}

func (o *Order) addOrderToDb() (id int64, err error) {
	db := db.Init()
	defer db.Close()
	ctx := context.Background()

	res, err := db.NewInsert().Model(o).Exec(ctx)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

func (o *Order) CreateOrderWithItems(items []OrderItem) (err error) {
	order := &Order{BusinessId: o.BusinessId}
	inserted_id, err := order.addOrderToDb()
	if err != nil {
		return errors.New("failed to add order to database")
	}

	o.Completed = false
	// TODO: this should be cleaned into a single sql query
	for _, item := range items {
		item.OrderId = inserted_id
		o.NumberOfItems += int64(item.Qty)
		product, err := GetItemById(item.ItemId)
		if err != nil {
			break
		}
		o.TotalPrice += product.Price * float32(item.Qty)
		err = item.AddToDb()
		if err != nil {
			break
		}
	}
	err = o.updateOrder() // update the order with the calculated totals
	return
}

func GetOrderById(id int64) (order Order, err error) {
	db := db.Init()
	defer db.Close()
	ctx := context.Background()

	orders := []Order{}

	err = db.NewSelect().Model(&orders).Where("id = ?", id).Scan(ctx)

	order = orders[0]
	return
}

func GetOrdersByBusinessId(businessId int64) (orders []FullOrder, err error) {
	db := db.Init()
	defer db.Close()
	ctx := context.Background()

	os := []Order{}

	err = db.NewSelect().Model(&os).Where("business_id = ?", businessId).Scan(ctx)

	for _, o := range os {
		fo, err := o.GetOrderWithItems()

		if err != nil {
			break
		}
		orders = append(orders, fo)
	}
	return
}

func GetOrdersByUserId(userId int64) (orders []FullOrder, err error) {
	db := db.Init()
	defer db.Close()
	ctx := context.Background()

	os := []Order{}

	err = db.NewSelect().Model(&os).Where("user_id = ?", userId).Scan(ctx)

	for _, o := range os {
		fo, err := o.GetOrderWithItems()

		if err != nil {
			break
		}
		orders = append(orders, fo)
	}
	return
}

func (o *Order) CompleteOrder() (err error) {
	o.Completed = true
	err = o.updateOrder()
	return
}
