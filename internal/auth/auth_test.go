package auth_test

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		headers http.Header
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Happy path",
			headers: http.Header{
				"Authorization": []string{"ApiKey 1234"},
			},
			want: "1234",
		},
		{
			name: "Ignores extra contents",
			headers: http.Header{
				"Authorization": []string{"ApiKey 4567 extra"},
			},
			want: "4567",
		},
		{
			name: "Ignores extra values",
			headers: http.Header{
				"Authorization": []string{"ApiKey 4567", "extra value"},
			},
			want: "4567",
		},

		{
			name:    "No header",
			headers: http.Header{},
			wantErr: true,
		},
		{
			name: "Empty header",
			headers: http.Header{
				"Authorization": []string{""},
			},
			wantErr: true,
		},
		{
			name: "Malformed header",
			headers: http.Header{
				"Authorization": []string{"Bearer 1234"},
			},
			wantErr: true,
		},
		{
			name: "Multiple values",
			headers: http.Header{
				"Authorization": []string{"Bearer 1234", "ApiKey 1234"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := auth.GetAPIKey(tt.headers)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetAPIKey() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetAPIKey() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
