package types

type PrivilegedDatas struct {
	Schemas      []string          `json:"schemas"`
	TotalResults int               `json:"totalResults"`
	ItemsPerPage int               `json:"itemsPerPage"`
	StartIndex   int               `json:"startIndex"`
	Resources    []PrivilegedDatas `json:"Resources"`
}

type PrivilegedData struct {
	Name                                             string                                           `json:"name"`
	Description                                      string                                           `json:"description,omitempty"`
	Type                                             string                                           `json:"type"`
	Schemas                                          []string                                         `json:"schemas"`
	Id                                               string                                           `json:"id"`
	ExternalId                                       string                                           `json:"externalId,omitempty"`
	Meta                                             Meta                                             `json:"meta"`
	UrnIetfParamsScimSchemasCyberark10PrivilegedData UrnIetfParamsScimSchemasCyberark10PrivilegedData `json:"urn:ietf:params:scim:schemas:cyberark:1.0:PrivilegedData"`
}

type Properties struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UrnIetfParamsScimSchemasCyberark10PrivilegedData struct {
	Safe       string       `json:"safe,omitempty"`
	Folder     string       `json:"folder,omitempty"`
	Password   string       `json:"password,omitempty"`
	Properties []Properties `json:"properties,omitempty"`
}

// Used in Privileged Data PATCH functions
type Value struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Used in Privileged Data PATCH functions
type Operations struct {
	Op    string  `json:"op"`
	Path  string  `json:"path"`
	Value []Value `json:"value"`
}
