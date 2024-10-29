package jwt

import (
	"net/http"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		header string
		want   string
		ok     bool
	}{
		{
			header: "Bearer example",
			want:   "example",
			ok:     true,
		},
		{
			header: "example",
			want:   "",
			ok:     false,
		},
		{
			header: "Bearerexample",
			want:   "",
			ok:     false,
		},
		{
			header: "",
			want:   "",
			ok:     false,
		},
	}

	req, err := http.NewRequest("GET", "testing", nil)
	if err != nil {
		t.Fatalf("could not create request for test %v", err)
	}

	for _, test := range tests {
		req.Header.Set("authorization", test.header)
		token, _ := GetTokenFromRequest(req)

		if token != test.want {
			t.Errorf(`token=%q. Want %q`, token, test.want)
		}
	}
}
