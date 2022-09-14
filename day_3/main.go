package main

import (
	"day_3/routes"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := ":8080"
	fmt.Printf("Successfully started on port %s\n", port)
	routes.Routes()
	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))
}
