package auth

import (
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name    string
		headers map[string][]string
		want    string
		wantErr bool
	}

	tests := []test{
		{
			name: "valid api key",
			headers: map[string][]string{
				"Authorization": {"ApiKey my-api-key"},
			},
			want:    "my-api-key",
			wantErr: false,
		},
		{
			name:    "no authorization header",
			headers: map[string][]string{},
			want:    "",
			wantErr: true,
		},
		{
			name: "malformed authorization header",
			headers: map[string][]string{
				"Authorization": {"Bearer my-api-key"},
			},
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
