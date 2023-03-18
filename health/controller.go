package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController interface {
	Health(ctx *gin.Context)
}

type controller struct{}

func NewController() HealthController {
	return &controller{}
}

func (*controller) Health(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}

func Route(r *gin.Engine, c HealthController) {
	r.GET("/health", c.Health)
}
