package main

import (
	"day_3/routes"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	port := ":8080"
	fmt.Printf("Successfully started on port %s\n", port)
	routes.Routes()
	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))
}
