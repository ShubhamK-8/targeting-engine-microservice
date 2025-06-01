package helper

import (
	"errors"
	webServiceSchema "targeting-engine/webService/schema"

	"github.com/gin-gonic/gin"
)

func FetchRequestParams(request *gin.Context) *webServiceSchema.DeliveryRequest {
	params := &webServiceSchema.DeliveryRequest{
		AppID:   request.Query("app"),
		Country: request.Query("country"),
		OS:      request.Query("os"),
	}
	return params
}

// ValidateDeliveryRequest checks if all required fields are present
func ValidateRequest(params *webServiceSchema.DeliveryRequest) error {
	if params.AppID == "" || params.OS == "" || params.Country == "" {
		return errors.New("missing one or more required parameters: app, os, country")
	}
	return nil
}
