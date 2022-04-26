package types

type Users struct {
	Schemas      []string `json:"schemas"`
	TotalResults int      `json:"totalResults"`
	ItemsPerPage int      `json:"itemsPerPage"`
	StartIndex   int      `json:"startIndex"`
	Resources    []User   `json:"Resources"`
}

type User struct {
	UserName                                          string                                            `json:"userName"`
	Name                                              Name                                              `json:"name,omitempty"`
	DisplayName                                       string                                            `json:"displayName,omitempty"`
	NickName                                          string                                            `json:"nickName,omitempty"`
	ProfileUrl                                        string                                            `json:"profileUrl,omitempty"`
	Title                                             string                                            `json:"title,omitempty"`
	UserType                                          string                                            `json:"userType,omitempty"`
	PreferredLanguage                                 string                                            `json:"preferredLanguage,omitempty"`
	Locale                                            string                                            `json:"locale,omitempty"`
	Timezone                                          string                                            `json:"timezone,omitempty"`
	Active                                            bool                                              `json:"active,omitempty"`
	Password                                          string                                            `json:"-,omitempty"`
	Emails                                            []Emails                                          `json:"emails,omitempty"`
	PhoneNumbers                                      []PhoneNumbers                                    `json:"phoneNumbers,omitempty"`
	Ims                                               []Ims                                             `json:"ims,omitempty"`
	Photos                                            []Photos                                          `json:"photos,omitempty"`
	Addresses                                         []Addresses                                       `json:"addresses,omitempty"`
	Groups                                            []UsersGroups                                     `json:"groups,omitempty"`
	Entitlements                                      []string                                          `json:"entitlements,omitempty"`
	Roles                                             []Roles                                           `json:"roles,omitempty"`
	X509Certificates                                  []X509Certificates                                `json:"x509Certificates,omitempty"`
	Schemas                                           []string                                          `json:"schemas"`
	Id                                                string                                            `json:"id"`
	ExternalId                                        string                                            `json:"externalId,omitempty"`
	Meta                                              Meta                                              `json:"meta"`
	UrnIetfParamsScimSchemasPam10LinkedObject         UrnIetfParamsScimSchemasPam10LinkedObject         `json:"urn:ietf:params:scim:schemas:pam:1.0:LinkedObject,omitempty"`
	UrnIetfParamsScimSchemasExtensionEnterprise20User UrnIetfParamsScimSchemasExtensionEnterprise20User `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User,omitempty"`
	UrnIetfParamsScimSchemasCyberark10User            UrnIetfParamsScimSchemasCyberark10User            `json:"urn:ietf:params:scim:schemas:cyberark:1.0:User,omitempty"`
	UrnScimSchemasExtensionCustom20                   UrnScimSchemasExtensionCustom20                   `json:"urn:scim:schemas:extension:custom:2.0,omitempty"`
}

type UsersGroups struct {
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
	Value   string `json:"value,omitempty"`
	Ref     string `json:"$ref,omitempty"`
	Display string `json:"display,omitempty"`
}

type Name struct {
	Formatted       string `json:"formatted,omitempty"`
	GivenName       string `json:"givenName,omitempty"`
	FamilyName      string `json:"familyName,omitempty"`
	MiddleName      string `json:"middleName,omitempty"`
	HonorificPrefix string `json:"honorificPrefix,omitempty"`
	HonorificSuffix string `json:"honorificSuffix,omitempty"`
}

type Emails struct {
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Ref     string `json:"$ref,omitempty"`
}

type PhoneNumbers struct {
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Ref     string `json:"$ref,omitempty"`
}

type Ims struct {
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Ref     string `json:"$ref,omitempty"`
}

type Photos struct {
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Ref     string `json:"$ref,omitempty"`
}

type Addresses struct {
	Formatted     string `json:"formatted,omitempty"`
	StreetAddress string `json:"streetAddress,omitempty"`
	Locality      string `json:"locality,omitempty"`
	Region        string `json:"region,omitempty"`
	PostalCode    string `json:"postalCode,omitempty"`
	Country       string `json:"country,omitempty"`
}

type Roles struct {
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Ref     string `json:"$ref,omitempty"`
}

type X509Certificates struct {
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Ref     string `json:"$ref,omitempty"`
}

type UrnIetfParamsScimSchemasPam10LinkedObject struct {
	Source           string `json:"source,omitempty"`
	NativeIdentifier string `json:"nativeIdentifier,omitempty"`
}

type UrnIetfParamsScimSchemasExtensionEnterprise20User struct {
	EmployeeNumber string    `json:"employeeNumber,omitempty"`
	CostCenter     string    `json:"costCenter,omitempty"`
	Organization   string    `json:"organization,omitempty"`
	Division       string    `json:"division,omitempty"`
	Department     string    `json:"department,omitempty"`
	Manager        []Manager `json:"manager,omitempty"`
}

type Manager struct {
	Value       string `json:"value,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Ref         string `json:"$ref,omitempty"`
}

type UrnIetfParamsScimSchemasCyberark10User struct {
	AuthenticationMethod  []string `json:"authenticationMethod,omitempty"`
	ExpiryDate            int64    `json:"expiryDate,omitempty"`
	ChangePassOnNextLogon bool     `json:"changePassOnNextLogon,omitempty"`
	PasswordNeverExpires  bool     `json:"passwordNeverExpires,omitempty"`
	DistinguishedName     string   `json:"distinguishedName,omitempty"`
	DirectoryType         string   `json:"directoryType,omitempty"`
}

type UrnScimSchemasExtensionCustom20 struct {
}
