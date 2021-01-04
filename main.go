package main

import (
	_ "github.com/joho/godotenv/autoload"
	"wumiao/routes"
)

func main() {
	go routes.BackendHtml()
	go routes.Api()
	routes.BackendStart()
}
