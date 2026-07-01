package gogs

import (
	"bytes"
	"context"
	"fmt"

	"github.com/open-beagle/bdpulse-scm/scm"
)

type contentService struct {
	client *wrapper
}

func (s *contentService) Find(ctx context.Context, repo, path, ref string) (*scm.Content, *scm.Response, error) {
	endpoint := fmt.Sprintf("api/v1/repos/%s/raw/%s/%s", repo, scm.TrimRef(ref), path)
	out := new(bytes.Buffer)
	res, err := s.client.do(ctx, "GET", endpoint, nil, out)
	return &scm.Content{
		Path: path,
		Data: out.Bytes(),
	}, res, err
}

func (s *contentService) Create(ctx context.Context, repo, path string, params *scm.ContentParams) (*scm.Response, error) {
	return nil, scm.ErrNotSupported
}

func (s *contentService) Update(ctx context.Context, repo, path string, params *scm.ContentParams) (*scm.Response, error) {
	return nil, scm.ErrNotSupported
}

func (s *contentService) Delete(ctx context.Context, repo, path string, params *scm.ContentParams) (*scm.Response, error) {
	return nil, scm.ErrNotSupported
}

func (s *contentService) List(ctx context.Context, repo, path, ref string, _ scm.ListOptions) ([]*scm.ContentInfo, *scm.Response, error) {
	return nil, nil, scm.ErrNotSupported
}
