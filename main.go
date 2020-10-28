package main

import (
	_ "github.com/joho/godotenv/autoload"
	"wumiao/routes"
)

func main() {
	go routes.FrontendStart()
	go routes.BackendHtml()
	routes.BackendStart()
}
