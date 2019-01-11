package jwt

// Type is used to declare how token should be carried
type Type string

const (
	// Bearer type is when a token carried in Authorization header
	Bearer Type = "Bearer"
)

// Response is used to encode and decode responses with JWT data
type Response struct {
	Token      string `json:"access_token"`
	Type       Type   `json:"token_type"`
	Expiration int64  `json:"expires_in"`
}
