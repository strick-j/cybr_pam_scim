package cybr_pam_scim

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/strick-j/cybr_pam_scim/pkg/cybr_pam_scim/types"
	"golang.org/x/exp/slices"
)

var (
	PrivilegedDatas types.PrivilegedDatas
	PrivilegedData  types.PrivilegedData
)

// GetPrivilegedData retrieves all Privileged Data (Accounts, SSHKeys, etc...) via the SCIM API.
// The response from the SCIM API is returned as the types.PrivilegedDatas struct
//
// Example Usage:
//		getPrivilegedData, err := s.GetPrivilegedData(context.Background)
//
func (s *Service) GetPrivilegedData(ctx context.Context) (*types.PrivilegedDatas, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s", "PrivilegedData"), &PrivilegedDatas); err != nil {
		return nil, fmt.Errorf("failed to get Privielged Data: %w", err)
	}

	return &PrivilegedDatas, nil
}

// GetPrivilegedDataIndex retrieves a limited subset of Privileged Data based on a starting index and count.
// The response from the SCIM API is returned as the types.PrivilegedDatas struct
//
// Requires PVWA 12.2+
//
// Example Usage:
//		getPrivielegedDataIndex, err := s.GetPrivilegedDataIndex(context.Background, 10, 5)
//
func (s *Service) GetPrivilegedDataIndex(ctx context.Context, startIndex int, count int) (*types.PrivilegedDatas, error) {
	pathEscapedQuery := url.PathEscape("startIndex=" + strconv.Itoa(startIndex) + "&count=" + strconv.Itoa(count))
	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "PrivilegedData", pathEscapedQuery), &PrivilegedDatas); err != nil {
		return nil, fmt.Errorf("failed to get Privileged Data: %w", err)
	}

	return &PrivilegedDatas, nil
}

// GetPrivilegedDataSort retrieves and sorts all Safes via the SCIM API based on provided
// The response from the SCIM API is returned as the types.Containers struct
// Supported "sortBy" fields: name, id, type, meta.created, meta.lastModified, or meta.location
// Supported "sortOrder" fileds are: ascending or descending
//
// Requires PVWA 12.2+
//
// Example Usage:
//		getPrivilegedDataSort, err := s.GetPrivilegedDataSort(context.Background, "name", "ascending")
//
func (s *Service) GetPrivilegedDataSort(ctx context.Context, sortBy string, sortOrder string) (*types.PrivilegedDatas, error) {
	var pathEscapedQuery string
	allowedSortBy := []string{"name", "id", "type", "meta.created", "meta.lastmodified", "meta.location"}
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
		return nil, fmt.Errorf("invalid sortBy value provided, the only accepted value is name, id, meta.created, meta.lastmodified, or meta.location")
	}

	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "PrivilegedData", pathEscapedQuery), &PrivilegedDatas); err != nil {
		return nil, fmt.Errorf("failed to get Privileged Data: %w", err)
	}

	return &PrivilegedDatas, nil
}

// GetPrivilegedDataById retrieves a data point based on Id via the SCIM API.
// The response from the SCIM API is returned as the types.PrivilegedData struct.
//
// Requires PVWA 12.2+
//
// Example Usage:
//		getPrivilegedDataById, err := s.GetPrivilegedDataById(context.Background, "92_2")
//
func (s *Service) GetPrivilegedDataById(ctx context.Context, id string) (*types.PrivilegedData, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s/%s", "PrivilegedData", id), &PrivilegedData); err != nil {
		return nil, fmt.Errorf("failed to get Privileged data %s: %w", id, err)
	}

	return &PrivilegedData, nil
}

// GetPrivilegedDataByFilter retrieves Privileged Data based on a provided filter via the SCIM API.
// The response from the SCIM API is returned as the types.PrivilegedDatas struct.
// filterType is the json key (e.g. name, id)
// filterQuery is the json value (e.g. VaultInternal)
//
// Notes: Filter query is case sensitive
//
// Example Usage:
//      getPrivilegedDataByFilter, err := s.GetPrivilegedDataByFilter(context.Background, "name", "exampleadmin")
//      getPrivilegedDataByFilter, err := s.GetPrivilegedDataByFilter(context.Background, "id", "92_3")
//
func (s *Service) GetPrivilegedDataByFilter(ctx context.Context, filterType string, filterQuery string) (*types.PrivilegedDatas, error) {
	pathEscapedQuery := url.PathEscape("filter=" + filterType + " eq " + filterQuery)
	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "PrivilegedDatas", pathEscapedQuery), &PrivilegedDatas); err != nil {
		return nil, fmt.Errorf("failed to get Safe Permissions based on filter parameters - %s = %s: %w", filterType, filterQuery, err)
	}

	return &PrivilegedDatas, nil
}

// AddPrivilegedData attempts a "POST" operation to add Privileged Data to the Vault
// and requires a types.PrivilegedData struct with the desired data.
// The response from the SCIM API is returned as the types.PrivilegedData struct.
//
// Example Usage:
//		PrivilegedData := types.PrivilegedData {
//			Name: "NewAccount",
//			Type: "password",
//			Schemas: []string{"urn:ietf:params:scim:schemas:pam:1.0:PrivilegedData", "urn:ietf:params:scim:schemas:cyberark:1.0:PrivilegedData"}
//			UrnIetfParamsScimSchemasCyberark10PrivilegedData: types.UrnIetfParamsScimSchemasCyberark10PrivilegedData {
//				Safe: "safeName",
//				Properties: []types.Properties {
//					{
//						Key: "username"
//						Value: "ExampleUser"
//					},
//					{
//						Key: "address"
//						Value: "ExampleAddress"
//					},
//					{
//						Key: "platformId"
//						Value: "ExamplePlatform"
//					},
//					{
//						Key: "secret"
//						Value: "ExampleSecret"
//					},
//				},
//			},
//		}
//      addPrivilegedData, err := s.AddPrivilegedData(context.Background, PrivilegedData)
//
func (s *Service) AddPrivilegedData(ctx context.Context, privilegedData types.PrivilegedData) (*types.PrivilegedData, error) {
	if err := s.client.Post(ctx, fmt.Sprintf("/%s", "PrivilegedData"), privilegedData, &PrivilegedData); err != nil {
		return nil, fmt.Errorf("failed to add permissions to safe: %w", err)
	}

	return &PrivilegedData, nil
}

// UpdatePrivilegedData attempts to perform a "PUT" operation against Privileged Data and requires
// a types.PrivilegedData struct with the desired udpates. The PrivilegedData Id must be part of the struct
// for this function to work properly.
//
// Requires PVWA 12.2+
//
// Example Usage:
// 		PrivilegedDataUpdate := types.PrivilegedData {
//			Name: "NewAccount",
//			Id: "62_3"
//			Schemas: []string{"urn:ietf:params:scim:schemas:pam:1.0:PrivilegedData", "urn:ietf:params:scim:schemas:cyberark:1.0:PrivilegedData"}
//			UrnIetfParamsScimSchemasCyberark10PrivilegedData: types.UrnIetfParamsScimSchemasCyberark10PrivilegedData {
//				Safe: "safeName",
//				Properties: []types.Properties {
//					{
//						Key: "username"
//						Value: "ExampleUser"
//					},
//					{
//						Key: "address"
//						Value: "ExampleAddress"
//					},
//					{
//						Key: "platformId"
//						Value: "ExampleNewPlatform"
//					},
//				},
//		}
//      addPrivilegedData, err := s.AddPrivilegedData(context.Background, PrivilegedData)
//
func (s *Service) UpdatePrivilegedData(ctx context.Context, privilegedData types.PrivilegedData) (*types.PrivilegedData, error) {
	if err := s.client.Put(ctx, fmt.Sprintf("/%s/%s", "PrivilegedData", privilegedData.Id), PrivilegedData, &PrivilegedData); err != nil {
		return nil, fmt.Errorf("failed to update Privileged Data: %w", err)
	}

	return &PrivilegedData, nil
}

// ModifyPrivilegedData attempts to perform a "PATCH" operation against Privileged Data and requires
// a types.PrivilegedData struct containing the "operation" substruct with the desired udpates.
// The PrivilegedData Id must be part of the struct for this function to work properly.
//
// Example Usage:
//		PrivilegedDataModify := types.PrivilegedData {
//			Id: "62_3"
//			Schemas: []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"}
//			Operations: []types.Operations {
//				{
//					Op: "replace",
//					Path: "urn:ietf:params:scim:schemas:cyberark:1.0:PrivilegedData.properties",
//					Value: []types.Value {
//						{
//							Key: "address",
//							Value: "newAddress",
//						},
//					},
//				},
//			},
//		}
//      modifyPrivilegedData, err := s.ModifyPrivilegedData(context.Background, PrivilegedDataModify)
//
func (s *Service) ModifyPrivilegedData(ctx context.Context, privilegedData types.PrivilegedData) (*types.PrivilegedData, error) {
	if err := s.client.Patch(ctx, fmt.Sprintf("/%s/%s", "PrivilegedData", privilegedData.Id), PrivilegedData, &PrivilegedData); err != nil {
		return nil, fmt.Errorf("failed to update Privileged Data: %w", err)
	}

	return &PrivilegedData, nil
}

// DeletePrivilegedData attempts to perform a "DELETE" operation against Privileged Data
// based on the provided Id. No response is provided if deletion is successful.
//
// An error will be returned if an attempt is made to delete multiple Safes or
// if deletion is attempted twice.
//
// Example Usage:
//		err := s.DeletePrivilegedData(context.Background, "62_3")
//
func (s *Service) DeletePrivilegedData(ctx context.Context, id string) error {
	if err := s.client.Delete(ctx, fmt.Sprintf("/%s/%s", "PrivilegedData", id), nil); err != nil {
		return fmt.Errorf("failed to delete privileged data %s: %w", id, err)
	}

	return nil
}
