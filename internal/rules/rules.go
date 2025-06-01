package rules

import (
	"strings"
	"targeting-engine/internal/models"
)

var Campaigns = []models.Campaign{
	{"spotify", "Spotify - Music for everyone", "https://somelink", "Download", "ACTIVE"},
	{"duolingo", "Duolingo: Best way to learn", "https://somelink2", "Install", "ACTIVE"},
	{"subwaysurfer", "Subway Surfer", "https://somelink3", "Play", "ACTIVE"},
}

var TargetingRules = []models.TargetingRule{
	{"spotify", nil, []string{"US", "Canada"}, nil, nil, nil},
	{"duolingo", nil, nil, []string{"US"}, []string{"Android", "iOS"}, nil},
	{"subwaysurfer", []string{"com.gametion.ludokinggame"}, nil, nil, []string{"Android"}, nil},
}

func MatchCampaigns(app, country, os string) []models.Campaign {
	var result []models.Campaign

	for _, campaign := range Campaigns {
		if strings.ToUpper(campaign.Status) != "ACTIVE" {
			continue
		}

		for _, rule := range TargetingRules {
			if rule.CampaignID != campaign.ID {
				continue
			}

			if !matchDimension(rule.IncludeApp, app, true) {
				continue
			}
			if !matchDimension(rule.IncludeCountry, country, true) {
				continue
			}
			if !matchDimension(rule.ExcludeCountry, country, false) {
				continue
			}
			if !matchDimension(rule.IncludeOS, os, true) {
				continue
			}
			if !matchDimension(rule.ExcludeOS, os, false) {
				continue
			}

			result = append(result, campaign)
			break
		}
	}

	return result
}

func matchDimension(values []string, value string, include bool) bool {
	if values == nil || len(values) == 0 {
		return true
	}
	found := false
	for _, v := range values {
		if strings.EqualFold(v, value) {
			found = true
			break
		}
	}
	if include {
		return found
	}
	return !found
}
