/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Custom thresholds for high error rate detection on Lambda service. If not present (null) then default thresholds will be used.
type LambdaHighErrorRateThresholds struct {

	// Detect high error rate when failed invocations rate is higher than N percent in 3 out of 5 samples.
	FailedInvocationsRate int32 `json:"failedInvocationsRate"`
}