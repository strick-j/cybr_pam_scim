package types

import (
	"time"
)

// Generic Shared Fields //////////////////////////////////////////////////////////////////

// Used in all responses
type Meta struct {
	ResourceType string    `json:"resourceType,omitempty"`
	Created      time.Time `json:"created,omitempty"`
	LastModified time.Time `json:"lastModified,omitempty"`
	Location     string    `json:"location,omitempty"`
	Version      string    `json:"version,omitempty"`
}

// Used with "ContainerPermissions and PrivilegedDataPermissions"
type PrivilegedDataRef struct {
	Value   string `json:"value,omitempty"`
	Ref     string `json:"$ref,omitempty"`
	Display string `json:"display,omitempty"`
	Type    string `json:"type,omitempty"`
}

// Used with "ContainerPermissions and PrivilegedDataPermissions"
type GroupRef struct {
	Value   string `json:"value,omitempty"`
	Ref     string `json:"$ref,omitempty"`
	Display string `json:"display,omitempty"`
}

// Used with "ContainerPermissions and PrivilegedDataPermissions"
type UserRef struct {
	Value   string `json:"value,omitempty"`
	Ref     string `json:"$ref,omitempty"`
	Display string `json:"display,omitempty"`
}

// SCIM Service Provider Config //////////////////////////////////////////////////////////////////
type ScimConfig struct {
	Schemas               []string                `json:"schemas"`
	Patch                 Patch                   `json:"patch"`
	Bulk                  Bulk                    `json:"bulk"`
	Filter                Filter                  `json:"filter"`
	ChangePassword        ChangePassword          `json:"changePassword"`
	Sort                  Sort                    `json:"sort"`
	Etag                  Etag                    `json:"etag"`
	AuthenticationSchemes []AuthenticationSchemes `json:"authenticationSchemes"`
	Meta                  Meta                    `json:"meta"`
}

type Patch struct {
	Supported bool `json:"supported"`
}

type Bulk struct {
	Supported      bool `json:"supported"`
	MaxOperations  int  `json:"maxOperations"`
	MaxPayloadSize int  `json:"maxPayloadSize"`
}

type Filter struct {
	Supported  bool `json:"supported"`
	MaxResults int  `json:"maxResults"`
}

type ChangePassword struct {
	Supported bool `json:"supported"`
}

type Sort struct {
	Supported bool `json:"supported"`
}

type Etag struct {
	Supported bool `json:"supported"`
}

type AuthenticationSchemes struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// SCIM Resource Types ///////////////////////////////////////////
type ResourceTypes struct {
	Schemas      []string       `json:"schemas"`
	TotalResults int            `json:"totalResults"`
	ItemsPerPage int            `json:"itemsPerPage"`
	StartIndex   int            `json:"startIndex"`
	Resources    []ResourceType `json:"Resources"`
}

type ResourceType struct {
	Name             string             `json:"name"`
	Endpoint         string             `json:"endpoint"`
	Schema           string             `json:"schema"`
	SchemaExtensions []SchemaExtensions `json:"schemaExtensions,omitempty"`
	Schemas          []string           `json:"schemas"`
	Id               string             `json:"id"`
	Meta             Meta               `json:"meta"`
}

type SchemaExtensions struct {
	Schema   string `json:"schema"`
	Required bool   `json:"required"`
}

// SCIM Schemas ///////////////////////////////////////////
type Schemas struct {
	Schemas      []string `json:"schemas"`
	TotalResults int      `json:"totalResults"`
	ItemsPerPage int      `json:"itemsPerPage"`
	StartIndex   int      `json:"startIndex"`
	Resources    []Schema `json:"Resources"`
}

type Schema struct {
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"`
	Attributes  []Attributes `json:"attributes,omitempty"`
	Id          string       `json:"id"`
	Meta        Meta         `json:"meta"`
}

type Attributes struct {
	Name          string          `json:"name"`
	Type          string          `json:"type"`
	MultiValued   bool            `json:"multiValued"`
	Required      bool            `json:"required"`
	CaseExact     bool            `json:"caseExact,omitempty"`
	SubAttributes []SubAttributes `json:"subAttributes,omitempty"`
	Mutability    string          `json:"mutability,omitempty"`
	Returned      string          `json:"returned,omitempty"`
}

type SubAttributes struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	MultiValued bool   `json:"multiValued"`
	Required    bool   `json:"required"`
	CaseExact   bool   `json:"caseExact"`
}
