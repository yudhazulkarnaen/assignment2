package main

import (
	"assignment2.id/orderapi/database"
	_ "assignment2.id/orderapi/docs"
	"assignment2.id/orderapi/routers"
)

// @title           Order API
// @version         1.0
// @description     Assignment 2.

// @contact.name   zulkarnaen
// @contact.email  premiumforspot@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	database.StartDB()
	var port = ":8080"
	routers.StartServer().Run(port)
}
