package main

import (
	"fmt"
	"gotest/config"
	"gotest/database"
	"gotest/routes"
)

func main() {
	c := config.NewConfig()

	r := routes.NewRouter(c)
	conn := database.NewDatabaseConnection(c)

	fmt.Println(conn)

	r.Serve()
}
