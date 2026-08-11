package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		name         string
		sampleHeader http.Header
		expectKey    string
		expectFail   bool
	}{
		{
			name: "Valid Authorization Header",
			sampleHeader: http.Header{
				"Authorization": []string{"ApiKey abc123"},
			},
			expectKey:  "abc123",
			expectFail: false,
		},
		{
			name:         "Empty Authorization Header",
			sampleHeader: http.Header{},
			expectFail:   true,
		},
		{
			name: "Split length less than 2",
			sampleHeader: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectFail: true,
		},
		{
			name: "Key is not 'ApiKey'",
			sampleHeader: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			expectFail: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tc.sampleHeader)

			if tc.expectFail {
				if err == nil {
					t.Fatal("expected GetAPIKey to fail")
				}
				return
			}

			if err != nil {
				t.Fatalf("GetAPIKey returned an unexpected error: %v", err)
			}
			if apiKey != tc.expectKey {
				t.Fatalf("GetAPIKey returned %q, want %q", apiKey, tc.expectKey)
			}
		})
	}
}
