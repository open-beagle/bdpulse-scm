package beagle

import (
	"net/url"
	"strconv"

	"github.com/open-beagle/go-scm/scm"
)

func encodeMemberListOptions(opts scm.ListOptions) string {
	params := url.Values{}
	// params.Set("membership", "true")
	if opts.Page != 0 {
		params.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.Size != 0 {
		params.Set("per_page", strconv.Itoa(opts.Size))
	}
	if len(opts.URL) > 0 {
		params.Set("scm", opts.URL)
	}
	return params.Encode()
}
