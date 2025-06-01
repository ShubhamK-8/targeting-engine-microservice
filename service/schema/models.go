package models

// Campaign represents an advertisement campaign.
type Campaign struct {
	ID       string `json:"cid"`    // Unique identifier for the campaign
	Name     string `json:"name"`   // Name of the campaign
	ImageURL string `json:"img"`    // URL of the image creative
	CTA      string `json:"cta"`    // Call To Action text (e.g., "Download", "Install")
	Status   string `json:"status"` // Current state of the campaign ("ACTIVE" or "INACTIVE")
}

// TargetingRule defines the conditions under which a campaign can be served.
type TargetingRule struct {
	CampaignID     string          `json:"campaign_id"`     // ID of the campaign this rule applies to
	IncludeCountry map[string]bool `json:"include_country"` // Countries to include (map for fast lookup)
	ExcludeCountry map[string]bool `json:"exclude_country"` // Countries to exclude
	IncludeOS      map[string]bool `json:"include_os"`      // Operating Systems to include
	ExcludeOS      map[string]bool `json:"exclude_os"`      // Operating Systems to exclude
	IncludeApp     map[string]bool `json:"include_app"`     // App IDs to include
	ExcludeApp     map[string]bool `json:"exclude_app"`     // App IDs to exclude
}
