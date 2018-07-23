/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Custom thresholds for low free storage space detection on RDS. If not present (null) then default thresholds will be used.
type RdsLowStorageThresholds struct {

	// Detect low storage when free storage space divided by allocated storage is lower than N percent in 3 out of 5 samples.
	FreeStoragePercentage int32 `json:"freeStoragePercentage"`
}
