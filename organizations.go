package datadog

type Organization struct {
	Billing      *OrganizationBilling      `json:"billing,omitempty"`
	Created      *string                   `json:"created,omitempty"`
	Description  *string                   `json:"description,omitempty"`
	Name         *string                   `json:"name,omitempty"`
	PublicId     *string                   `json:"public_id,omitempty"`
	Settings     *OrganizationSettings     `json:"settings,omitempty"`
	Subscription *OrganizationSubscription `json:"subscription,omitempty"`
}

type OrganizationBilling struct {
	Type *string `json:"type,omitempty"`
}

type OrganizationSettings struct {
	PrivateWidgetShare         *bool                                   `json:"private_widget_share,omitempty"`
	Saml                       *OrganizationSaml                       `json:"subscription,omitempty"`
	SamlAutocreateAccessRole   *string                                 `json:"saml_autocreate_access_role"`
	SamlAutocreateUsersDomains *OrganizationSamlAutocreateUsersDomains `json:"saml_autocreate_users_domains,omitempty"`
	SamlCanBeEnabled           *bool                                   `json:"saml_can_be_enabled,omitempty"`
	SamlIdpEndpoint            *string                                 `json:"saml_idp_endpoint,omitempty"`
	SamlIdpInitiatedLogin      *OrganizationSamlIdpInitiatedLogin      `json:"saml_idp_initiated_login,omitempty"`
	SamlIdpMetadataUploaded    *bool                                   `json:"saml_idp_metadata_uploaded,omitempty"`
	SamlLoginUrl               *string                                 `json:"saml_login_url,omitempty"`
	SamlStrictMode             *OrganizationSamlStrictMode             `json:"saml_strict_mode,omitempty"`
}

type OrganizationSaml struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type OrganizationSamlAutocreateUsersDomains struct {
	Domains []string `json:"domains,omitempty"`
	Enabled *bool    `json:"enabled,omitempty"`
}

type OrganizationSamlIdpInitiatedLogin struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type OrganizationSamlStrictMode struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type OrganizationSubscription struct {
	Type *string `json:"type,omitempty"`
}

// GetOrganizations returns a list of all orgs.
func (client *Client) GetOrganizations() ([]Organization, error) {
	var out struct {
		Organizations []Organization `json:"orgs"`
	}

	uri := "/v1/org"
	if err := client.doJsonRequest("GET", uri, nil, &out); err != nil {
		return nil, err
	}
	return out.Organizations, nil
}

// GetOrganization get org by public id.
func (client *Client) GetOrganization(publicId string) (*Organization, error) {
	var out struct {
		Organization Organization `json:"org"`
	}

	uri := "/v1/org/" + publicId
	if err := client.doJsonRequest("GET", uri, nil, &out); err != nil {
		return nil, err
	}
	return &out.Organization, nil
}
