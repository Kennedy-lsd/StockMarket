package handlers

import (
	"net/http"

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
