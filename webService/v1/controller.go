package v1

import (
	"fmt"
	"net/http"
	"time"

	prometheousInit "targeting-engine/init/prometheous"
	campaignService "targeting-engine/service/campaignService"
	webServiceHelper "targeting-engine/webService/helper"
	webServiceSchema "targeting-engine/webService/schema"

	"github.com/gin-gonic/gin"
)

func handleDelivery(request *gin.Context) {
	start := time.Now()
	var statusCode int // Declare statusCode here

	defer func() {
		// Capture the final status code from the Gin context's writer
		statusCode = request.Writer.Status()
		// Increment total HTTP requests counter
		prometheousInit.HttpRequestsTotal.WithLabelValues(request.Request.URL.Path, request.Request.Method, fmt.Sprintf("%d", statusCode)).Inc()
		// Observe HTTP request duration
		prometheousInit.HttpRequestDuration.WithLabelValues(request.Request.URL.Path, request.Request.Method).Observe(time.Since(start).Seconds())
	}()

	response := &webServiceSchema.ResponseEntity{}
	params := webServiceHelper.FetchRequestParams(request)
	err := webServiceHelper.ValidateRequest(params)

	if err != nil {
		response.SetError(err)
		request.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	matchedCampaigns, err := campaignService.GetCampaignsList(params)

	if err != nil {
		response.SetError(err)
		request.IndentedJSON(http.StatusInternalServerError, response)
		return
	}

	if len(matchedCampaigns) == 0 {
		// No campaigns match, send HTTP 204 No Content.
		request.IndentedJSON(http.StatusNoContent, response)
		return
	}

	response.SetSuccess(true).SetData(matchedCampaigns)
	request.IndentedJSON(http.StatusOK, response)
}
