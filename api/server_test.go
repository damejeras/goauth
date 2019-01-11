package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: add table testing
func TestServeToken(t *testing.T) {
	config := Config{"../keys/jwtRS256.key", ""}
	server := NewServer(config)
	req, err := http.NewRequest("GET", "localhost:8080/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()

	server.get()(rec, req)

	res := rec.Result()
	if res.StatusCode != 200 {
		t.Errorf("expected 200, got %v", res.StatusCode)
	}
	result := map[string]interface{}{}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if result["token_type"] != "Bearer" {
		t.Errorf("wanted 'Bearer' token_type, got '%v'", result["token_type"])
	}
}
