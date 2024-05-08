package tests

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func Test_GetAPIKey(t *testing.T) {
	tests := map[string]struct {
		header http.Header
		key    string
		err    error
	}{
		"NoAuthHeader":           {header: http.Header{}, key: "", err: errors.New("no authorization header included")},
		"AuthHeaderWithDiffName": {header: http.Header{"Another": []string{"ApiKey 1234"}}, key: "", err: errors.New("no authorization header included")},
		"NoToken":                {header: http.Header{"Authorization": []string{"ApiKey"}}, key: "", err: errors.New("malformed authorization header")},
		"AnotherName":            {header: http.Header{"Authorization": []string{"Bearer 1234"}}, key: "", err: errors.New("malformed authorization header")},
		"Ok":                     {header: http.Header{"Authorization": []string{"ApiKey 1234567"}}, key: "1234567", err: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := auth.GetAPIKey(tc.header)

			if output != tc.key {
				t.Errorf("%v : expected -> %v -- got -> %v", name, tc.key, output)
			}

			// Check error handling
			if (err != nil && tc.err == nil) || (err == nil && tc.err != nil) {
				t.Errorf("%v: Expected error %v, got %v", name, tc.err, err)
			}

			// If there's an error, ensure it matches the expected error
			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("%v: Expected error message %v, got %v", name, tc.err.Error(), err.Error())
			}
		})
	}
}
