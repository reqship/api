package models

import (
	"context"
	"reqship-api/helpers/db"

	"github.com/uptrace/bun"
)

type OrderItem struct {
	bun.BaseModel `bun:"table:order_items,alias:o_items"`

	ID      int64 `bun:",pk,autoincrement"`
	OrderId int64
	Qty     int8
	ItemId  int64
	// TotalPrice float32 // should be added eventually when better queries are used
}

func GetOrderItemsByOrderId(orderId int64) (orderItems []OrderItem, err error) {
	db := db.Init()
	defer db.Close()
	ctx := context.Background()

	_, err = db.NewSelect().Model(&orderItems).Where("order_id = ?", orderId).Exec(ctx)
	return
}

func GetSingleOrderItemById(id int64) (orderItem OrderItem, err error) {
	db := db.Init()
	defer db.Close()
	ctx := context.Background()

	_, err = db.NewSelect().Model(&orderItem).Where("id = ?", id).Exec(ctx)
	return
}

func (oi *OrderItem) AddToDb() (err error) {
	db := db.Init()
	defer db.Close()
	ctx := context.Background()

	_, err = db.NewInsert().Model(oi).Exec(ctx)
	return
}

func (oi *OrderItem) RemoveFromDb() (err error) {
	db := db.Init()
	ctx := context.Background()
	defer db.Close()

	_, err = db.NewDelete().Model(oi).WherePK().Exec(ctx)
	return
}
