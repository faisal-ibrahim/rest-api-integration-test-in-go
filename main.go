package main

import (
	"github.com/faisal-ibrahim/rest-api-integration-test-in-go.git/internal/database"
	"github.com/faisal-ibrahim/rest-api-integration-test-in-go.git/internal/transport"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := transport.Setup()
	database.InitDatabase()
	app.Listen(3000)
}
