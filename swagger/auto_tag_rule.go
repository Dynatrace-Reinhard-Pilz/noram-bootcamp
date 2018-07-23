/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// A rule defines when to associate a tag with an entity. Each rule is evaluated independently of all other rules.
type AutoTagRule struct {

	// Type of entities for which the tag rule should be evaluated and associated to if it matches.
	Type_ string `json:"type"`

	// Tag evaluation rule enabled/disabled
	Enabled bool `json:"enabled"`

	// Value format for the tag value. If specified, the format of the tag will be <name>:<evaluatedValueFormat>. You can use various placeholders in here like {ProcessGroup:DetectedName} which are evaluated at run-time. (Only some of those placeholders are allowed based on the specified 'type'.) You can also specify a regex for extraction like this {ProcessGroup:DetectedName\\.*}.
	ValueFormat string `json:"valueFormat,omitempty"`

	// Types of propagation to be done on a matching tag. Only some may be valid for the specified 'type'.
	PropagationTypes []string `json:"propagationTypes,omitempty"`

	// Actual conditions of the rule when the rule should match for an entity of the specified 'type'.Only if all conditions match for an entity, the tag evaluation is positive and the tag is associated with the entity.
	Conditions []EntityRuleEngineCondition `json:"conditions"`
}
