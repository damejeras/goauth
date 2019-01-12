package scopes

type createNewRequest struct {
	Scope string `json:"scope"`
}

func (r createNewRequest) validate() map[string][]string {
	result := make(map[string][]string)
	return result
}
