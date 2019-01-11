package users

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// RegistrationHandler returns a handler to handle requests for user registration
func RegistrationHandler() http.HandlerFunc {
	db, err := sql.Open("mysql", "root:my-secret-pw@/goauth")
	if err != nil {
		log.Fatalf("could not establish database connection: %v", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		req := registrationRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "bad request", 400)
			return
		}

		errors := req.validate()
		if len(errors) > 0 {
			w.WriteHeader(422)
			response := map[string]interface{}{
				"errors": errors,
			}
			err := encoder.Encode(response)
			if err != nil {
				log.Fatalf("could not encode response: %v", err)
			}
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
		if err != nil {
			log.Fatalf("could not hash a password: %v", err)
		}

		stmt, err := db.Prepare("INSERT INTO users (email, password) VALUES (?, ?)")
		if err != nil {
			log.Fatalf("could not create insert statement: %v", err)
		}

		_, err = stmt.Exec(req.Email, hash)
		if err != nil {
			if driverErr, ok := err.(*mysql.MySQLError); ok { // Now the error number is accessible directly
				if driverErr.Number == 1062 {
					http.Error(w, "user already exists", 409)
					return
				}
			}
			log.Fatalf("could not insert user: %v", err)
		}

		resp := map[string]string{
			"message": "account created",
		}
		err = encoder.Encode(resp)
		if err != nil {
			log.Fatalf("could not decode response: %v", err)
		}
	}
}

// LoginHandler returns a handler to handle login requests
func LoginHandler() http.HandlerFunc {
	db, err := sql.Open("mysql", "root:my-secret-pw@/goauth")
	if err != nil {
		log.Fatalf("could not establish database connection: %v", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		req := registrationRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "bad request", 400)
			return
		}

		user := User{}
		row := db.QueryRow("SELECT * FROM users WHERE email = ?", req.Email)
		err = row.Scan(&user.ID, &user.Email, &user.Password)
		if err == sql.ErrNoRows {
			w.WriteHeader(401)
			res := map[string]string{
				"error": "invalid credentials",
			}
			encoder.Encode(res)
		} else if err != nil {
			http.Error(w, "internal error", 500)
			log.Fatalf("error while searching for user: %v", err)
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			w.WriteHeader(401)
			res := map[string]string{
				"error": "invalid credentials",
			}
			encoder.Encode(res)
			return
		}

		err = encoder.Encode(map[string]string{
			"message": "logged in",
		})
		if err != nil {
			http.Error(w, "internal error", 500)
			log.Fatalf("error while encoding response: %v", err)
		}
	}
}
