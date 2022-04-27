package types

type Containers struct {
	Schemas      []string    `json:"schemas"`
	TotalResults int         `json:"totalResults"`
	ItemsPerPage int         `json:"itemsPerPage"`
	StartIndex   int         `json:"startIndex"`
	Resources    []Container `json:"Resources"`
}

type Container struct {
	Name                                   string                                 `json:"name"`
	DisplayName                            string                                 `json:"displayName,omitempty"`
	Description                            string                                 `json:"description,omitempty"`
	Type                                   string                                 `json:"type,omitempty"`
	Parent                                 Parent                                 `json:"parent,omitempty"`
	Owner                                  Owner                                  `json:"owner,omitempty"`
	PrivilegedData                         []PrivilegedDataRef                    `json:"privilegedData,omitempty"`
	Schemas                                []string                               `json:"schemas"`
	Id                                     string                                 `json:"id"`
	ExternalId                             string                                 `json:"externalId,omitempty"`
	Meta                                   Meta                                   `json:"meta"`
	UrnIetfParamsScimSchemasCyberark10Safe UrnIetfParamsScimSchemasCyberark10Safe `json:"urn:ietf:params:scim:schemas:cyberark:1.0:Safe"`
}

// Container Attribute
type UrnIetfParamsScimSchemasCyberark10Safe struct {
	NumberOfDaysRetention int    `json:"NumberOfDaysRetention,omitempty"`
	ManagingCPM           string `json:"ManagingCPM,omitempty"`
}

type Owner struct {
	Value   string `json:"value,omitempty"`
	Ref     string `json:"$ref,omitempty"`
	Display string `json:"display,omitempty"`
}

type Parent struct {
	Value   string `json:"value,omitempty"`
	Ref     string `json:"$ref,omitempty"`
	Display string `json:"display,omitempty"`
}
