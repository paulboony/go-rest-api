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

	var task Task

	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorsPayload(err.Error()))
		return
	}

	if _, err := c.service.FindById(id); err != nil {
		ctx.JSON(http.StatusNotFound, util.ErrorsPayload(err.Error()))
		return
	}

	task.ID = id
	v, err := c.service.Update(task)

	if err != nil {
		ctx.JSON(http.StatusNotFound, util.ErrorsPayload(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.ResourcePayload(v))
}

func (c *controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := c.service.FindById(id); err != nil {
		ctx.JSON(http.StatusNotFound, util.ErrorsPayload(err.Error()))
	}
	c.service.Delete(id)
	ctx.JSON(http.StatusNoContent, nil)
}
