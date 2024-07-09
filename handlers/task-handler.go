package handlers

import (
	"job_task/models"
	"job_task/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	task repositories.TaskLogic
}

func TaskHandler(task repositories.TaskLogic) *Handler {
	return &Handler{task}
}

//Task One
func (h *Handler) TaskOne(c echo.Context) error {
	req := new(models.TaskOneRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.task.TaskOne(req.Nums, req.Target)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.TaskOneResponse{Result: result})
}

//Task Two
func (h *Handler) TaskTwo(c echo.Context) error {
	req := new(models.TaskTwoRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := h.task.TaskTwo(req.Nums)
	return c.JSON(http.StatusOK, models.TaskTwoResponse{Result: result})
}

//Task Three Tidak Selesai
func (h *Handler) TaskThree(c echo.Context) error {
	req := new(models.TaskThreeRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := h.task.TaskThree(req.S, req.Words)
	return c.JSON(http.StatusOK, models.TaskThreeResponse{Result: result})
}