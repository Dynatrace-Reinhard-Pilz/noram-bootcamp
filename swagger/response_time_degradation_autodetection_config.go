/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Parameters of the response time degradation auto-detection. Required if the **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.   Violation of any criterion triggers an alert.
type ResponseTimeDegradationAutodetectionConfig struct {

	// Alert if the response time degrades by more than *X* milliseconds.
	ResponseTimeDegradationMilliseconds int32 `json:"responseTimeDegradationMilliseconds"`

	// Alert if the response time degrades by more than *X* %.
	ResponseTimeDegradationPercent int32 `json:"responseTimeDegradationPercent"`

	// Alert if the response time of the slowest 10% degrades by more than *X* milliseconds.
	SlowestResponseTimeDegradationMilliseconds int32 `json:"slowestResponseTimeDegradationMilliseconds"`

	// Alert if the response time of the slowest 10% degrades by more than *X* %.
	SlowestResponseTimeDegradationPercent int32 `json:"slowestResponseTimeDegradationPercent"`

	// Mininal service load to detect response time degradation.   Response time degradation of services with smaller load won't trigger alerts.
	LoadThreshold string `json:"loadThreshold"`
}
