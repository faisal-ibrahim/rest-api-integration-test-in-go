//go:build integration

package book_test

import "github.com/faisal-ibrahim/rest-api-integration-test-in-go.git/internal/transport"

type BookTestSuite struct {
	suite.Suite
	app *fiber.App
}

func (suite *BookTestSuite) SetupSuite() {
	suite.app = transport.Setup()
}
