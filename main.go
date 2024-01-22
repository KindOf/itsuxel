package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ValueResponse struct {
	Value  string `json:"value"`
	Result string `json:"result"`
}

func NewValueResponse(value string, result string) *ValueResponse {
	return &ValueResponse{value, result}
}

type TableResponse = []ValueResponse

func NewLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod: true,
		LogStatus: true,
		LogURI:    true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("%s %v: %v\n", v.Method, v.Status, v.URI)
			return nil
		},
	})
}

func main() {
    // USE EXPR https://expr-lang.org/
	e := echo.New()
	// Root level middleware
	e.Use(NewLogger())
	e.Use(middleware.Recover())

	e.POST("/api/v1/table/:sheet/:cell", func(c echo.Context) error {
		return c.JSON(http.StatusOK, NewValueResponse("value", "cell"))
	})

	e.GET("/api/v1/table/:sheet/:cell", func(c echo.Context) error {
		return c.JSON(http.StatusOK, NewValueResponse("value", "cell"))
	})

	e.GET("/api/v1/table/:sheet", func(c echo.Context) error {
		tr := []*ValueResponse{
			NewValueResponse("value1", "cell1"),
			NewValueResponse("value2", "cell2"),
			NewValueResponse("value3", "cell3"),
			NewValueResponse("value4", "cell4"),
		}

		return c.JSON(http.StatusOK, tr)
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start("localhost:3000"))
}
