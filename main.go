package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"

	"github.com/KindOf/itsuxel/api"
	"github.com/KindOf/itsuxel/data"
	_ "github.com/KindOf/itsuxel/docs"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

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

// @host		localhost:3000
// @BasePath	/api/v1
func main() {
	// USE EXPR https://expr-lang.org/
	data.ConnectStorage()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// Root level middleware
	e.Use(NewLogger())
	e.Use(middleware.Recover())

	e.POST("/api/v1/table/:sheet/:cell", api.CreateCell)

	e.GET("/api/v1/table/:sheet/:cell", api.GetCell)

	e.GET("/api/v1/table/:sheet", api.GetSheet)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":3000"))
}
