package scm

import "context"

// Linker provides deep links to resources.
type Linker interface {
	// Resource returns a link to the resource.
	Resource(ctx context.Context, repo string, ref Reference) (string, error)

	// Diff returns a link to the diff.
	Diff(ctx context.Context, repo string, source, target Reference) (string, error)
}
