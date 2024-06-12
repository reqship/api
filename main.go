package main

import (
	"os"
	db "reqship-api/helpers/migrate"
	"reqship-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	args := []string{}
	if len(args) > 1 {
		args = os.Args[1:]
	}

	if len(args) > 0 && args[0] == "migrate" {
		db.Migrate()
		return
	}

	router := gin.Default()
	routes.Init(router)

}
