package types

type PrivilegedDataPermissions struct {
	Schemas      []string                   `json:"schemas"`
	TotalResults int                        `json:"totalResults"`
	ItemsPerPage int                        `json:"itemsPerPage"`
	StartIndex   int                        `json:"startIndex"`
	Resources    []PrivilegedDataPermission `json:"Resources"`
}

type PrivilegedDataPermission struct {
	PrivilegedData PrivilegedDataRef `json:"privilegedData"`
	User           UserRef           `json:"user,omitempty"`
	Group          GroupRef          `json:"group,omitempty"`
	Rights         []string          `json:"rights"`
	Schemas        []string          `json:"schemas"`
	Id             string            `json:"id"`
	ExternalId     string            `json:"externalId,omitempty"`
	Meta           Meta              `json:"meta"`
}
