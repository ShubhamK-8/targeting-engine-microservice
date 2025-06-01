package campaignsservice

import (
	serviceHelper "targeting-engine/service/campaignService/helper"
	webServiceSchema "targeting-engine/webService/schema"
)

func GetCampaignsList(params *webServiceSchema.DeliveryRequest) ([]webServiceSchema.CampaignResponse, error) {

	serviceHelper.MatchCampaigns()

}
