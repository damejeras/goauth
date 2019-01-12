package scopes

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleNew returns a handler to handle Scope creation requests
func HandleNew() http.HandlerFunc {
	db, err := sql.Open("mysql", "root:my-secret-pw@/goauth")
	if err != nil {
		log.Fatalf("could not establish database connection: %v", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		vars := mux.Vars(r)
		email, ok := vars["email"]
		if !ok {
			log.Fatalln("no `email` in URL for HandleNew() handler")
		}

		var id int64
		row := db.QueryRow("SELECT id FROM users WHERE email = ?", email)
		err := row.Scan(&id)
		if err == sql.ErrNoRows {
			w.WriteHeader(404)
			err := encoder.Encode(map[string]string{
				"error": "email not found",
			})
			if err != nil {
				panic(err)
			}
			return
		}
		var req createNewRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// TODO make response format consistent
			http.Error(w, "bad request", 400)
			return
		}

		errors := req.validate()
		if len(errors) > 0 {
			w.WriteHeader(400)
			err := encoder.Encode(map[string]interface{}{
				"errors": errors,
			})
			if err != nil {
				panic(err)
			}
			return
		}

		stmt, err := db.Prepare("INSERT INTO scopes (user_id, scope) VALUES (?, ?)")
		if err != nil {
			http.Error(w, "internal error", 500)
			log.Printf("could not create insert statement: %v", err)
			return
		}

		_, err = stmt.Exec(id, req.Scope)
		if err != nil {
			http.Error(w, "internal error", 500)
			log.Printf("error while inserting scope: %v", err)
			return
		}

		w.WriteHeader(201)
		encoder.Encode(map[string]string{
			"message": "scope created",
		})
	}
}
