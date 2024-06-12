package models

import "github.com/uptrace/bun"

type Business struct {
	bun.BaseModel `bun:"table:businesses,alias:business"`

	ID          int64 `bun:",pk,autoincrement:"`
	Name        string
	Description string
	// Items
	// Orders

}
