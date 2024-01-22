package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
    "github.com/swaggo/echo-swagger"

	_ "github.com/KindOf/itsuxel/docs"
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

// @title ITSUXEL
// @version 1.0
// @description Educational Excel-like API

// @contact.email iostapovychweb@gmail.com

// @host localhost:3000
// @BasePath /api/v1
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start("localhost:3000"))
}
