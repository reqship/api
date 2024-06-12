package models

import "github.com/uptrace/bun"

type Order struct {
	bun.BaseModel `bun:"table:orders,alias:orders"`

	ID         int64 `bun:",pk,autoincrement"`
	BusinessId int64
}
