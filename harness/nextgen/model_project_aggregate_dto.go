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

type ProjectAggregateDto struct {
	ProjectResponse *ProjectResponse `json:"projectResponse"`
	Organization *Organization `json:"organization,omitempty"`
	HarnessManagedOrg bool `json:"harnessManagedOrg,omitempty"`
	Admins []UserMetadataDto `json:"admins,omitempty"`
	Collaborators []UserMetadataDto `json:"collaborators,omitempty"`
}