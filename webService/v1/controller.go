package v1

import (
	"net/http"

	webServiceHelper "targeting-engine/webService/helper"
	webServiceSchema "targeting-engine/webService/schema"

	"github.com/gin-gonic/gin"
)

func handleDelivery(request *gin.Context) {
	response := &webServiceSchema.CampaignResponse{}
	params := webServiceHelper.FetchRequestParams(request)
	


	trackData, err := webServiceHelper.GetPackageStatus(fetchCriteria, params.Verbose, maskAddressFields)
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
