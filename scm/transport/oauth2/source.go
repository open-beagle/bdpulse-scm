package oauth2

import (
	"context"

	"github.com/open-beagle/bdpulse-scm/scm"
)

// StaticTokenSource returns a TokenSource that always
// returns the same token. Because the provided token t
// is never refreshed, StaticTokenSource is only useful
// for tokens that never expire.
func StaticTokenSource(t *scm.Token) scm.TokenSource {
	return staticTokenSource{t}
}

type staticTokenSource struct {
	token *scm.Token
}

func (s staticTokenSource) Token(context.Context) (*scm.Token, error) {
	return s.token, nil
}

// ContextTokenSource returns a TokenSource that returns
// a token from the http.Request context.
func ContextTokenSource() scm.TokenSource {
	return contextTokenSource{}
}

type contextTokenSource struct {
}

func (s contextTokenSource) Token(ctx context.Context) (*scm.Token, error) {
	token, _ := ctx.Value(scm.TokenKey{}).(*scm.Token)
	return token, nil
}
