// Package gitlab implements a GitLab client.
package gitlab

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/open-beagle/go-scm/scm"
)

// New returns a new GitLab API client.
func New(uri string) (*scm.Client, error) {
	// 确保 URI 以斜杠结尾，避免 url.Parse 拼接时丢失 host
	// 例如：https://gitlab.wodcloud.com + api/v4/... 会变成 https:/api/v4/...
	// 正确：https://gitlab.wodcloud.com/ + api/v4/... = https://gitlab.wodcloud.com/api/v4/...
	if !strings.HasSuffix(uri, "/") {
		uri = uri + "/"
	}

	base, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(base.Path, "/") {
		base.Path = base.Path + "/"
	}
	client := &wrapper{new(scm.Client)}
	client.BaseURL = base
	// initialize services
	client.Driver = scm.DriverGitlab
	client.Linker = &linker{base.String()}
	client.Contents = &contentService{client}
	client.Git = &gitService{client}
	client.Issues = &issueService{client}
	client.Organizations = &organizationService{client}
	client.Milestones = &milestoneService{client}
	client.PullRequests = &pullService{client}
	client.Repositories = &repositoryService{client}
	client.Releases = &releaseService{client}
	client.Reviews = &reviewService{client}
	client.Users = &userService{client}
	client.Webhooks = &webhookService{client}
	return client.Client, nil
}

// NewDefault returns a new GitLab API client using the
// default gitlab.com address.
func NewDefault() *scm.Client {
	client, _ := New("https://gitlab.com")
	return client
}

// wraper wraps the Client to provide high level helper functions
// for making http requests and unmarshaling the response.
type wrapper struct {
	*scm.Client
}

// do wraps the Client.Do function by creating the Request and
// unmarshalling the response.
func (c *wrapper) do(ctx context.Context, method, path string, in, out interface{}) (*scm.Response, error) {
	req := &scm.Request{
		Method: method,
		Path:   path,
	}

	// if we are posting or putting data, we need to
	// write it to the body of the request.
	if in != nil {
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(in)
		req.Header = map[string][]string{
			"Content-Type": {"application/json"},
		}
		req.Body = buf
	}

	// execute the http request
	res, err := c.Client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// parse the gitlab request id.
	res.ID = res.Header.Get("X-Request-Id")

	// parse the gitlab rate limit details.
	res.Rate.Limit, _ = strconv.Atoi(
		res.Header.Get("RateLimit-Limit"),
	)
	res.Rate.Remaining, _ = strconv.Atoi(
		res.Header.Get("RateLimit-Remaining"),
	)
	res.Rate.Reset, _ = strconv.ParseInt(
		res.Header.Get("RateLimit-Reset"), 10, 64,
	)

	// snapshot the request rate limit
	c.Client.SetRate(res.Rate)

	// if an error is encountered, unmarshal and return the
	// error response.
	if res.Status > 300 {
		err := new(Error)
		json.NewDecoder(res.Body).Decode(err)
		return res, err
	}

	if out == nil {
		return res, nil
	}

	// if a json response is expected, parse and return
	// the json response.
	return res, json.NewDecoder(res.Body).Decode(out)
}

// Error represents a GitLab error.
type Error struct {
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}
