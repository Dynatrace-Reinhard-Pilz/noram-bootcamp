/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TagFilterDto struct {

	// The tag context.
	Context string `json:"context,omitempty"`

	// The tag key.
	Key string `json:"key"`

	// The tag value.
	Value string `json:"value,omitempty"`
}