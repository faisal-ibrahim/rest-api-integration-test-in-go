//go:build integration

package book_test

import (
	"github.com/faisal-ibrahim/rest-api-integration-test-in-go.git/internal/book"
	"github.com/faisal-ibrahim/rest-api-integration-test-in-go.git/internal/database"
	"github.com/faisal-ibrahim/rest-api-integration-test-in-go.git/internal/transport"
	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
)

type BookTestSuite struct {
	suite.Suite
	app *fiber.App
}

func (suite *BookTestSuite) SetupSuite() {
	suite.app = transport.Setup()
	database.InitDatabase()
	database.DBConn.AutoMigrate(&book.Book{})
}

func (suite *BookTestSuite) TearDownSuite() {
	err := os.Remove("books.db")
	if err != nil {

	}
}

func (suite *BookTestSuite) TestCreateBook() {
	req := httptest.NewRequest(
		"POST",
		"/api/v1/book",
		strings.NewReader(`{"ISIN": 12345, "title": "Test book", "author": "Faisal", "rating": 5}`),
	)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))

	res, err := suite.app.Test(req, -1)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)
}

func TestBookTestSuite(t *testing.T) {
	suite.Run(t, new(BookTestSuite))
}
