package azure

import (
	"context"

	"github.com/open-beagle/bdpulse-scm/scm"
)

type userService struct {
	client *wrapper
}

func (s *userService) Find(ctx context.Context) (*scm.User, *scm.Response, error) {
	return nil, nil, scm.ErrNotSupported
}

func (s *userService) FindLogin(ctx context.Context, login string) (*scm.User, *scm.Response, error) {
	return nil, nil, scm.ErrNotSupported
}

func (s *userService) FindEmail(ctx context.Context) (string, *scm.Response, error) {
	return "", nil, scm.ErrNotSupported
}

func (s *userService) ListEmail(context.Context, scm.ListOptions) ([]*scm.Email, *scm.Response, error) {
	return nil, nil, scm.ErrNotSupported
}

func (s *userService) FindNetrc(ctx context.Context, id string) (*scm.Netrc, *scm.Response, error) {
	return nil, nil, scm.ErrNotSupported
}
