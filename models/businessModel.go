package models

import (
	"context"
	"errors"
	"math"
	"reqship-api/helpers/db"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type Business struct {
	bun.BaseModel `bun:"table:businesses,alias:business"`

	ID          int64 `bun:",pk,autoincrement:"`
	Name        string
	Description string
	UserID      int64
	ImageUrl    string
}

func GetPaginatedBusinesses(page int, count int, query string) (res gin.H, err error) {
	db := db.Init()
	ctx := context.Background()
	defer db.Close()

	b := []Business{}

	selectQuery := db.NewSelect()
	if query != "" {
		selectQuery = selectQuery.Where("name LIKE ?", "%"+query+"%")
	}

	item_count, err := selectQuery.Model(&b).Offset((page - 1) * count).Limit(count).ScanAndCount(ctx)

	res = gin.H{"total_items": item_count, "pages": math.Ceil(float64(item_count) / float64(count)), "current_page": page, "items": b}
	return
}

func GetBusinessByID(id int64) (b Business, err error) {
	db := db.Init()
	ctx := context.Background()
	defer db.Close()

	businesses := []Business{}

	err = db.NewSelect().Model(&businesses).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return
	}
	if len(businesses) == 0 {
		err = errors.New("cannot find business")
		return
	}
	b = businesses[0]
	return
}

func (b *Business) Create() (business *Business, err error) {
	db := db.Init()
	ctx := context.Background()
	defer db.Close()

	_, err = db.NewInsert().Model(b).Exec(ctx)
	return
}

func (b *Business) Update() (err error) {
	db := db.Init()
	ctx := context.Background()
	defer db.Close()

	_, err = db.NewUpdate().Model(b).WherePK().Exec(ctx)
	return
}

func (b *Business) Delete() (err error) {
	db := db.Init()
	ctx := context.Background()
	defer db.Close()

	_, err = db.NewDelete().Model(b).WherePK().Exec(ctx)
	return
}

func (b *Business) GetItems() (items []Item, err error) {
	return GetItemsByBusinessId(b.ID)
}
