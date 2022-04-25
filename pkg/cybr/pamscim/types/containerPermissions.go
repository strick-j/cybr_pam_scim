package types

type ScimContainerPermissions struct {
	Schemas      []string               `json:"schemas"`
	TotalResults int                    `json:"totalResults"`
	ItemsPerPage int                    `json:"itemsPerPage"`
	StartIndex   int                    `json:"startIndex"`
	Resources    []ContainerPermissions `json:"Resources"`
}

type ContainerPermissions struct {
	Container ContainerRef `json:"container"`
	User      UserRef      `json:"user,omitempty"`
	Rights    []string     `json:"rights"`
	Schemas   []string     `json:"schemas"`
	ID        string       `json:"id"`
	Meta      Meta         `json:"meta"`
	Group     GroupRef     `json:"group,omitempty"`
}

type ContainerRef struct {
	Value   string `json:"value"`
	Ref     string `json:"$ref"`
	Name    string `json:"name"`
	Display string `json:"display"`
}

type GroupRef struct {
	Value   string `json:"value"`
	Ref     string `json:"$ref"`
	Display string `json:"display"`
}

type UserRef struct {
	Value   string `json:"value"`
	Ref     string `json:"$ref"`
	Display string `json:"display"`
}
