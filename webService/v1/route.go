package v1

import (
	"github.com/gin-gonic/gin"
)

// AddRoutes Routes request to its request controller
func AddRoutes(router *gin.RouterGroup) {
	router.GET("/v1/delivery", handleDelivery)
}
