/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type DelegateProfileDetailsNg struct {
	Uuid string `json:"uuid,omitempty"`
	AccountId string `json:"accountId,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Primary bool `json:"primary,omitempty"`
	ApprovalRequired bool `json:"approvalRequired,omitempty"`
	StartupScript string `json:"startupScript,omitempty"`
	ScopingRules []ScopingRuleDetailsNg `json:"scopingRules,omitempty"`
	Selectors []string `json:"selectors,omitempty"`
	CreatedBy *EmbeddedUserDetails `json:"createdBy,omitempty"`
	LastUpdatedBy *EmbeddedUserDetails `json:"lastUpdatedBy,omitempty"`
	CreatedAt int64 `json:"createdAt,omitempty"`
	LastUpdatedAt int64 `json:"lastUpdatedAt,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	NumberOfDelegates int64 `json:"numberOfDelegates,omitempty"`
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
}