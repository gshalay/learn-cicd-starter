package auth

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey_NoHeader(t *testing.T) {
	emptyHeader := http.Header{}

	_, err := auth.GetAPIKey(emptyHeader)

	if err == nil {
		t.Fatalf("expected header error, but got none")
	}
}

func TestGetAPIKey_BlankHeader(t *testing.T) {
	blankAuthHeader := http.Header{}
	blankAuthHeader.Set("Authorization", " ")
	_, err := auth.GetAPIKey(blankAuthHeader)

	if err == nil {
		t.Fatalf("expected empty header error, but got none")
	}
}

func TestGetAPIKey_NonApiKeyHeader(t *testing.T) {
	apiHeader := http.Header{}
	apiHeader.Set("Authorization", "Bunk CaptainBunk420")
	_, err := auth.GetAPIKey(apiHeader)

	if err == nil {
		t.Fatalf("expected malformed header error, but got none")
	}
}

func TestGetAPIKey_ValidHeader(t *testing.T) {
	apiHeader := http.Header{}
	apiHeader.Set("Authorization", "ApiKey SomeApiKey")
	key, err := auth.GetAPIKey(apiHeader)

	if err != nil {
		t.Fatalf("expected no err but got: %v", err)
	} else if key != "SomeApiKey" {
		t.Fatalf("expected key but got SomeApiKey but got %v", key)
	}
}
