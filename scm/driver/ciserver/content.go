package ciserver

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/open-beagle/bdpulse-scm/scm"
)

type contentService struct {
	client *wrapper
}

func (s *contentService) Find(ctx context.Context, repo, path, ref string) (*scm.Content, *scm.Response, error) {
	var (
		endpoint string
		err      error
	)
	if strings.Contains(repo, "/") {
		endpoint, err = formatPath(s.client.paths.ContentBySlug, repo, ref, path)
	} else {
		endpoint, err = formatPath(s.client.paths.ContentByID, repo, ref, path)
	}
	if err != nil {
		return nil, nil, err
	}
	out := new(content)
	res, err := s.client.do(ctx, "GET", endpoint, nil, out)
	raw, berr := base64.StdEncoding.DecodeString(out.Content)
	if berr != nil {
		return nil, nil, err
	}
	return &scm.Content{
		Path:   out.FilePath,
		Data:   raw,
		Sha:    out.LastCommitID,
		BlobID: out.BlobID,
	}, res, err
}

func (s *contentService) Create(ctx context.Context, repo, path string, params *scm.ContentParams) (*scm.Response, error) {
	endpoint, err := formatPath(s.client.paths.ContentWrite, encode(repo))
	if err != nil {
		return nil, err
	}
	in := &createUpdateContent{
		FilePath:      path,
		Branch:        params.Branch,
		Content:       params.Data,
		CommitMessage: params.Message,
	}
	res, err := s.client.do(ctx, "POST", endpoint, in, nil)
	return res, err
}

func (s *contentService) Update(ctx context.Context, repo, path string, params *scm.ContentParams) (*scm.Response, error) {
	endpoint, err := formatPath(s.client.paths.ContentWrite, encode(repo))
	if err != nil {
		return nil, err
	}
	in := &createUpdateContent{
		FilePath:      path,
		Branch:        params.Branch,
		Content:       params.Data,
		CommitMessage: params.Message,
	}
	res, err := s.client.do(ctx, "POST", endpoint, in, nil)
	return res, err
}

func (s *contentService) Delete(ctx context.Context, repo, path string, params *scm.ContentParams) (*scm.Response, error) {
	return nil, scm.ErrNotSupported
}

func (s *contentService) List(ctx context.Context, repo, path, ref string, opts scm.ListOptions) ([]*scm.ContentInfo, *scm.Response, error) {
	return nil, nil, scm.ErrNotSupported
}

type content struct {
	FilePath     string `json:"filePath"`
	Encoding     string `json:"encoding"`
	Content      string `json:"content"`
	Ref          string `json:"ref"`
	BlobID       string `json:"blobId"`
	CommitID     string `json:"commitId"`
	LastCommitID string `json:"lastCommitId"`
}

type createUpdateContent struct {
	FilePath      string `json:"filePath"`
	Branch        string `json:"branch"`
	Content       []byte `json:"content"`
	CommitMessage string `json:"commit_message"`
}
