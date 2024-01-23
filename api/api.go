package api

import (
	"net/http"
	"sync"

	"github.com/KindOf/itsuxel/data"
	"github.com/labstack/echo/v4"
)

var lock = sync.Mutex{}

type (
	SheetParam struct {
		Sheet string `json:"sheet" param:"sheet" validate:"required"`
	}

	CellParam struct {
		Cell string `json:"cell" param:"cell" validate:"required"`
	}

	CellAddr struct {
		SheetParam
		CellParam
	}

	CellValue struct {
		Value string `json:"value" validate:"required"`
	}

	SetCellValue struct {
		CellAddr
		CellValue
	}

	ValueResponse struct {
		CellAddr
		Value  string `json:"value"`
		Result string `json:"result"`
	}

	HTTPError struct {
		Code     int         `json:"-"`
		Message  interface{} `json:"message"`
		Internal error       `json:"-"`
	}
)

func NewValueResponse(addr *CellAddr, value string, result string) *ValueResponse {
	return &ValueResponse{*addr, value, result}
}

// CreateCell adds value to cell
//
//	@Summary		adds value to cell
//	@Description	Add a new pet to the store
//	@Accept			json
//	@Produce		json
//	@Param			sheet	path		string		true	"sheet name"
//	@Param			cell	path		string		true	"cell address"
//	@Param			json	body		CellValue	true	"Set Cell Value"
//	@Success		201		{object}	string		"ok"
//	@Failure		400		{object}	HTTPError
//	@Router			/table/{sheet}/{cell} [post]
func CreateCell(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	cell := &SetCellValue{}

	if err := c.Bind(cell); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(cell); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := data.SetCell(cell.Sheet, cell.Cell, cell.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = c.String(http.StatusCreated, "ok")
	if err != nil {
		return nil
	}
	return nil
}

// GetCell returns cell value
//
//	@Summary	returns cell value
//	@Accept		json
//	@Produce	json
//	@Param		sheet	path		string			true	"sheet name"
//	@Param		cell	path		string			true	"cell address"
//	@Success	200		{object}	ValueResponse	"ok"
//	@Failure	400		{object}	HTTPError
//	@Failure	404		{object}	HTTPError
//	@Router		/table/{sheet}/{cell} [get]
func GetCell(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	addr := &CellAddr{}

	if err := c.Bind(addr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(addr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	val, err := data.GetCell(addr.Sheet, addr.Cell)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, NewValueResponse(addr, val, val))
}

// GetSheet returns sheet cells
//
//	@Summary	returns cell value
//	@Accept		json
//	@Produce	json
//	@Param		sheet	path		string			true	"sheet name"
//	@Success	200		{object}	[]ValueResponse	"ok"
//	@Failure	404		{object}	HTTPError
//	@Failure	400		{object}	HTTPError
//	@Router		/table/{sheet} [get]
func GetSheet(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	addr := &SheetParam{}

	if err := c.Bind(addr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(addr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	m, err := data.GetSheet(addr.Sheet)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	tr := []*ValueResponse{}

	for k, v := range m {
		tr = append(tr, NewValueResponse(&CellAddr{SheetParam{addr.Sheet}, CellParam{k}}, v, v))
	}

	return c.JSON(http.StatusOK, tr)
}
