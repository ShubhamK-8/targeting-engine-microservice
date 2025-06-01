package campaignsservice

import (
	db "targeting-engine/connection"
	serviceHelper "targeting-engine/service/campaignService/helper"
	webServiceSchema "targeting-engine/webService/schema"
)

func GetCampaignsList(params *webServiceSchema.DeliveryRequest) ([]webServiceSchema.CampaignResponse, error) {

	campaigns := serviceHelper.MatchCampaigns(*params, db.Campaigns, db.Rules)
	return campaigns, nil

}
