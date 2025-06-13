package main 

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponce struct {
	Message string `json:"message"`
}

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) CreateTodoTask(c *gin.Context) {
	var task Todo

	if err := c.BindJSON(&task); err != nil {
		fmt.Printf("failed to bind Todo task: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return 
	}

	h.storage.Insert(&task)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": task.ID,
	})
}

func (h *Handler) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("Failed to convert  id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	var task Todo

	if err := c.BindJSON(&task); err != nil {
		fmt.Printf("failed to bind task: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return 
	}

	h.storage.Update(id, task)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": task.ID,
	})
}

func (h *Handler) GetTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	task, err := h.storage.Get(id)
	if err != nil {
		fmt.Printf("failed to get task %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) DeleteTask(c *gin.Context) {
	id, err  := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	h.storage.Delete(id)

	c.String(http.StatusOK, "task deleted")
}