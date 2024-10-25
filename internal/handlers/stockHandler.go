package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Kennedy-lsd/StockMarket/data"
	"github.com/Kennedy-lsd/StockMarket/internal/services"
	"github.com/labstack/echo/v4"
)

type StockHandler struct {
	StockService *services.StockService
}

func NewStockHandler(s *services.StockService) *StockHandler {
	return &StockHandler{
		StockService: s,
	}
}

func (h *StockHandler) GetAllStocks(c echo.Context) error {
	stocks, err := h.StockService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(stocks) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "No stocks found",
		})
	}

	return c.JSON(http.StatusOK, stocks)
}

func (h *StockHandler) CreateStock(c echo.Context) error {
	stock := new(data.CreatedStock)

	if err := c.Bind(stock); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input: "+err.Error())
	}

	err := h.StockService.Create(stock)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating stock: "+err.Error())
	}

	return c.JSON(http.StatusCreated, stock)
}

func (h *StockHandler) GetStockById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid id param"})
	}
	stock, err := h.StockService.GetById(id)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Comment not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Faild to fetch"})
	}

	return c.JSON(http.StatusOK, stock)
}

func (h *StockHandler) DeleteStockById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Id param"})
	}

	deleteError := h.StockService.DeleteById(id)

	if deleteError != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Not Found"})
	}
	return c.JSON(http.StatusOK, map[string]string{"error": "Stock was deleted"})
}
