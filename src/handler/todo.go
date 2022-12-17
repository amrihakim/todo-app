package handler

import (
	"net/http"
	"todo-app/src/models"
	"todo-app/src/payload/response"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetTodoList(c echo.Context) error {
	ctx := c.Request().Context()
	param := c.QueryParam("activity_group_id")

	resp, err := h.useCase.GetTodoList(ctx, param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.GlobalResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.GlobalResponse{
		Status:  "Success",
		Message: "Success",
		Data:    resp,
	})
}

func (h *Handler) GetTodo(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	resp, err := h.useCase.GetTodo(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.GlobalResponse{
			Status:  "Not Found",
			Message: err.Error(),
			Data:    resp,
		})
	}
	return c.JSON(http.StatusOK, response.GlobalResponse{
		Status:  "Success",
		Message: "Success",
		Data:    resp,
	})
}

func (h *Handler) CreateTodo(c echo.Context) error {
	ctx := c.Request().Context()
	a := new(models.Todo)
	if err := c.Bind(a); err != nil {
		return c.JSON(http.StatusBadRequest, response.GlobalResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
	}

	resp, err := h.useCase.CreateTodo(ctx, a)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.GlobalResponse{
			Status:  "Bad Request",
			Message: err.Error(),
			Data:    resp,
		})
	}
	return c.JSON(http.StatusOK, response.GlobalResponse{
		Status:  "Success",
		Message: "Success",
		Data:    resp,
	})
}

func (h *Handler) DeleteTodo(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	err := h.useCase.DeleteTodo(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.GlobalResponse{
			Status:  "Not Found",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.GlobalResponse{
		Status:  "Success",
		Message: "Success",
	})
}

func (h *Handler) UpdateTodo(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	body := new(models.Todo)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, response.GlobalResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
	}

	resp, err := h.useCase.UpdateTodo(ctx, body, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.GlobalResponse{
			Status:  "Not Found",
			Message: err.Error(),
			Data:    resp,
		})
	}
	return c.JSON(http.StatusOK, response.GlobalResponse{
		Status:  "Success",
		Message: "Success",
		Data:    resp,
	})
}
