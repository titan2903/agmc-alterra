package main

import (
	"day_2/routes"
	"fmt"
)

func main() {
	port := ":8080"
	fmt.Printf("Successfully started on port %s\n", port)
	routes.Routes()
	e := routes.Routes()
	e.Logger.Fatal(e.Start(":8080"))
}
