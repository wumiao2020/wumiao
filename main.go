package main

import (
	_ "github.com/joho/godotenv/autoload"
	"wumiao/routes"
)

func main() {
	go routes.BackendHtml()
	go routes.FrontendHtml()
	go routes.Frontend()
	routes.Backend()
}
