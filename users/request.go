package users

type registrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *registrationRequest) validate() map[string][]string {
	errors := make(map[string][]string)
	if len(r.Password) < 6 {
		errors["password"] = append(errors["passwords"], "password must be atleast 6 symbols long")
	}
	return errors
}
