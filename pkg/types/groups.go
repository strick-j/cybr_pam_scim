package types

type Groups struct {
	Schemas      []string `json:"schemas"`
	TotalResults int      `json:"totalResults"`
	Resources    []Groups `json:"Resources"`
	StartIndex   int      `json:"startIndex,omitempty"`
	ItemsPerPage int      `json:"itemsPerPage"`
}

type Group struct {
	DisplayName                             string                                  `json:"displayName"`
	Members                                 []Members                               `json:"members,omitempty"`
	Schemas                                 []string                                `json:"schemas"`
	Id                                      string                                  `json:"id"`
	Meta                                    Meta                                    `json:"meta"`
	ExternalId                              string                                  `json:"externalId,omitempty"`
	UrnIetfParamsScimSchemasCyberark10Group UrnIetfParamsScimSchemasCyberark10Group `json:"urn:ietf:params:scim:schemas:cyberark:1.0:Group,omitempty"`
}

type Members struct {
	Value   string `json:"value,omitempty"`
	Type    string `json:"type,omitempty"`
	Ref     string `json:"$ref,omitempty"`
	Display string `json:"display,omitempty"`
}

type UrnIetfParamsScimSchemasCyberark10Group struct {
	DirectoryType string `json:"directoryType,omitempty"`
	DirectoryName string `json:"directoryName,omitempty"`
}
