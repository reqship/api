package models

import "github.com/uptrace/bun"

type Item struct {
	bun.BaseModel `bun:"table:items,alias:items"`

	ID          int64 `bun:",pk,autoincrement"`
	BusinessID  int64
	Name        string
	Price       float32
	Description string
	ImageUrl    string
}
