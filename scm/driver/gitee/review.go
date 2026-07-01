package gitee

import (
	"context"

	"github.com/open-beagle/bdpulse-scm/scm"
)

type reviewService struct {
	client *wrapper
}

func (s *reviewService) Find(context.Context, string, int, int) (*scm.Review, *scm.Response, error) {
	return nil, nil, scm.ErrNotSupported
}

func (s *reviewService) List(context.Context, string, int, scm.ListOptions) ([]*scm.Review, *scm.Response, error) {
	return nil, nil, scm.ErrNotSupported
}

func (s *reviewService) Create(context.Context, string, int, *scm.ReviewInput) (*scm.Review, *scm.Response, error) {
	return nil, nil, scm.ErrNotSupported
}

func (s *reviewService) Delete(context.Context, string, int, int) (*scm.Response, error) {
	return nil, scm.ErrNotSupported
}
