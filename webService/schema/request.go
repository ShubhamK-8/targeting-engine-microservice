package schema

// DeliveryRequest represents an incoming request from an end-user/app.
type DeliveryRequest struct {
	AppID   string `json:"app_id"`  // Unique identifier of the app/game
	Country string `json:"country"` // Country of the user
	OS      string `json:"os"`      // Operating System (e.g., "Android", "iOS", "Web")
}
