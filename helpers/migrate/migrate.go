package migrate

import (
	"context"
	"fmt"

	"reqship-api/helpers/db"
	"reqship-api/models"

	"github.com/uptrace/bun"
)

var tables []any = []any{
	(*models.User)(nil),
	(*models.Business)(nil),
	(*models.Item)(nil),
	(*models.Order)(nil),
	(*models.OrderItem)(nil),
}

func createTable(database *bun.DB, table any) {
	ctx := context.Background()
	if err := database.ResetModel(ctx, table); err != nil {
		fmt.Println(err)
	}
}

func Migrate() {
	database := db.Init()
	for _, table := range tables {
		createTable(database, table)
	}
}
