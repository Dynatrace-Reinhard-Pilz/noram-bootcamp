/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Type-specific key to identify the data for which we want to match against.This is an abstract version of the key, in its place use one of the actual implementations for the allowed types: * STRING -> EntityRuleEngineDynamicStringConditionKey * PROCESS_PREDEFINED_METADATA_KEY -> EntityRuleEngineDynamicPgIdCalcInputConditionKey * PROCESS_CUSTOM_METADATA_KEY -> EntityRuleEngineDynamicCustomPgMetadataConditionKey * null (no type is set) -> EntityRuleEngineStaticConditionKey
type EntityRuleEngineConditionKey struct {

	// Defines the attribute used for the filter. Only some are allowed based on the feature it is used for and the specified 'type' outside this condition.
	Attribute string `json:"attribute"`

	Type_ string `json:"type,omitempty"`
}
