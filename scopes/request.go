package scopes

type createNewRequest struct {
	Scope string `json:"scope"`
}

func (r createNewRequest) validate() map[string][]string {
	result := make(map[string][]string)
	if r.Scope == "" {
		result["scope"] = []string{"scope must be set"}
	}
	return result
}
