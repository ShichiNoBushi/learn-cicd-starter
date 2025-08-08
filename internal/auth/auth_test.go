package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey1(t *testing.T) {
	testHeaders := make(http.Header)

	testHeaders.Set("Authorization", "ApiKey testing")

	got, err := GetAPIKey(testHeaders)
	want := "testing"

	if err != nil {
		t.Fatalf("expected: nil error, got: %s", err)
	}

	if got != want {
		t.Fatalf("expected: %s, got: %s", want, got)
	}
}

func TestGetAPIKey2(t *testing.T) {
	testHeaders := make(http.Header)

	testHeaders.Set("Authorization", "")

	_, err := GetAPIKey(testHeaders)

	if err == nil {
		t.Fatalf("expected: an error, got: %s", err)
	}

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected: no authorization error, got: %s", err)
	}
}

func TestGetAPIKey3(t *testing.T) {
	testHeaders := make(http.Header)

	testHeaders.Set("Authorization", "incorrect")

	_, err := GetAPIKey(testHeaders)

	if err == nil {
		t.Fatalf("expected: an error, got: %s", err)
	}

	if err.Error() != "malformed authorization header" {
		t.Fatalf("expected: invalid authorization error, got: %s", err)
	}
}
