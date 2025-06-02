package v1

import (
	"net/http"

	campaignService "targeting-engine/service/campaignService"
	webServiceHelper "targeting-engine/webService/helper"
	webServiceSchema "targeting-engine/webService/schema"

	"github.com/gin-gonic/gin"
)

func handleDelivery(request *gin.Context) {
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
