package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyValid(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey valid_key")
	key, err := GetAPIKey(h)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	if key != "valid_key" {
		t.Errorf("expected key to be valid_key, got %s", key)
	}
}

func TestGetAPIEmptyAuth(t *testing.T) {
	h := http.Header{}
	key, err := GetAPIKey(h)
	if err == nil {
		t.Errorf("expected error, got none")
	}
	if key != "" {
		t.Errorf("expected empty key, got %s", key)
	}
}

func TestGetAPIKeyMalformed(t *testing.T) {
	h := http.Header{}
	key, err := GetAPIKey(h)
	h.Set("Authorization", "bearer valid_key")

	if err == nil {
		t.Errorf("expected error, got none")
	}
	if key != "" {
		t.Errorf("expected empty key, got %s", key)
	}
}
