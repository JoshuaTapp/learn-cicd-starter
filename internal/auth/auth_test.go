package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectedErr error
	}{
		{
			name:        "No Authorization Header",
			headers:     http.Header{},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Authorization Header 1",
			headers: http.Header{
				"Authorization": []string{"InvalidHeader"},
			},
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
		{
			name: "Malformed Authorization Header 2",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
		{
			name: "Valid Authorization Header",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-api-key"},
			},
			expectedKey: "my-secret-api-key",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)
			if apiKey != tt.expectedKey {
				t.Errorf("expected %v, got %v", tt.expectedKey, apiKey)
			}
			if (err != nil && tt.expectedErr == nil) || (err == nil && tt.expectedErr != nil) {
				t.Errorf("expected error %v, got error %v", tt.expectedErr, err)
			}
			if err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("expected error %v, got error %v", tt.expectedErr, err)
			}
		})
	}
}
