package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ValueResponse struct {
	Value  string `json:"value"`
	Result string `json:"result"`
}

func NewValueResponse(value string, result string) *ValueResponse {
	return &ValueResponse{value, result}
}

// CreateCell adds value to cell
//
//	@Summary		adds value to cell
//	@Description	Add a new pet to the store
//	@Accept			json
//	@Produce		json
//	@Param			sheet	path		string	true	"sheet name"
//	@Param			cell	path		string	true	"cell address"
//	@Success		200		{string}	string	"ok"
//	@Router			/api/v1/table/{sheet}/{cell} [post]
func CreateCell(c echo.Context) error {
	return c.JSON(http.StatusOK, NewValueResponse("value", "cell"))
}

func GetCell(c echo.Context) error {
	return c.JSON(http.StatusOK, NewValueResponse("value", "cell"))
}

func GetSheet(c echo.Context) error {
	tr := []*ValueResponse{
		NewValueResponse("value1", "cell1"),
		NewValueResponse("value2", "cell2"),
		NewValueResponse("value3", "cell3"),
		NewValueResponse("value4", "cell4"),
	}

	return c.JSON(http.StatusOK, tr)
}