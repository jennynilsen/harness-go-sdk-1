/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// FeatureState : The state of a flag either off or on
type FeatureState string

// List of FeatureState
const (
	ON_FeatureState  FeatureState = "on"
	OFF_FeatureState FeatureState = "off"
)