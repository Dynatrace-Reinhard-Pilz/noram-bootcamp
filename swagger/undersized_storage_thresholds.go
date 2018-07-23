/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Custom thresholds for undersized storage device detection. If not present (null) then default thresholds will be used.
type UndersizedStorageThresholds struct {

	// Detect undersized storage device always when average queue command latency is higher than N milliseconds in 3 out of 5 samples.
	AverageQueueCommandLatency int32 `json:"averageQueueCommandLatency"`

	// Detect undersized storage device always when peak queue command latency is higher than N milliseconds in 3 out of 5 samples.
	PeakQueueCommandLatency int32 `json:"peakQueueCommandLatency"`
}
