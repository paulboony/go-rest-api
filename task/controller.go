package task

import (
	"net/http"
	"paulboony/go-rest-api/util"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	service TaskService
}

func NewController(s TaskService) TaskController {
	return &controller{
		service: s,
	}
}

func (c *controller) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, util.ResourcePayload(c.service.FindAll()))
}

func (c *controller) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	v, err := c.service.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, util.ErrorsPayload(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.ResourcePayload(v))
}

func (c *controller) Create(ctx *gin.Context) {
	var task Task

	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorsPayload(err.Error()))
		return
	}

	v, err := c.service.Create(task)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorsPayload(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, util.ResourcePayload(v))
}

func (c *controller) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var newTask Task

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorsPayload(err.Error()))
		return
	}

	currentTask, err := c.service.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, util.ErrorsPayload(err.Error()))
		return
	}

	currentTask.Title = newTask.Title
	updatedTask, err := c.service.Update(*currentTask)

	if err != nil {
		ctx.JSON(http.StatusNotFound, util.ErrorsPayload(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.ResourcePayload(updatedTask))
}

func (c *controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := c.service.FindById(id); err != nil {
		ctx.JSON(http.StatusNotFound, util.ErrorsPayload(err.Error()))
	}
	c.service.Delete(id)
	ctx.JSON(http.StatusNoContent, nil)
}

func Route(r *gin.Engine, c TaskController) {
	r.GET("/tasks", c.FindAll)
	r.GET("/tasks/:id", c.FindById)
	r.POST("/tasks", c.Create)
	r.PATCH("/tasks/:id", c.Update)
	r.DELETE("/tasks/:id", c.Delete)
}
