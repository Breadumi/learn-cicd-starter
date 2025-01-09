package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	test_req, _ := http.NewRequest("GET", "throwaway.com", nil)

	test_req.Header.Set("Blah", "nonsense")
	resp, err := GetAPIKey(test_req.Header)
	if err == nil || resp != "" {
		t.Fatalf("expected: %v, %v\n, got: %v, %v", "", errors.New("no authorization header included"), resp, err)
	}

	test_key := "oausdf897sdfh134jh4kriu5"
	test_req.Header.Set("Authorization", "ApiKey "+test_key)
	resp, err = GetAPIKey(test_req.Header)
	if resp != test_key || err != nil {
		t.Fatalf("expected: %v, %v\n, got: %v, %v", test_key, nil, resp, err)
	}

}
