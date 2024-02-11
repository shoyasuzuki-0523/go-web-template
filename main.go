package main

import (
	"kakeru-pro-web/common/config"
	"kakeru-pro-web/common/db"
	"kakeru-pro-web/common/router"
)

func main() {
	config.Init()
	db := db.Init()
	router.Init(db)
}
