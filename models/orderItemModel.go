package models

import "github.com/uptrace/bun"

type OrderItem struct {
	bun.BaseModel `bun:"table:order_items,alias:o_items"`

	ID      int64 `bun:",pk,autoincrement"`
	OrderId int64
	Qty     int8
	ItemId  int64
}
