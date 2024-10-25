package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

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

func (h *CommentHandler) GetCommentById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid comment ID"})
	}

	comment, err := h.CommentService.GetById(id)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Comment not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch comment"})
	}

	return c.JSON(http.StatusOK, comment)

}

func (h *CommentHandler) DeleteCommentById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid comment ID"})
	}

	deleteError := h.CommentService.DeleteById(id)

	if deleteError != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Not Found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Comment was deleted"})

}

func (h *CommentHandler) UpdateCommentById(c echo.Context) error {
	comment := new(data.CommentUpdate)

	if err := c.Bind(comment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input: "+err.Error())
	}
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid comment ID"})
	}

	updateError := h.CommentService.UpdateById(id, comment)

	if updateError != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Not Found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Comments was updated"})
}
