package jwt

import (
	"net/http"
	"testing"
)

func TestHello(t *testing.T) {
	want := struct {
		Token string
		OK    bool
	}{
		Token: "example",
		OK:    true,
	}

	req, err := http.NewRequest("GET", "testing", nil)
	if err != nil {
		t.Fatalf("could not create request for test %v", err)
	}
	req.Header.Set("authorization", "Bearer example")
	authenticator := &JWTAuthenticator{}
	token, ok := authenticator.GetTokenFromRequest(req)

	if token == "" || !ok {
		t.Fatalf(" want token:%q,ok: %t | got: %q, %t", want.Token, want.OK, "", false)
	}
}
