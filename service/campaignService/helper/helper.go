package helper

import (
	serviceSchema "targeting-engine/service/schema"
	webServiceSchema "targeting-engine/webService/schema"
)

// MatchCampaigns takes DeliveryRequest and returns a list of matching CampaignResponse objects.
func MatchCampaigns(req webServiceSchema.DeliveryRequest) []webServiceSchema.CampaignResponse {

	var matched []webServiceSchema.CampaignResponse

	// Iterate through all active campaigns.
	for _, campaign := range s.campaigns {
		// Only consider active campaigns.
		if campaign.Status != "ACTIVE" {
			continue
		}

		// Get the targeting rule for the current campaign.
		rule, exists := s.rules[campaign.ID]
		if !exists {
			// If no rule exists, it means the campaign has no targeting restrictions,ss
			matched = append(matched, webServiceSchema.CampaignResponse{
				CID: campaign.ID,
				Img: campaign.ImageURL,
				CTA: campaign.CTA,
			})
			continue
		}

		// Evaluate targeting rules.
		if evaluateRule(req, rule) {
			matched = append(matched, webServiceSchema.CampaignResponse{
				CID: campaign.ID,
				Img: campaign.ImageURL,
				CTA: campaign.CTA,
			})
		}
	}

	return matched
}

// evaluateRule if a given DeliveryRequest matches a TargetingRule.
func evaluateRule(req webServiceSchema.DeliveryRequest, rule serviceSchema.TargetingRule) bool {
	// Handle Exclude rules first. If any exclude rule matches, the campaign does not qualify.

	// Exclude Country
	if len(rule.ExcludeCountry) > 0 && rule.ExcludeCountry[req.Country] {
		return false
	}
	// Exclude OS
	if len(rule.ExcludeOS) > 0 && rule.ExcludeOS[req.OS] {
		return false
	}
	// Exclude App
	if len(rule.ExcludeApp) > 0 && rule.ExcludeApp[req.AppID] {
		return false
	}

	// Handle Include rules. If any include rule exists for a dimension,

	// Include Country
	if len(rule.IncludeCountry) > 0 && !rule.IncludeCountry[req.Country] {
		return false
	}
	// Include OS
	if len(rule.IncludeOS) > 0 && !rule.IncludeOS[req.OS] {
		return false
	}
	// Include App
	if len(rule.IncludeApp) > 0 && !rule.IncludeApp[req.AppID] {
		return false
	}

	return true
}
