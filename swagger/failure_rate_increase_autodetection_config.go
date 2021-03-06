/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Parameters of the failure rate increase auto-detection. Required if the **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.   Both absolute and relative threshold must exceed to trigger an alert.   Example: If the expected error rate is 1.5%, and you set absolute increase of 1%, and relative increase of 50%, the thresholds will be:  Absolute: 1.5% + **1%** = 2.5%  Relative: 1.5% + 1.5% * **50%** = 2.25%
type FailureRateIncreaseAutodetectionConfig struct {

	// Absolute increase of failing service calls to trigger an alert, %.
	FailingServiceCallPercentageIncreaseAbsolute int32 `json:"failingServiceCallPercentageIncreaseAbsolute"`

	// Relative increase of failing service calls to trigger an alert, %.
	FailingServiceCallPercentageIncreaseRelative int32 `json:"failingServiceCallPercentageIncreaseRelative"`
}
