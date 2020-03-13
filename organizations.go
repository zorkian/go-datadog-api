package datadog

type OrganizationResponse struct {
	Data *Organization `json:"data,omitempty"`
}

// TODO: Fill this out with everything you can get back when getting an organization
type Organization struct {
	Type *string `json:"type,omitempty"`
	Id   *string `json:"id,omitempty"`
}
