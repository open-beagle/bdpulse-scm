package scm

import (
	"context"
	"time"
)

type (
	// User represents a user account.
	User struct {
		ID      string
		Login   string
		Name    string
		Email   string
		Avatar  string
		Created time.Time
		Updated time.Time
	}

	// Email represents a user email.
	Email struct {
		Value    string
		Primary  bool
		Verified bool
	}

	Netrc struct {
		Login string
		Token string
	}

	// UserService provides access to user account resources.
	UserService interface {
		// Find returns the authenticated user.
		Find(context.Context) (*User, *Response, error)

		// FindEmail returns the authenticated user email.
		FindEmail(context.Context) (string, *Response, error)

		// FindLogin returns the user account by username.
		FindLogin(context.Context, string) (*User, *Response, error)

		// ListEmail returns the user email list.
		ListEmail(context.Context, ListOptions) ([]*Email, *Response, error)

		// FindNetrc returns netrc credentials when supported by the driver.
		FindNetrc(context.Context, string) (*Netrc, *Response, error)
	}
)
