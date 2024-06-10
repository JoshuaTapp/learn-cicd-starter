package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRespondWithError(t *testing.T) {
	tests := []struct {
		code            int
		message         string
		expectedLog     string
		expectedBody    string
		expectedHeaders map[string]string
	}{
		{
			code:            400,
			message:         "Bad Request",
			expectedBody:    `{"error":"Bad Request"}`,
			expectedHeaders: map[string]string{"Content-Type": "application/json"},
		},
		{
			code:            500,
			message:         "Internal Server Error",
			expectedLog:     "Responding with 5XX error: Internal Server Error",
			expectedBody:    `{"error":"Internal Server Error"}`,
			expectedHeaders: map[string]string{"Content-Type": "application/json"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.message, func(t *testing.T) {
			rr := httptest.NewRecorder()

			respondWithError(rr, tt.code, tt.message)

			// Verify Response Code
			assert.Equal(t, tt.code, rr.Code)

			// Verify Headers
			for key, val := range tt.expectedHeaders {
				assert.Equal(t, val, rr.Header().Get(key))
			}

			// Verify Body
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestRespondWithJSON(t *testing.T) {
	type response struct {
		Message string `json:"message"`
	}

	tests := []struct {
		code            int
		payload         interface{}
		expectedBody    string
		expectedHeaders map[string]string
		shouldError     bool
	}{
		{
			code:            200,
			payload:         response{Message: "Success"},
			expectedBody:    `{"message":"Success"}`,
			expectedHeaders: map[string]string{"Content-Type": "application/json"},
		},
		{
			code:            500,
			payload:         make(chan int), // Non-marshallable payload to trigger error
			expectedHeaders: map[string]string{"Content-Type": "application/json"},
			shouldError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.expectedBody, func(t *testing.T) {
			rr := httptest.NewRecorder()

			respondWithJSON(rr, tt.code, tt.payload)

			// Verify Response Code
			if tt.shouldError {
				assert.Equal(t, http.StatusInternalServerError, rr.Code)
			} else {
				assert.Equal(t, tt.code, rr.Code)
			}

			// Verify Headers
			for key, val := range tt.expectedHeaders {
				assert.Equal(t, val, rr.Header().Get(key))
			}

			// Verify Body
			if !tt.shouldError {
				assert.JSONEq(t, tt.expectedBody, rr.Body.String())
			}
		})
	}
}
