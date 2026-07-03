package ciserver

import (
	"fmt"
	"strings"

	"github.com/open-beagle/bdpulse-scm/scm"
)

// Paths defines endpoint templates used by the CI-server adapter.
//
// Product-specific routes are supplied by the caller so this public SCM package
// does not own deployment URL decisions.
type Paths struct {
	UserInfo         string
	CreateProject    string
	RepositoryByID   string
	RepositoryBySlug string
	RepositoryPerms  string
	RepositoryList   string
	ContentByID      string
	ContentBySlug    string
	ContentWrite     string
	Commit           string
	Branches         string
	Groups           string
}

// Option configures a CI-server client.
type Option func(*options)

type options struct {
	paths Paths
}

// WithPaths configures endpoint templates for CI-server compatibility APIs.
func WithPaths(paths Paths) Option {
	return func(opts *options) {
		opts.paths = paths
	}
}

func formatPath(tmpl string, args ...interface{}) (string, error) {
	tmpl = strings.TrimLeft(strings.TrimSpace(tmpl), "/")
	if tmpl == "" {
		return "", scm.ErrNotSupported
	}
	return fmt.Sprintf(tmpl, args...), nil
}
