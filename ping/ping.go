package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingResponse struct {
	Message string `json:"message"`
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, PingResponse{
		Message: "pong",
	})
}
