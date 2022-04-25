package pamscim

import (
	"context"
	"fmt"

	"github.com/strick-j/cybr-pam-scim/pkg/cybr/pamscim/types"
)

var (
	Users types.Users
	User  types.User
)

func (s *Service) GetUsers(ctx context.Context) (*types.Users, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s", "users"), &Users); err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	return &Users, nil
}

func (s *Service) GetUser(ctx context.Context, id string) (*types.Users, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s/%s", "users", id), &Users); err != nil {
		return nil, fmt.Errorf("failed to get user %s: %w", id, err)
	}
	return &Users, nil
}

func (s *Service) AddUser(ctx context.Context, ob types.User) (*types.User, error) {
	if err := s.client.Post(ctx, fmt.Sprintf("/%s", "users"), ob, &Users); err != nil {
		return nil, fmt.Errorf("failed to add %s: %w", ob.UserName, err)
	}

	return &User, nil
}

func (s *Service) UpdateUser(ctx context.Context, id string, ob types.User) (*types.User, error) {
	if err := s.client.Put(ctx, fmt.Sprintf("/%s/%s", "users", id), ob, &Users); err != nil {
		return nil, fmt.Errorf("failed to update user %s: %w", id, err)
	}
	return &User, nil
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	if err := s.client.Delete(ctx, fmt.Sprintf("/%s/%s", "users", id), nil); err != nil {
		return fmt.Errorf("failed to delete user %s: %w", id, err)
	}
	return nil
}
