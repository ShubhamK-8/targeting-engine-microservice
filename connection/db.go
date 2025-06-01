package connection

import (
	serviceSchema "targeting-engine/service/schema"
)

// Predefined set of campaigns (some active, some inactive for test coverage)
var Campaigns = []serviceSchema.Campaign{
	{
		ID:       "spotify",
		Name:     "Spotify - Music for everyone",
		ImageURL: "https://somelink",
		CTA:      "Download",
		Status:   "ACTIVE",
	},
	{
		ID:       "duolingo",
		Name:     "Duolingo: Best way to learn",
		ImageURL: "https://somelink2",
		CTA:      "Install",
		Status:   "ACTIVE",
	},
	{
		ID:       "subwaysurfer",
		Name:     "Subway Surfer",
		ImageURL: "https://somelink3",
		CTA:      "Play",
		Status:   "ACTIVE",
	},
	{
		ID:       "inactive_test",
		Name:     "Inactive Campaign",
		ImageURL: "https://inactive.link",
		CTA:      "Inactive",
		Status:   "INACTIVE",
	},
}

// Targeting rules defining which requests are eligible for which campaigns
var Rules = map[string]serviceSchema.TargetingRule{
	"spotify": {
		CampaignID:     "spotify",
		IncludeCountry: map[string]bool{"US": true, "Canada": true},
	},
	"duolingo": {
		CampaignID:     "duolingo",
		IncludeOS:      map[string]bool{"Android": true, "iOS": true},
		ExcludeCountry: map[string]bool{"US": true},
	},
	"subwaysurfer": {
		CampaignID: "subwaysurfer",
		IncludeOS:  map[string]bool{"Android": true},
		IncludeApp: map[string]bool{"com.gametion.ludokinggame": true},
	},
}
