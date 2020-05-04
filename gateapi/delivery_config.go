/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DeliveryConfig struct {
	ApiVersion string `json:"apiVersion,omitempty"`

	Application string `json:"application,omitempty"`

	Environments []Environment `json:"environments,omitempty"`

	ServiceAccount string `json:"serviceAccount,omitempty"`

	Name string `json:"name,omitempty"`

	Artifacts []Mapstringobject `json:"artifacts,omitempty"`
}
