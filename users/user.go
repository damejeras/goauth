package users

// User carries logged in user information
type User struct {
	ID       int64
	Email    string
	Password string
}
