package types

type Users struct {
	Schemas      []string `json:"schemas"`
	TotalResults int      `json:"totalResults"`
	ItemsPerPage int      `json:"itemsPerPage"`
	StartIndex   int      `json:"startIndex"`
	Resources    []User   `json:"Resources"`
}

type User struct {
	Active                                            bool                                              `json:"active,omitempty"`
	Description                                       string                                            `json:"description,omitempty"`
	DisplayName                                       string                                            `json:"displayName,omitempty"`
	Emails                                            []Emails                                          `json:"emails,omitempty"`
	PhoneNumbers                                      []PhoneNumbers                                    `json:"phoneNumbers,omitempty"`
	Entitlements                                      []string                                          `json:"entitlements,omitempty"`
	Groups                                            []UsersGroups                                     `json:"groups,omitempty"`
	ID                                                string                                            `json:"id,omitempty"`
	Meta                                              Meta                                              `json:"meta,omitempty"`
	Name                                              Name                                              `json:"name,omitempty"`
	Password                                          string                                            `json:"password,omitempty"`
	Schemas                                           []string                                          `json:"schemas,omitempty"`
	UserName                                          string                                            `json:"userName,omitempty"`
	UserType                                          string                                            `json:"userType,omitempty"`
	UrnIetfParamsScimSchemasPam10LinkedObject         UrnIetfParamsScimSchemasPam10LinkedObject         `json:"urn:ietf:params:scim:schemas:pam:1.0:LinkedObject,omitempty"`
	UrnIetfParamsScimSchemasExtensionEnterprise20User UrnIetfParamsScimSchemasExtensionEnterprise20User `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User,omitempty"`
	UrnIetfParamsScimSchemasCyberark10User            UrnIetfParamsScimSchemasCyberark10User            `json:"urn:ietf:params:scim:schemas:cyberark:1.0:User"`
}

type UsersGroups struct {
	Type    string `json:"type"`
	Value   string `json:"value"`
	Ref     string `json:"$ref"`
	Display string `json:"display"`
}

type Name struct {
	Formatted string `json:"formatted"`
	GivenName string `json:"givenName"`
}

type Emails struct {
	Type    string `json:"type"`
	Primary bool   `json:"primary"`
	Value   string `json:"value"`
}

type PhoneNumbers struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type UrnIetfParamsScimSchemasPam10LinkedObject struct {
	Source           string `json:"source"`
	NativeIdentifier string `json:"nativeIdentifier"`
}

type UrnIetfParamsScimSchemasExtensionEnterprise20User struct {
	Organization string `json:"organization"`
}

type UrnIetfParamsScimSchemasCyberark10User struct {
	AuthenticationMethod  []string `json:"authenticationMethod"`
	ExpiryDate            int64    `json:"expiryDate"`
	ChangePassOnNextLogon bool     `json:"changePassOnNextLogon"`
	PasswordNeverExpires  bool     `json:"passwordNeverExpires"`
	DirectoryType         string   `json:"directoryType"`
}
