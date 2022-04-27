package cybr_pam_scim

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/strick-j/cybr-pam-scim/pkg/types"
)

var (
	ContainerPermissions types.ContainerPermissions
	ContainerPermission  types.ContainerPermission
)

// GetSafePermissions retrieves all Safes via the SCIM API.
// The response from the SCIM API is returned as the types.Containers struct
//
// Example Usage:
// 		getSafePermissions, err := s.GetSafePermissions(context.Background)
//
func (s *Service) GetSafePermissions(ctx context.Context) (*types.ContainerPermissions, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s", "ContainerPermissions"), &ContainerPermissions); err != nil {
		return nil, fmt.Errorf("failed to get Safe Permissions: %w", err)
	}

	return &ContainerPermissions, nil
}

// GetSafePermissionsIndex retrieves a limited subset of Safe Permissions based on a starting index and count.
// The response from the SCIM API is returned as the types.Containers struct
//
// Example Usage:
//        indexSafePermissions, err := s.GetSafePermissionsIndex(context.Background, 10, 5)
//
func (s *Service) GetSafePermissionsIndex(ctx context.Context, startIndex int, count int) (*types.ContainerPermissions, error) {
	pathEscapedQuery := url.PathEscape("startIndex=" + strconv.Itoa(startIndex) + "&count=" + strconv.Itoa(count))
	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "ContainerPermissions", pathEscapedQuery), &ContainerPermissions); err != nil {
		return nil, fmt.Errorf("failed to get Safe Permissions: %w", err)
	}

	return &ContainerPermissions, nil
}

// GetSafePermissionsSort retrieves and sorts all Safes via the SCIM API based on provided
// The response from the SCIM API is returned as the types.Containers struct
// Supported "sortBy" fields: id
// Supported "sortOrder" fileds are: ascending or descending
//
// Example Usage:
//        sortSafePermissions, err := s.GetSafePermissionsSort(context.Background, "SafeName", "ascending")
//
func (s *Service) GetSafePermissionsSort(ctx context.Context, sortBy string, sortOrder string) (*types.ContainerPermissions, error) {
	var pathEscapedQuery string
	// Input validations:
	if sortBy == "id" {
		if sortOrder == "ascending" || sortOrder == "descending" {
			pathEscapedQuery = url.PathEscape("sortBy=" + sortBy + "&sortOrder=" + sortOrder)
		} else if sortOrder == "" {
			pathEscapedQuery = url.PathEscape("sortBy=" + sortBy)
		} else {
			return nil, fmt.Errorf("invalid sortOrder provided, accepted values are ascending, descending, or no input")
		}
	} else {
		return nil, fmt.Errorf("invalid sortBy value provided, the only accepted value is id")
	}

	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "ContainerPermissions", pathEscapedQuery), &ContainerPermissions); err != nil {
		return nil, fmt.Errorf("failed to get Safes: %w", err)
	}

	return &ContainerPermissions, nil
}

// GetSafePermissionsByName retrieves a single Safe by Safe Name amd a User or Group Name via the SCIM API.
// The response from the SCIM API is returned as the types.Container struct.
//
// Requires PVWA 12.2
//
// Example Usage:
//        getSafePermissionsByName, err := s.GetSafePermissionsByName(context.Background, "VaultInternal", "EPMAgent")
//
func (s *Service) GetSafePermissionsByName(ctx context.Context, safeName string, userOrGroupName string) (*types.ContainerPermission, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s/%s:%s", "ContainerPermissions", safeName, userOrGroupName), &ContainerPermission); err != nil {
		return nil, fmt.Errorf("failed to get User (%s) permissions on Safe %s: %w", userOrGroupName, safeName, err)
	}

	return &ContainerPermission, nil
}

// GetSafePermissionsByFilter retrieves a single Safe based on a provided filter via the SCIM API.
// The response from the SCIM API is returned as the types.Containers struct.
// filterType is the json key (e.g. container.name, container.display, container.value, user.display, user.value, group.value, group.display)
// filterQuery is the json value (e.g. VaultInternal)
//
// Notes: Filter query is case sensitive
//
// Example Usage:
//      // Return all permission on a single safe
//      getSafePermissionsByFilter, err := s.GetSafePermissionsByFilter(context.Background, "container.name", "PVWATicketingSystem")
//
//      // Return specific users permissions on all safes
//      getSafePermissionsByFilter, err := s.GetSafePermissionsByFilter(context.Background, "user.display", "EPMAgent")
//
//		// Return specific group permissions on all safes
//		getSafePermissionsByFilter, err := s.GetSafePermissionsByFilter(context.Background, "group.value", "18")
//
func (s *Service) GetSafePermissionByFilter(ctx context.Context, filterType string, filterQuery string) (*types.ContainerPermissions, error) {
	pathEscapedQuery := url.PathEscape("filter=" + filterType + " eq " + filterQuery)
	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "ContainerPermissions", pathEscapedQuery), &ContainerPermissions); err != nil {
		return nil, fmt.Errorf("failed to get Safe Permissions based on filter parameters - %s = %s: %w", filterType, filterQuery, err)
	}

	return &ContainerPermissions, nil
}

// AddSafePermissions attempts a "POST" operation to addpermissions to a single Safe for a
// specific user and requires a types.ContainerPermission struct with the desired
// User, Safe, and Rights information.
// The response from the SCIM API is returned as the types.ContainerPermission struct.
//
// Example Usage:
// 		safePermission := types.ContainerPermission {
//    		Schemas: []string{"urn:ietf:params:scim:schemas:pam:1.0:ContainerPermission"},
//		    UserRef.Display: "john.smith@example.com",
//			ContainerRef.Name: "ExampleContainer",
//          Rights: []string{"UseAccounts","RetrieveAccounts","ListAccounts"},
// 		}
//      addSafePermissions, err := s.AddSafePermissions(context.Background, safePermission)
//
func (s *Service) AddSafePermissions(ctx context.Context, safePermission types.ContainerPermission) (*types.ContainerPermission, error) {
	if err := s.client.Post(ctx, fmt.Sprintf("/%s", "ContainerPermissions"), ContainerPermission, &ContainerPermission); err != nil {
		return nil, fmt.Errorf("failed to add permissions to safe: %w", err)
	}

	return &ContainerPermission, nil
}

// UpdateSafe attempts to perform a "PUT" operation against a single Safe and requires
// a types.ContainerPermission struct with the desired udpates. Null values are supported enabling
// the removal of Container attribute values.
//
// Requires PVWA 12.2+
//
// Example Usage:
// 		safePermissionUpdate := types.ContainerPermission {
//    		Schemas: []string{"urn:ietf:params:scim:schemas:pam:1.0:ContainerPermission"},
//		    UserRef.Display: "john.smith@example.com",
//			ContainerRef.Name: "ExampleContainer",
//          Rights: []string{"UseAccounts","RetrieveAccounts","ListAccounts","ManageSafe"},
// 		}
//      updateSafePermissions, err := s.UpdateSafePermissions(context.Background, safePermissionUpdate)
//
func (s *Service) UpdateSafePermissions(ctx context.Context, safePermission types.ContainerPermission) (*types.Container, error) {
	if err := s.client.Put(ctx, fmt.Sprintf("/%s/%s:%s", "ContainerPermissions", safePermission.Container.Name, safePermission.User.Display), ContainerPermission, &ContainerPermission); err != nil {
		return nil, fmt.Errorf("failed to update Safe Permissions: %w", err)
	}

	return &Container, nil
}

// DeleteSafe attempts to perform a "DELETE" operation against a single Safe for a
// single user or group Safe Name and the User or Group Name.
//
// An error will be returned if an attempt is made to delete multiple Safes or
// if deletion is attempted twice.
//
// Example Usage:
//        err := s.DeleteSafePermission(context.Background, "ExampleSafe", "ExampleUser")
//
func (s *Service) DeleteSafePermission(ctx context.Context, safeName string, userOrGroupName string) error {
	if err := s.client.Delete(ctx, fmt.Sprintf("/%s/%s:%s", "ContainerPermissions", safeName, userOrGroupName), nil); err != nil {
		return fmt.Errorf("failed to remove %s Permissiosn from Safe %s: %w", userOrGroupName, safeName, err)
	}

	return nil
}
