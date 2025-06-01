package campaignsservice

import (
	webServiceSchema "targeting-engine/webService/schema"
	serviceHelper "targeting-engine/service/helper"

)

func GetCampaignsList(params *webServiceSchema.DeliveryRequest) ([]webServiceSchema.CampaignResponse, error) {

	serviceHelper.MatchCampaigns()

}
