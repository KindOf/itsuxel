package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"

	"github.com/KindOf/itsuxel/api"
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

//	@title			ITSUXEL
//	@version		1.0
//	@description	Educational Excel-like API

//	@contact.email	iostapovychweb@gmail.com

//	@host		localhost:3000
//	@BasePath	/api/v1
func main() {
	// USE EXPR https://expr-lang.org/
	e := echo.New()
	// Root level middleware
	e.Use(NewLogger())
	e.Use(middleware.Recover())

	e.POST("/api/v1/table/:sheet/:cell", api.CreateCell)

	e.GET("/api/v1/table/:sheet/:cell", api.GetCell)

	e.GET("/api/v1/table/:sheet", api.GetSheet)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start("localhost:3000"))
}
