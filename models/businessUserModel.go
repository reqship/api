package models

import "github.com/uptrace/bun"

type BusinessUser struct {
	bun.BaseModel `bun:"table:business_users,alias:b_users"`

	ID           int64 `bun:",pk,autoincrement:"`
	BusinessName string
	Email        string
	Password     string
	BusinessId   int64
	// ProfileImage string

}
