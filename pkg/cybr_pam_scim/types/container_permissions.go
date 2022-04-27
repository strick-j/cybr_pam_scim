package types

type ContainerPermissions struct {
	Schemas      []string              `json:"schemas"`
	TotalResults int                   `json:"totalResults"`
	ItemsPerPage int                   `json:"itemsPerPage"`
	StartIndex   int                   `json:"startIndex"`
	Resources    []ContainerPermission `json:"Resources"`
}

type ContainerPermission struct {
	Container                                    ContainerRef                                 `json:"container"`
	User                                         UserRef                                      `json:"user,omitempty"`
	Group                                        GroupRef                                     `json:"group,omitempty"`
	Rights                                       []string                                     `json:"rights"`
	Schemas                                      []string                                     `json:"schemas"`
	Id                                           string                                       `json:"id"`
	ExternalId                                   string                                       `json:"externalId,omitempty"`
	Meta                                         Meta                                         `json:"meta"`
	UrnIetfParamsScimSchemasCyberark10SafeMember UrnIetfParamsScimSchemasCyberark10SafeMember `json:"urn:ietf:params:scim:schemas:cyberark:1.0:SafeMember"`
}

type ContainerRef struct {
	Value   string `json:"value,omitempty"`
	Ref     string `json:"$ref,omitempty"`
	Name    string `json:"name,omitempty"`
	Display string `json:"display,omitempty"`
}

type UrnIetfParamsScimSchemasCyberark10SafeMember struct {
	MembershipExpirationDate int    `json:"membershipExpirationDate,omitempty"`
	MemberType               string `json:"memberType,omitempty"`
	SearchIn                 string `json:"searchIn,omitempty"`
}
