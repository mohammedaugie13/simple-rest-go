package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Rest) HealthCheck(c *gin.Context) {
	service := h.Service.HealthCheckService.HealthCheck()
	c.JSON(http.StatusOK, service)
}
