# AutoTagRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Type of entities for which the tag rule should be evaluated and associated to if it matches. | [default to null]
**Enabled** | **bool** | Tag evaluation rule enabled/disabled | [default to null]
**ValueFormat** | **string** | Value format for the tag value. If specified, the format of the tag will be &lt;name&gt;:&lt;evaluatedValueFormat&gt;. You can use various placeholders in here like {ProcessGroup:DetectedName} which are evaluated at run-time. (Only some of those placeholders are allowed based on the specified &#39;type&#39;.) You can also specify a regex for extraction like this {ProcessGroup:DetectedName\\.*}. | [optional] [default to null]
**PropagationTypes** | **[]string** | Types of propagation to be done on a matching tag. Only some may be valid for the specified &#39;type&#39;. | [optional] [default to null]
**Conditions** | [**[]EntityRuleEngineCondition**](EntityRuleEngineCondition.md) | Actual conditions of the rule when the rule should match for an entity of the specified &#39;type&#39;.Only if all conditions match for an entity, the tag evaluation is positive and the tag is associated with the entity. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


