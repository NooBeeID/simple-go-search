package main

import (
	"fmt"
	"simple-go-search/config"
	"simple-go-search/db"
	"simple-go-search/server/router"
)

func main() {
	appDB := db.NewGormDB().GetConnection()
	if appDB != nil {
		fmt.Println("connected")
	}

	port := fmt.Sprintf(":%s", config.GetString(config.APP_PORT, "3001"))
	appRotuer := router.NewRouterGin(port)
	appRotuer.Start()

}
