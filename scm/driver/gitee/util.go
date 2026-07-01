package gitee

import (
	"net/url"
	"strconv"

	"github.com/open-beagle/bdpulse-scm/scm"
)

func encodeListOptions(opts scm.ListOptions) string {
	params := url.Values{}
	if opts.Page != 0 {
		params.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.Size != 0 {
		params.Set("per_page", strconv.Itoa(opts.Size))
	}
	return params.Encode()
}

func encodeCommitListOptions(opts scm.CommitListOptions) string {
	params := url.Values{}
	if opts.Page != 0 {
		params.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.Size != 0 {
		params.Set("per_page", strconv.Itoa(opts.Size))
	}
	if opts.Ref != "" {
		params.Set("sha", opts.Ref)
	}
	return params.Encode()
}

func encodeIssueListOptions(opts scm.IssueListOptions) string {
	params := url.Values{}
	if opts.Page != 0 {
		params.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.Size != 0 {
		params.Set("per_page", strconv.Itoa(opts.Size))
	}
	// state: open, progressing, closed, rejected, all; default: open
	if opts.Open && opts.Closed {
		params.Set("state", "all")
	} else if opts.Closed {
		params.Set("state", "closed")
	} else if opts.Open {
		params.Set("state", "open")
	}
	return params.Encode()
}

func encodePullRequestListOptions(opts scm.PullRequestListOptions) string {
	params := url.Values{}
	if opts.Page != 0 {
		params.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.Size != 0 {
		params.Set("per_page", strconv.Itoa(opts.Size))
	}
	// state: open, closed, merged, all; default: open
	if opts.Open && opts.Closed {
		params.Set("state", "all")
	} else if opts.Closed {
		params.Set("state", "closed")
	} else if opts.Open {
		params.Set("state", "open")
	}
	return params.Encode()
}
