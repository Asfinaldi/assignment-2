package main

import (
	"assignment-2/database"
	"assignment-2/routers"
)

func main() {
	database.StartDB()
	routers.StartServer().Run(":8080")
}
