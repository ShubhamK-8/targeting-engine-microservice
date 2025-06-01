package v1

import (
	"net/http"

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

	trackData, err := webServiceHelper.GetCampaignsList(fetchCriteria, params.Verbose, maskAddressFields)
	//zap.L().Info("Tracking Response", zap.Any("trackData", trackData), zap.Any("error", err))
	if err != nil {
		response.SetSuccess(false).SetDescription(err.GetErrorString()).SetError(err)
		request.IndentedJSON(err.GetStatusCode(), response)
		return
	} else {
		response.SetSuccess(true).SetData(trackData)
		request.IndentedJSON(http.StatusOK, response)
		return
	}

}
