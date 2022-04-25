package types

type Groups struct {
	Schemas      []string `json:"schemas"`
	TotalResults int      `json:"totalResults"`
	Resources    []Groups `json:"Resources"`
}

type Group struct {
	DisptlayName                            string                                  `json:"disptlayName,omitempty"`
	Members                                 []Members                               `json:"members"`
	Schemas                                 []string                                `json:"schemas"`
	ID                                      string                                  `json:"id"`
	Meta                                    Meta                                    `json:"meta"`
	DisplayName                             string                                  `json:"displayName,omitempty"`
	ExternalID                              string                                  `json:"externalId,omitempty"`
	UrnIetfParamsScimSchemasCyberark10Group UrnIetfParamsScimSchemasCyberark10Group `json:"urn:ietf:params:scim:schemas:cyberark:1.0:Group,omitempty"`
}

type Members struct {
	Value   string `json:"value"`
	Ref     string `json:"$ref"`
	Display string `json:"display"`
}

type UrnIetfParamsScimSchemasCyberark10Group struct {
	DirectoryType string `json:"directoryType"`
	DirectoryName string `json:"directoryName"`
}
