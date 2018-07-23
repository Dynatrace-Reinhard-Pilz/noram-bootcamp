/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type EntityRuleEngineApplicationTypeComparisonInfo struct {

	// Operator comparing the extracted value to the comparison value.
	Operator string `json:"operator"`

	// Holds the actual value which should be used for comparison. Other parameters of this condition define how this value should be compared to actually stored values.
	Value string `json:"value,omitempty"`

	// Negate the comparison.
	Negate bool `json:"negate"`
}