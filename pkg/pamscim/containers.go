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
	Containers types.Containers
	Container  types.Container
)

// GetSafes retrieves all Safes via the SCIM API.
// The response from the SCIM API is returned as the types.Containers struct
//
// Example Usage:
// 		getSafes, err := s.GetSafes(context.Bacground)
//
func (s *Service) GetSafes(ctx context.Context) (*types.Containers, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s", "Containers"), &Containers); err != nil {
		return nil, fmt.Errorf("failed to get Safes: %w", err)
	}

	return &Containers, nil
}

// GetSafesIndex retrieves a limited subset of Safes based on a starting index and count.
// The response from the SCIM API is returned as the types.Containers struct
//
// Example Usage:
//        indexSafes, err := s.GetSafesIndex(context.Background, 10, 5)
//
func (s *Service) GetSafesIndex(ctx context.Context, startIndex int, count int) (*types.Containers, error) {
	pathEscapedQuery := url.PathEscape("startIndex=" + strconv.Itoa(startIndex) + "&count=" + strconv.Itoa(count))
	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "Containers", pathEscapedQuery), &Containers); err != nil {
		return nil, fmt.Errorf("failed to get Safes: %w", err)
	}

	return &Containers, nil
}

// GetSafesSort retrieves and sorts all Safes via the SCIM API based on provided
// The response from the SCIM API is returned as the types.Containers struct
// Supported "sortBy" fields: name, displayName, description, id, meta.created, meta.lastModified, or meta.location
// Supported "sortOrder" fileds are: ascending or descending
//
// Example Usage:
//        sortSafes, err := s.GetSafesSort(context.Background, "SafeName", "ascending")
//
func (s *Service) GetSafesSort(ctx context.Context, sortBy string, sortOrder string) (*types.Containers, error) {
	allowedSortBy := []string{"name", "displayName", "description", "id", "meta.created", "meta.lastmodified", "meta.location"}
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
		return nil, fmt.Errorf("invalid sortBy value provided, accepted values are name, displayName, description, id, meta.created, meta.lastModified, or meta.location")
	}

	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "Containers", pathEscapedQuery), &Containers); err != nil {
		return nil, fmt.Errorf("failed to get Safes: %w", err)
	}

	return &Containers, nil
}

// GetSafeByName retrieves a single Safe by Safe Name via the SCIM API.
// The response from the SCIM API is returned as the types.Container struct.
//
// Requires PVWA 12.2
//
// Example Usage:
//        getSafeByName, err := s.GetSafeByName(context.Background, "NotificationEngine")
//
func (s *Service) GetSafeByName(ctx context.Context, safeName string) (*types.Container, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s/%s", "Containers", safeName), &Container); err != nil {
		return nil, fmt.Errorf("failed to get Safe %s: %w", safeName, err)
	}

	return &Container, nil
}

// GetSafeByFilter retrieves a single Safe based on a provided filter via the SCIM API.
// The response from the SCIM API is returned as the types.Container struct.
// filterType is the json key (e.g. name, displayName, description, etc...)
// filterQuery is the json value (e.g. PVWATicketingSystem)
//
// Notes: Filter query is case sensitive
//
// Example Usage:
//      getContainer, err := s.GetContainerByFilter(context.Background, "name", "PVWATicketingSystem")
//
func (s *Service) GetSafeByFilter(ctx context.Context, filterType string, filterQuery string) (*types.Container, error) {
	pathEscapedQuery := url.PathEscape("filter=" + filterType + " eq " + filterQuery)
	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "Safes", pathEscapedQuery), &Container); err != nil {
		return nil, fmt.Errorf("failed to get Container based on filter parameters - %s = %s: %w", filterType, filterQuery, err)
	}

	return &Container, nil
}

// AddSafe attempts add a single Safe and requires a types.Container struct with the
// desired Safe information. The response from the SCIM API is returned as the types.Container struct.
//
// At a minimum the types.Container struct must contain the following:
//       Name
//       Schemas
//
// Example Usage:
// 		safe := types.Container {
//			Name:     "ExampleSafe",
//			Schemas:  []string{"urn:ietf:params:scim:schemas:core:1.0:Container"},
// 		}
//      addSafe, err := s.AddSafe(context.Background, safe)
//
func (s *Service) AddSafe(ctx context.Context, safe types.Container) (*types.Container, error) {
	if err := s.client.Post(ctx, fmt.Sprintf("/%s", "Safes"), safe, &Container); err != nil {
		return nil, fmt.Errorf("failed to add Container %s: %w", safe.Name, err)
	}

	return &Container, nil
}

// UpdateSafe attempts to perform a "PUT" operation against a single Safe and requires
// a passed object in the form of types.Container. Null values are supported enabling
// the removal of Container attribute values.
//
// Requires PVWA 12.2+
//
// Example Usage:
// 		safe := types.Container {
//			Name:        "ExampleSafe,
//          Displayname: "ExampleSafeDisplayName" // Updated name
//			Schemas:     []string{"urn:ietf:params:scim:schemas:core:1.0:Container"},
//		}
//      updateSafe, err := s.UpdateContainer(context.Background, safe)
//
func (s *Service) UpdateSafe(ctx context.Context, safe types.Container) (*types.Container, error) {
	if err := s.client.Put(ctx, fmt.Sprintf("/%s/%s", "Safes", safe.Id), Container, &Container); err != nil {
		return nil, fmt.Errorf("failed to update Container %s: %w", safe.Id, err)
	}

	return &Container, nil
}

// DeleteSafe attempts to perform a "DELETE" operation against a single Safe by
// Safe Name via the SCIM API and does not return a response is successful.
// An error will be returned if an attempt is made to delete multiple Safes or
// if deletion is attempted twice.
//
// Example Usage:
//        err := s.DeleteSafe(context.Background, "ExampleSafe")
//
func (s *Service) DeleteSafe(ctx context.Context, name string) error {
	if err := s.client.Delete(ctx, fmt.Sprintf("/%s/%s", "Safes", name), nil); err != nil {
		return fmt.Errorf("failed to delete Container %s: %w", name, err)
	}

	return nil
}
