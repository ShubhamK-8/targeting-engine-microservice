package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	coreUtils "targeting-engine/coreUtils"
	webService "targeting-engine/webService/v1"
)

var principalRouter *gin.Engine

// InitEnvironment Initialises Gin
func InitEnvironment() {
	principalRouter = gin.New()
	registerRoutes(principalRouter)
	principalRouter.Run(coreUtils.ServerPort)
}

// RegisterRoutes Registers the application routes according to basepath.
func registerRoutes(principalRouter *gin.Engine) {
	servieRoutes := principalRouter.Group(coreUtils.Basepath)
	healthCheckRoutes := principalRouter.Group(coreUtils.HealthCheckBasepath)
	webService.AddRoutes(servieRoutes)
	fmt.Println("Added Routes: for Targeting Engine Service, Basepath", coreUtils.Basepath)
	healthCheckRoutes.GET("v1/check", healthCheck)
}

func healthCheck(request *gin.Context) {
	request.IndentedJSON(http.StatusOK, gin.H{"message": "service is up"})
	return
}
