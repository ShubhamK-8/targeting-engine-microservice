package campaignsservice

import (
	"errors"
	coreUtils "targeting-engine/coreUtils"
	dbConnection "targeting-engine/database/elasticSerach"
	webServiceSchema "targeting-engine/webService/schema"
)

func GetCampaignsList(params *webServiceSchema.DeliveryRequest) (campaigns []webServiceSchema.CampaignResponse, err error) {

	esClient, err := dbConnection.NewElasticsearchClient(coreUtils.ElasticsearchHost)
	if err != nil {
		println("Error while makign connection with", err)
		err = errors.New("Internal error occuered")
		return
	}
	campaigns, err = dbConnection.QueryElasticsearch(esClient, params.AppID, params.Country, params.OS)
	return
}
