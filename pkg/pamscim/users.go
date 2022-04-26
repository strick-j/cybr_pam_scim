package pamscim

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/strick-j/cybr-pam-scim/pkg/pamscim/types"
	"golang.org/x/exp/slices"
)

var (
	Users types.Users
	User  types.User
)

// GetUsers retrieves all users via the SCIM API.
// The response from the SCIM API is returned as the types.Users struct
//
// Example Usage:
// 		getUsers, err := s.GetUsers(context.Bacground)
//
func (s *Service) GetUsers(ctx context.Context) (*types.Users, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s", "users"), &Users); err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return &Users, nil
}

// GetUsersIndex retrieves a limited subset of Users based on a starting index and count.
// The response from the SCIM API is returned as the types.Users struct
//
// Example Usage:
//        indexUsers, err := s.GetUsersIndex(context.Background, 1, 5)
//
func (s *Service) GetUsersIndex(ctx context.Context, startIndex int, count int) (*types.Users, error) {
	pathEscapedQuery := url.PathEscape("startIndex=" + strconv.Itoa(startIndex) + "&count=" + strconv.Itoa(count))
	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "Users", pathEscapedQuery), &Users); err != nil {
		return nil, fmt.Errorf("failed to get Users: %w", err)
	}

	return &Users, nil
}

// GetUsersSort retrieves and sorts all users via the SCIM API based on provided
// The response from the SCIM API is returned as the types.Users struct
// Supported "sortBy" fields: active, userName, displayName, name.givenName, name.familyName, userType, id, meta.created, meta.lastmodified
// Supported "sortOrder" fileds are: ascending or descending
//
// Example Usage:
//        sortUsers, err := s.GetUsersSort(context.Background, "userName", "ascending")
//
func (s *Service) GetUsersSort(ctx context.Context, sortBy string, sortOrder string) (*types.Users, error) {
	allowedSortBy := []string{"active", "userName", "displayName", "name.familyName", "name.givenName", "userType", "id", "meta.created", "meta.lastmodified", "meta.location"}
	var pathEscapedQuery string
	// Input validations:
	if slices.Contains(allowedSortBy, sortBy) {
		if sortOrder == "ascending" || sortOrder == "descending" {
			pathEscapedQuery = url.PathEscape("sortBy=" + sortBy + "&sortOrder=" + sortOrder)
		} else if sortOrder == "" {
			pathEscapedQuery = url.PathEscape("sortBy=" + sortBy)
		} else {
			return nil, fmt.Errorf("invalid sortOrder provided, accepted values are ascending, descending, or no input")
		}
	} else {
		return nil, fmt.Errorf("invalid sortBy value provided, accepted values are active, userName, displayName, name.givenName, name.familyName, userType, id, meta.created, meta.lastmodified, or meta.location")
	}
	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "users", pathEscapedQuery), &Users); err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return &Users, nil
}

// GetUserById retrieves a single user by User Id via the SCIM API.
// The response from the SCIM API is returned as the types.User struct.
//
// Requires PVWA 12.2
//
// Example Usage:
//        getUser, err := s.GetUserById(context.Background, "1")
//
func (s *Service) GetUserById(ctx context.Context, id string) (*types.User, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s/%s", "users", id), &User); err != nil {
		return nil, fmt.Errorf("failed to get user %s: %w", id, err)
	}

	return &User, nil
}

// GetUserByFilter retrieves a single user based on a provided filter via the SCIM API.
// The response from the SCIM API is returned as the types.User struct.
// filterType is the json key (e.g. displayName, name.givenName, name.familyName)
// filterQuery is the json value (e.g. john.smith@example.com, "John", "Smith")
//
// Notes: Filter query is case sensitive
//
// Example Usage:
//      getUser, err := s.GetUserByFilter(context.Background, "userName", "john.smith@example.com")
// 		getUser, err := s.GetUserByFilter(context.Background, "name.familyName", "Smith")
//
func (s *Service) GetUserByFilter(ctx context.Context, filterType string, filterQuery string) (*types.User, error) {
	pathEscapedQuery := url.PathEscape("filter=" + filterType + " eq " + filterQuery)
	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "users", pathEscapedQuery), &User); err != nil {
		return nil, fmt.Errorf("failed to get user based on filter parameters - %s = %s: %w", filterType, filterQuery, err)
	}

	return &User, nil
}

// AddUser attempts add a single user and requires a types.User struct with the
// desired user information. The response from the SCIM API is returned as the types.User struct.
//
// At a minimum the types.User struct must contain the following:
//       UserName
//       Password
//       Schemas
//
// Example Usage:
// 		user := types.User {
//			UserName: "JohnDoe@example.com",
//			Password: "ExamplePass",
//			Schemas:  []string{"urn:ietf:params:scim:schemas:core:2.0:User"},
// 		}
//      addUser, err := s.AddUser(context.Background, user)
//
func (s *Service) AddUser(ctx context.Context, user types.User) (*types.User, error) {
	if err := s.client.Post(ctx, fmt.Sprintf("/%s", "users"), User, &User); err != nil {
		return nil, fmt.Errorf("failed to add user %s: %w", user.UserName, err)
	}

	return &User, nil
}

// UpdateUser attempts to perform a "PUT" operation against a single User and requires
// a passed object in the form of types.User. Null values are supported enabling
// the removal of user attribute values.
//
// Example Usage:
// 		user := types.User {
//			UserName:    "JohnDoe@example.com",
//          Active:      false, // Disable User via PUT
//			Schemas:     []string{"urn:ietf:params:scim:schemas:core:2.0:User"},
//		}
//      updateUser, err := s.UpdateUser(context.Background, user)
//
func (s *Service) UpdateUser(ctx context.Context, user types.User) (*types.User, error) {
	if err := s.client.Put(ctx, fmt.Sprintf("/%s/%s", "users", user.Id), User, &User); err != nil {
		return nil, fmt.Errorf("failed to update user %s: %w", user.Id, err)
	}

	return &User, nil
}

// DeleteUser attempts to perform a "DELETE" operation against a single User by
// User Id via the SCIM API and does not return a response is successful.
// An error will be returned if an attempt is made to delete multiple Users or
// if deletion is attempted twice.
//
// Example Usage:
//        err := s.DeleteUser(context.Background, "8")
//
func (s *Service) DeleteUser(ctx context.Context, id string) error {
	if err := s.client.Delete(ctx, fmt.Sprintf("/%s/%s", "users", id), nil); err != nil {
		return fmt.Errorf("failed to delete user %s: %w", id, err)
	}

	return nil
}
