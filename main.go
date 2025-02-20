package main

import (
	"SalesBandung/config"
	"SalesBandung/router"
	"log"
)

func main() {
	config.ConnectDb()

	r := router.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
