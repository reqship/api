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
