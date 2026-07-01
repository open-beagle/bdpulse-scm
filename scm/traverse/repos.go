package traverse

import (
	"context"

	"github.com/open-beagle/bdpulse-scm/scm"
)

// Repos returns the full repository list, traversing and
// combining paginated responses if necessary.
func Repos(ctx context.Context, client *scm.Client) ([]*scm.Repository, error) {
	list := []*scm.Repository{}
	opts := scm.ListOptions{Size: 100}
	for {
		result, meta, err := client.Repositories.List(ctx, opts)
		if err != nil {
			return nil, err
		}
		for _, src := range result {
			if src != nil {
				list = append(list, src)
			}
		}
		opts.Page = meta.Page.Next
		opts.URL = meta.Page.NextURL

		if opts.Page == 0 && opts.URL == "" {
			break
		}
	}
	return list, nil
}
