package enrich

import (
	"context"

	"github.com/open-beagle/go-scm/scm"
)

// Webhook enriches the webhook payload with missing
// information not included in the webhook payload.
func Webhook(ctx context.Context, client *scm.Client, webhook *scm.Webhook) error {
	return nil // TODO
}
