/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Custom thresholds for Java out-of-threads detection. If not present (null) then default thresholds will be used.
type OutOfThreadsThresholds struct {

	// Detect Java out of threads when number of out-of-threads exceptions per minute is N or more.
	OutOfThreadsExceptionsNumber int32 `json:"outOfThreadsExceptionsNumber"`
}