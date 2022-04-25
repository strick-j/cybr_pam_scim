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
	DisplayName                            string                                 `json:"displayName"`
	Description                            string                                 `json:"description"`
	Type                                   string                                 `json:"type"`
	Owner                                  Owner                                  `json:"owner"`
	PrivilegedData                         []PrivilegedDataRef                    `json:"privilegedData"`
	Schemas                                []string                               `json:"schemas"`
	ID                                     string                                 `json:"id"`
	Meta                                   Meta                                   `json:"meta"`
	UrnIetfParamsScimSchemasCyberark10Safe UrnIetfParamsScimSchemasCyberark10Safe `json:"urn:ietf:params:scim:schemas:cyberark:1.0:Safe"`
}

// Container Attribute
type UrnIetfParamsScimSchemasCyberark10Safe struct {
	NumberOfDaysRetention int    `json:"NumberOfDaysRetention"`
	ManagingCPM           string `json:"ManagingCPM"`
}

type Owner struct {
	Value   string `json:"value"`
	Ref     string `json:"$ref"`
	Display string `json:"display"`
}

type PrivilegedDataRef struct {
	Value   string `json:"value"`
	Ref     string `json:"$ref"`
	Display string `json:"display"`
}
