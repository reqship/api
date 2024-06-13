package models

import (
	"context"
	"reqship-api/helpers/db"

	"github.com/uptrace/bun"
)

type Item struct {
	bun.BaseModel `bun:"table:items,alias:items"`

	ID          int64 `bun:",pk,autoincrement"`
	BusinessID  int64
	Name        string
	Price       float32
	Description string
	ImageUrl    string
}

func (i *Item) Create() (err error) {
	db := db.Init()
	ctx := context.Background()
	defer db.Close()

	_, err = db.NewInsert().Model(i).Exec(ctx)
	return
}

func GetItemById(id int64) (item Item, err error) {
	db := db.Init()
	ctx := context.Background()
	defer db.Close()
	items := []Item{}
	err = db.NewSelect().Model(&items).Where("id = ?", id).Scan(ctx)
	item = items[0]
	return
}

func GetItemsByBusinessId(businessId int64) (items []Item, err error) {
	db := db.Init()
	ctx := context.Background()
	defer db.Close()

	err = db.NewSelect().Model(&items).Where("business_id = ?", businessId).Scan(ctx)
	return
}
