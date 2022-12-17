package handler

import (
	"net/http"
	"todo-app/src/models"
	"todo-app/src/payload/response"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetActivityList(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := h.useCase.GetActivityList(ctx)
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

func (h *Handler) GetActivity(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	resp, err := h.useCase.GetActivity(ctx, id)
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

func (h *Handler) CreateActivity(c echo.Context) error {
	ctx := c.Request().Context()
	a := new(models.Activity)
	if err := c.Bind(a); err != nil {
		return c.JSON(http.StatusBadRequest, response.GlobalResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
	}

	resp, err := h.useCase.CreateActivity(ctx, a)
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

func (h *Handler) DeleteActivity(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	err := h.useCase.DeleteActivity(ctx, id)
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

func (h *Handler) UpdateActivity(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	body := new(models.Activity)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, response.GlobalResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
	}

	resp, err := h.useCase.UpdateActivity(ctx, body, id)
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
