package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/damejeras/goauth/jwt"
	jwtgo "github.com/dgrijalva/jwt-go"
)

func (s *server) get() http.HandlerFunc {
	key, err := ioutil.ReadFile(s.config.KeyPath)
	if err != nil {
		log.Fatalf("error reading private key: %s\n", err)
	}
	privateKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		log.Fatalf("error parsing private key: %s\n", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, jwt.MakeClaims())

		signed, err := token.SignedString(privateKey)
		if err != nil {
			http.Error(w, "internal error", 500)
			log.Fatalln(err)
			return
		}

		response := jwt.Response{
			Token:      signed,
			Type:       jwt.Bearer,
			Expiration: 3600,
		}

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "internal error", 500)
			log.Fatalln(err)
		}
	}
}
