package main

import (
	"be13/ca/config"
	"be13/ca/factory"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.ConnectToDB()
	defer db.Close()

	e := fiber.New()

	factory.InitFactory(e, db)
	e.Listen(fmt.Sprintf(":%d", 8080))

}
