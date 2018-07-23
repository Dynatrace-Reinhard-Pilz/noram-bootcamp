/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// A condition holds the settings for how to execute matching logic for an entity.
type EntityRuleEngineCondition struct {

	// Type-specific key to identify the data for which we want to match against.
	Key *EntityRuleEngineConditionKey `json:"key"`

	// Type-specific comparison information for attributes.
	ComparisonInfo *EntityRuleEngineComparisonInfoObjectObject `json:"comparisonInfo"`
}