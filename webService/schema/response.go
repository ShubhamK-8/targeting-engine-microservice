package schema

// CampaignResponse represents the simplified campaign details sent back to the client.
type CampaignResponse struct {
	CID string `json:"cid"` // Campaign ID
	Img string `json:"img"` // Image URL
	CTA string `json:"cta"` // Call To Action
}

type ResponseEntity struct {
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success,omitempty"`
}

func (ResponseEntity *ResponseEntity) SetError(err error) *ResponseEntity {
	ResponseEntity.Error = err.Error()
	return ResponseEntity
}

func (ResponseEntity *ResponseEntity) SetData(data interface{}) *ResponseEntity {
	ResponseEntity.Data = data
	return ResponseEntity
}

func (ResponseEntity *ResponseEntity) SetSuccess(success bool) *ResponseEntity {
	ResponseEntity.Success = success
	return ResponseEntity
}
