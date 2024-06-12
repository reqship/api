package main

import (
	"os"
	db "reqship-api/helpers/migrate"
	"reqship-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 && args[0] == "migrate" {
		db.Migrate()
	} else {
		router := gin.Default()
		routes.Init(router)
	}
}
