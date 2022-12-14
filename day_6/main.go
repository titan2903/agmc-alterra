package main

import (
	"day_6/internal/routes"
	"fmt"
)

func main() {
	port := ":8080"
	e := routes.NewRoutes()
	e.Logger.Fatal(e.Start(":8080"))
	fmt.Printf("Successfully started on port %s\n", port)
}
