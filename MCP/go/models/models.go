package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// TagDescriptionContract represents the TagDescriptionContract schema from the OpenAPI specification
type TagDescriptionContract struct {
	Name string `json:"name,omitempty"` // Resource name.
	TypeField string `json:"type,omitempty"` // Resource type for API Management resource.
	Id string `json:"id,omitempty"` // Resource ID.
}

// TagCreateUpdateParameters represents the TagCreateUpdateParameters schema from the OpenAPI specification
type TagCreateUpdateParameters struct {
	Properties TagContractProperties `json:"properties,omitempty"` // Tag contract Properties.
}

// TagDescriptionCollection represents the TagDescriptionCollection schema from the OpenAPI specification
type TagDescriptionCollection struct {
	Nextlink string `json:"nextLink,omitempty"` // Next page link if any.
	Value []TagDescriptionContract `json:"value,omitempty"` // Page values.
}

// TagCollection represents the TagCollection schema from the OpenAPI specification
type TagCollection struct {
	Nextlink string `json:"nextLink,omitempty"` // Next page link if any.
	Value []interface{} `json:"value,omitempty"` // Page values.
}

// TagContract represents the TagContract schema from the OpenAPI specification
type TagContract struct {
	Id string `json:"id,omitempty"` // Resource ID.
	Name string `json:"name,omitempty"` // Resource name.
	TypeField string `json:"type,omitempty"` // Resource type for API Management resource.
}

// TagContractProperties represents the TagContractProperties schema from the OpenAPI specification
type TagContractProperties struct {
	Displayname string `json:"displayName"` // Tag name.
}

// TagDescriptionContractProperties represents the TagDescriptionContractProperties schema from the OpenAPI specification
type TagDescriptionContractProperties struct {
	Description string `json:"description,omitempty"` // Description of the Tag.
	Externaldocsdescription string `json:"externalDocsDescription,omitempty"` // Description of the external resources describing the tag.
	Externaldocsurl string `json:"externalDocsUrl,omitempty"` // Absolute URL of external resources describing the tag.
}

// TagDescriptionCreateParameters represents the TagDescriptionCreateParameters schema from the OpenAPI specification
type TagDescriptionCreateParameters struct {
	Properties TagDescriptionBaseProperties `json:"properties,omitempty"` // Parameters supplied to the Create TagDescription operation.
}

// TagDescriptionBaseProperties represents the TagDescriptionBaseProperties schema from the OpenAPI specification
type TagDescriptionBaseProperties struct {
	Description string `json:"description,omitempty"` // Description of the Tag.
	Externaldocsdescription string `json:"externalDocsDescription,omitempty"` // Description of the external resources describing the tag.
	Externaldocsurl string `json:"externalDocsUrl,omitempty"` // Absolute URL of external resources describing the tag.
}
