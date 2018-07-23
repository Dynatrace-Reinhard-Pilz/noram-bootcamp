/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Fixed thresholds for response time degradation detection. Required if the **detectionMode** is `DETECT_USING_FIXED_THRESHOLDS`. Not applicable otherwise.
type ResponseTimeDegradationThresholdConfig struct {

	// Response time during any 5-minute period to trigger an alert, milliseconds.
	ResponseTimeThresholdMilliseconds int32 `json:"responseTimeThresholdMilliseconds"`

	// Response time of the 10% slowest during any 5-minute period to trigger an alert, milliseconds.
	SlowestResponseTimeThresholdMilliseconds int32 `json:"slowestResponseTimeThresholdMilliseconds"`

	// Mininal service load to detect response time degradation.   Response time degradation of services with smaller load won't trigger alerts.
	LoadThreshold string `json:"loadThreshold"`

	// Sensitivity of the threshold.   With `low` sensitivity high statistical confidence is used, so brief violations (for example, due to a surge in load) won't trigger alerts.   With `high` sensitivity no statistical confidence is used. Each violation triggers alert.
	Sensitivity string `json:"sensitivity"`
}
