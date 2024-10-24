package handlers

import (
	"net/http"

	"github.com/Kennedy-lsd/StockMarket/data"
	"github.com/Kennedy-lsd/StockMarket/internal/services"
	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	CommentService *services.CommentService
}

func NewCommentHandler(s *services.CommentService) *CommentHandler {
	return &CommentHandler{
		CommentService: s,
	}
}

func (h *CommentHandler) GetAllComments(c echo.Context) error {
	comments, err := h.CommentService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(comments) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Not Found",
		})
	}

	return c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) CreateComment(c echo.Context) error {
	comment := new(data.CommentCreate)

	if err := c.Bind(comment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input: "+err.Error())
	}

	err := h.CommentService.Create(comment)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating comment: "+err.Error())
	}

	return c.JSON(http.StatusCreated, comment)
}
