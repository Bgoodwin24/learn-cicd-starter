package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		request     *http.Request
		expectedKey string
		expectError bool
	}{
		{
			name:        "Valid API key in Authorization header",
			request:     createRequestWithHeader("Authorization", "ApiKey my-api-key"),
			expectedKey: "my-api-key",
			expectError: false,
		},
		{
			name:        "Missing Authorization header",
			request:     createRequestWithHeader("", ""),
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "Malformed Authorization header",
			request:     createRequestWithHeader("Authorization", "Bearer my-api-key"),
			expectedKey: "",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			key, err := GetAPIKey(tc.request.Header)

			if tc.expectError && err == nil {
				t.Fatalf("expected error but got none")
			}
			if !tc.expectError && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.expectError && key != tc.expectedKey {
				t.Fatalf("expected key %q, got %q", tc.expectedKey, key)
			}
		})
	}
}

func createRequestWithHeader(key, value string) *http.Request {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	if key != "" {
		req.Header.Add(key, value)
	}
	return req
}
