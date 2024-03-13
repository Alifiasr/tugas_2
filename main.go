package main

import (
	"tugas__2/config"
	"tugas__2/controller"
	"tugas__2/router"
)

var PORT = ":8080"

func main() {
	db := config.StarDB()
	OrderDB := controller.New(db)

	router.StarOrder(OrderDB).Run(PORT)

}
