package main

import (
	"github.com/Braullio/linkito/internal/database"
	"github.com/Braullio/linkito/internal/v1/links"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"log"

	_ "github.com/Braullio/linkito/docs"
	"github.com/gofiber/fiber/v2"
)

func init() {
	err := godotenv.Load("configs/development.env")
	if err != nil {
		log.Fatal("Error loading development.env file")
	}

	database.StartDB()
}

// @title                     Linkito API
// @version                   1.0
// @description               Encurtador de URL.
// @termsOfService            http://swagger.io/terms/
// @host                      localhost:3000
// @license.name              MIT License
// @license.url               https://opensource.org/license/MIT
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	http := fiber.New()

	routes(http)

	log.Fatal(
		http.Listen(":3000"),
	)
}

func routes(http *fiber.App) {
	http.Get("/swagger/*", swagger.HandlerDefault)
	http.Post("/links/migrate", links.Migrate)

	http.Get("/:id", links.Redirect)

	http.Post("/v1/links", links.Create)
	http.Get("/v1/links", links.ListAll)
	http.Get("/v1/links/:id", links.Search)
}
