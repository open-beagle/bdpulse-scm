package beagle

import (
	"fmt"
	"strings"

	"github.com/open-beagle/bdpulse-scm/scm"
)

// Paths defines the Beagle API endpoints used by this adapter.
//
// The adapter intentionally does not provide product-specific defaults. Callers
// that need compatibility endpoints must inject them at construction time.
type Paths struct {
	UserInfo       string
	Netrc          string
	Content        string
	Commit         string
	Repository     string
	RepositoryList string
}

// Option configures a Beagle client.
type Option func(*options)

type options struct {
	paths Paths
}

// WithPaths configures endpoint templates for Beagle compatibility APIs.
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
