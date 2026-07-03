package oauth2

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/open-beagle/bdpulse-scm/scm"
)

func TestRefresherClientCredentials(t *testing.T) {
	tests := []struct {
		name                      string
		clientCredentialsInBody   bool
		expectBasicAuth           bool
		expectCredentialsFormBody bool
	}{
		{
			name:            "basic auth by default",
			expectBasicAuth: true,
		},
		{
			name:                      "credentials in form body",
			clientCredentialsInBody:   true,
			expectCredentialsFormBody: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if err := r.ParseForm(); err != nil {
					t.Fatalf("parse form: %v", err)
				}

				username, password, hasBasicAuth := r.BasicAuth()
				if hasBasicAuth != test.expectBasicAuth {
					t.Fatalf("basic auth present = %v, want %v", hasBasicAuth, test.expectBasicAuth)
				}
				if test.expectBasicAuth && (username != "client-id" || password != "client-secret") {
					t.Fatalf("basic auth = %q/%q, want client-id/client-secret", username, password)
				}

				hasClientID := r.Form.Get("client_id") != ""
				if hasClientID != test.expectCredentialsFormBody {
					t.Fatalf("client_id present = %v, want %v", hasClientID, test.expectCredentialsFormBody)
				}
				if test.expectCredentialsFormBody {
					if got := r.Form.Get("client_id"); got != "client-id" {
						t.Fatalf("client_id = %q, want client-id", got)
					}
					if got := r.Form.Get("client_secret"); got != "client-secret" {
						t.Fatalf("client_secret = %q, want client-secret", got)
					}
				}

				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"access_token":"new-token","refresh_token":"new-refresh","expires_in":3600}`))
			}))
			defer server.Close()

			token := &scm.Token{
				Token:   "old-token",
				Refresh: "old-refresh",
				Expires: time.Now().Add(-time.Hour),
			}
			refresher := &Refresher{
				ClientID:                "client-id",
				ClientSecret:            "client-secret",
				Endpoint:                server.URL,
				ClientCredentialsInBody: test.clientCredentialsInBody,
				Source:                  StaticTokenSource(token),
				Client:                  server.Client(),
			}

			got, err := refresher.Token(context.Background())
			if err != nil {
				t.Fatalf("refresh token: %v", err)
			}
			if got.Token != "new-token" {
				t.Fatalf("token = %q, want new-token", got.Token)
			}
			if got.Refresh != "new-refresh" {
				t.Fatalf("refresh = %q, want new-refresh", got.Refresh)
			}
		})
	}
}
