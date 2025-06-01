package models

type Campaign struct {
    ID     string   `json:"cid"`
    Image  string   `json:"img"`
    CTA    string   `json:"cta"`
    Status string   `json:"status"` // ACTIVE or INACTIVE
}

type TargetingRule struct {
    CampaignID     string
    IncludeAppIDs  []string
    IncludeOS      []string
    IncludeCountry []string
    ExcludeAppIDs  []string
    ExcludeOS      []string
    ExcludeCountry []string
}
