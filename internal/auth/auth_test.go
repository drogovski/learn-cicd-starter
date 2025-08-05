package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name    string
		input   http.Header
		want    string
		wantErr error
	}

	headerWithAPIKey := http.Header{"Authorization": []string{"ApiKey JP2137test"}}
	expectedAPIKey := "JP2137test"
	headerWithEmptyAPIKey := http.Header{"Authorization": []string{""}}
	headerWithWrongAPIKey := http.Header{"Authorization": []string{"sadsafdsf3223423dasd32"}}
	headerWithoutAuthorization := http.Header{}

	tests := []test{
		{name: "Header with proper api key", input: headerWithAPIKey, want: expectedAPIKey, wantErr: nil},
		{name: "Header with empty api key", input: headerWithEmptyAPIKey, want: "", wantErr: ErrNoAuthHeaderIncluded},
		{name: "Header with wrong api key", input: headerWithWrongAPIKey, want: "", wantErr: ErrMalformedAuthorizationHeader},
		{name: "Header without authorization", input: headerWithoutAuthorization, want: "", wantErr: ErrNoAuthHeaderIncluded},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if (err != nil) != (tc.wantErr != nil) || !errors.Is(err, tc.wantErr) {
				t.Errorf("expected error %v, got %v", tc.wantErr, err)
			}
			if got != tc.want {
				t.Errorf("expected value = %v, got = %v", tc.want, got)
			}
		})
	}
}
