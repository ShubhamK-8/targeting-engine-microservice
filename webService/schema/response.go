package schema

// CampaignResponse represents the simplified campaign details sent back to the client.
type CampaignResponse struct {
	CID string `json:"cid"` // Campaign ID
	Img string `json:"img"` // Image URL
	CTA string `json:"cta"` // Call To Action
}
