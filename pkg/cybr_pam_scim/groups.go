package cybr_pam_scim

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/strick-j/cybr_pam_scim/pkg/cybr_pam_scim/types"
)

var (
	Groups types.Groups
	Group  types.Group
)

// GetGroups retrieves all groups via the SCIM API and returns them in the form of the
// The response from the SCIM API is returned as the types.Groups struct.
//
// Example Usage:
//		getGroups, err := s.GetGroups(context.Background)
//
func (s *Service) GetGroups(ctx context.Context) (*types.Groups, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s", "groups"), &Groups); err != nil {
		return nil, fmt.Errorf("failed to get groups: %w", err)
	}

	return &Groups, nil
}

// GetGroupsIndex retrieves a limited subset of groups based on a starting and count.
// The response from the SCIM API is returned as the types.Groups struct.
//
// Requires PVWA 12.2+
//
// Example Usage:
//		getGroupsIndex, err := s.GetGroupsIndex(context.Background, 1, 5)
//
func (s *Service) GetGroupsIndex(ctx context.Context, startIndex int, count int) (*types.Groups, error) {
	pathEscapedQuery := url.PathEscape("startIndex=" + strconv.Itoa(startIndex) + "&count=" + strconv.Itoa(count))
	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "Groups", pathEscapedQuery), &Groups); err != nil {
		return nil, fmt.Errorf("failed to get Groups: %w", err)
	}

	return &Groups, nil
}

// GetGroupsSort retrieves and sorts all Groups via the SCIM API.
// The response from the SCIM API is returned as the types.Groups struct.
// Supported "sortBy" fields are: displayName
// Supported "sortOrder" fileds are: ascending or descending
//
// Requires PVWA 12.2+
//
// Example Usage:
//		getGroupsSort, err := s.GetGroupsSort(context.Background, "displayName", "ascending")
//
func (s *Service) GetGroupsSort(ctx context.Context, sortBy string, sortOrder string) (*types.Groups, error) {
	var pathEscapedQuery string
	// Input validations:
	if sortBy == "displayName" {
		if sortOrder == "ascending" || sortOrder == "descending" {
			pathEscapedQuery = url.PathEscape("sortBy=" + sortBy + "&sortOrder=" + sortOrder)
		} else if sortOrder == "" {
			pathEscapedQuery = url.PathEscape("sortBy=" + sortBy)
		} else {
			return nil, fmt.Errorf("invalid sortOrder provided, accepted values are ascending, descending, or no input")
		}
	} else {
		return nil, fmt.Errorf("invalid sortBy value provide, accepted value is displayName")
	}

	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "Groups", pathEscapedQuery), &Groups); err != nil {
		return nil, fmt.Errorf("failed to get Groups: %w", err)
	}

	return &Groups, nil
}

// GetGroupById retrieves a single Group by Group Id via the SCIM API.
// The response from the SCIM API is returned as the types.Group struct.
//
// Requires PVWA 12.2+
//
// Example Usage:
//		getGroupById, err := s.GetGroupById(context.Background, "8")
//
func (s *Service) GetGroupById(ctx context.Context, id string) (*types.Group, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s/%s", "Groups", id), &Group); err != nil {
		return nil, fmt.Errorf("failed to get Group %s: %w", id, err)
	}

	return &Group, nil
}

// GetGroupByFilter retrieves a single Group based on a provided filter.
// The response from the SCIM API is returned as the types.Group struct.
// filterType should be displayName
// filterQuery should be the actual query (e.g. Auditors)
// The Groups struct response will only contain a single resource.
//
// Example Usage:
//		getGroupByFilter, err := s.GetGroupByFilter(context.Background, "displayName", "Auditors")
//
func (s *Service) GetGroupByFilter(ctx context.Context, filterType string, filterQuery string) (*types.Group, error) {
	var pathEscapedQuery string
	if filterType == "id" || filterType == "displayName" {
		pathEscapedQuery = url.PathEscape("filter=" + filterType + " eq \"" + filterQuery + "\"")
	} else {
		return nil, fmt.Errorf("invalid filterType provided, accepted types are id or displayName")
	}
	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "Groups", pathEscapedQuery), &Group); err != nil {
		return nil, fmt.Errorf("failed to get Group based on filter parameters - %s = %s: %w", filterType, filterQuery, err)
	}

	return &Group, nil
}

// AddGroup attempts add a single Group and requires a passed object in the form of
// types.Group. The response from the SCIM API is returned as the types.Group struct.
// At a minimum the struct must contain the following:
//		DisplayName
//		Schemas
//
// Example Usage:
//		Group := types.Group {
//			DisplayName: "ExampleGroup",
//			Schemas:     []string{"urn:ietf:params:scim:schemas:core:2.0:Group"},
//		}
//		addGroup, err := s.AddGroup(context.Background, Group)
//
func (s *Service) AddGroup(ctx context.Context, group types.Group) (*types.Group, error) {
	if err := s.client.Post(ctx, fmt.Sprintf("/%s", "Groups"), group, &Group); err != nil {
		return nil, fmt.Errorf("failed to add Group %s: %w", group.DisplayName, err)
	}

	return &Group, nil
}

// UpdateGroup attempts to perform a "PUT" operation against a single Group and requires
// a passed object in the form of types.Group. The "PUT" operation replaces an existing
// group with an updated version or creates a new group entirely.
//
// Example Usage:
//		Group := types.Group {
//			DisplayName: "ExampleGroup",
//			Schemas:     []string{"urn:ietf:params:scim:schemas:core:2.0:Group"},
//		}
//		addGroup, err := s.UpdateGroup(context.Background, Group)
//
func (s *Service) UpdateGroup(ctx context.Context, group types.Group) (*types.Group, error) {
	if err := s.client.Put(ctx, fmt.Sprintf("/%s/%s", "Groups", group.Id), group, &Group); err != nil {
		return nil, fmt.Errorf("failed to update Group %s: %w", group.Id, err)
	}

	return &Group, nil
}

// DeleteGroup attempts to perform a "DELETE" operation against a single Group by
// Group Id via the SCIM API and returns does not return a response is successful.
// An error will be returned if an attempt is made to delete multiple Groups or
// if deletion is attempted twice.
//
// Example Usage:
//		err := s.DeleteGroup(context.Background, "8")
//
func (s *Service) DeleteGroup(ctx context.Context, id string) error {
	if err := s.client.Delete(ctx, fmt.Sprintf("/%s/%s", "Groups", id), nil); err != nil {
		return fmt.Errorf("failed to delete Group %s: %w", id, err)
	}

	return nil
}
