package rules

import "targeting-engine/internal/models"

var Campaigns = []models.Campaign{
	{"spotify", "https://somelink", "Download", "ACTIVE"},
	{"duolingo", "https://somelink2", "Install", "ACTIVE"},
	{"subwaysurfer", "https://somelink3", "Play", "ACTIVE"},
}

var Rules = []models.TargetingRule{
	{CampaignID: "spotify", IncludeCountry: []string{"US", "Canada"}},
	{CampaignID: "duolingo", IncludeOS: []string{"Android", "iOS"}, ExcludeCountry: []string{"US"}},
	{CampaignID: "subwaysurfer", IncludeOS: []string{"Android"}, IncludeAppIDs: []string{"com.gametion.ludokinggame"}},
}
