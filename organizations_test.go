package datadog

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOrganization(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/organizations_response.json")
		if err != nil {
			t.Fatal(err)
		}
		w.Write(response)
	}))
	defer ts.Close()

	datadogClient := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	organization, err := datadogClient.GetOrganization("xxx-xxx-xxx")
	if err != nil {
		t.Fatal(err)
	}

	expectedCreated := "2020-08-01 10:00:00"
	if created := organization.GetCreated(); created != expectedCreated {
		t.Fatalf("expect created %s. Got %s", expectedCreated, created)
	}

	expectedDescription := "description"
	if description := organization.GetDescription(); description != expectedDescription {
		t.Fatalf("expect description %s. Got %s", expectedDescription, description)
	}

	expectedName := "test"
	if name := organization.GetName(); name != expectedName {
		t.Fatalf("expect name %s. Got %s", expectedName, name)
	}

	expectedPublicId := "xxx-xxx-xxx"
	if publicId := organization.GetPublicId(); publicId != expectedPublicId {
		t.Fatalf("expect public_id %s. Got %s", expectedPublicId, publicId)
	}

	billing := organization.GetBilling()

	expectedBillingType := "parent_billing"
	if billingType := billing.GetType(); billingType != expectedBillingType {
		t.Fatalf("expect billing.type  %s. Got %s", expectedBillingType, billingType)
	}

	settings := organization.GetSettings()

	expectedSettingsPrivateWidgetShare := false
	if settingsPrivateWidgetShare := settings.GetPrivateWidgetShare(); settingsPrivateWidgetShare != expectedSettingsPrivateWidgetShare {
		t.Fatalf("expect settings.private_widget_share  %v. Got %v", expectedSettingsPrivateWidgetShare, settingsPrivateWidgetShare)
	}

	expectedSettingsSamlAutocreateAccessRole := "st"
	if settingsSamlAutocreateAccessRole := settings.GetSamlAutocreateAccessRole(); settingsSamlAutocreateAccessRole != expectedSettingsSamlAutocreateAccessRole {
		t.Fatalf("expect settings.saml_autocreate_access_role  %s. Got %s", expectedSettingsSamlAutocreateAccessRole, settingsSamlAutocreateAccessRole)
	}

	expectedSettingsSamlCanBeEnabled := false
	if settingsSamlCanBeEnabled := settings.GetSamlCanBeEnabled(); settingsSamlCanBeEnabled != expectedSettingsSamlCanBeEnabled {
		t.Fatalf("expect settings.saml_can_be_enabled  %v. Got %v", expectedSettingsSamlCanBeEnabled, settingsSamlCanBeEnabled)
	}

	expectedSettingsSamlIdpEndpoint := "https://my.saml.endpoint"
	if settingsSamlIdpEndpoint := settings.GetSamlIdpEndpoint(); settingsSamlIdpEndpoint != expectedSettingsSamlIdpEndpoint {
		t.Fatalf("expect settings.saml_idp_endpoint  %s. Got %s", expectedSettingsSamlIdpEndpoint, settingsSamlIdpEndpoint)
	}

	expectedSettingsSamlIdpMetadataUploaded := false
	if settingsSamlIdpMetadataUploaded := settings.GetSamlIdpMetadataUploaded(); settingsSamlIdpMetadataUploaded != expectedSettingsSamlIdpMetadataUploaded {
		t.Fatalf("expect settings.saml_idp_metadata_uploaded  %v. Got %v", expectedSettingsSamlIdpMetadataUploaded, settingsSamlIdpMetadataUploaded)
	}

	expectedSettingsSamlLoginUrl := "https://my.saml.login.url"
	if settingsSamlLoginUrl := settings.GetSamlLoginUrl(); settingsSamlLoginUrl != expectedSettingsSamlLoginUrl {
		t.Fatalf("expect settings.saml_login_url  %s. Got %s", expectedSettingsSamlLoginUrl, settingsSamlLoginUrl)
	}

	expectedSettingsSaml := false
	if settingsSaml := settings.Saml.GetEnabled(); settingsSaml != expectedSettingsSaml {
		t.Fatalf("expect settings.saml.enabled  %v. Got %v", expectedSettingsSaml, settingsSaml)
	}

	domains := settings.SamlAutocreateUsersDomains.Domains
	expectedSettingsSamlAutocreateUsersDomainsDomains := 1
	if cnt := len(domains); cnt != expectedSettingsSamlAutocreateUsersDomainsDomains {
		t.Fatalf("expect domains count should be  %d. Got %d", expectedSettingsSamlAutocreateUsersDomainsDomains, cnt)
	}

	expectedSettingsSamlAutocreateUsersDomainsEnabled := false
	if settingsSamlAutocreateUsersDomainsEnabled := settings.SamlAutocreateUsersDomains.GetEnabled(); settingsSamlAutocreateUsersDomainsEnabled != expectedSettingsSamlAutocreateUsersDomainsEnabled {
		t.Fatalf("expect settings.saml_autocreate_users_domains.enabled  %v. Got %v", expectedSettingsSamlAutocreateUsersDomainsEnabled, settingsSamlAutocreateUsersDomainsEnabled)
	}

	expectedSettingsSamlIdpInitiatedLoginEnabled := false
	if settingsSamlIdpInitiatedLoginEnabled := settings.SamlIdpInitiatedLogin.GetEnabled(); settingsSamlIdpInitiatedLoginEnabled != expectedSettingsSamlIdpInitiatedLoginEnabled {
		t.Fatalf("expect settings.saml_idp_initiated_login.enalbed  %v. Got %v", expectedSettingsSamlAutocreateUsersDomainsEnabled, settingsSamlIdpInitiatedLoginEnabled)
	}

	expectedSettingsSamlStrictMode := false
	if settingsSamlStrictModeEnabled := settings.SamlStrictMode.GetEnabled(); settingsSamlStrictModeEnabled != expectedSettingsSamlStrictMode {
		t.Fatalf("expect settings.saml_strict_mode  %v. Got %v", expectedSettingsSamlStrictMode, expectedSettingsSamlStrictMode)
	}

	subscription := organization.GetSubscription()

	expectedSubscriptionType := "pro"
	if subscriptionType := subscription.GetType(); subscriptionType != expectedSubscriptionType {
		t.Fatalf("expect subscription.type  %s. Got %s", expectedSubscriptionType, subscriptionType)
	}
}
