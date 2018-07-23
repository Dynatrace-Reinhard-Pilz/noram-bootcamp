/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Custom thresholds for high CPU saturation detection. If not present (null) then default thresholds will be used.
type HighCpuSaturationThresholds struct {

	// Detect high CPU saturation when CPU usage is higher than N percent in 3 out of 5 samples.
	CpuSaturation int32 `json:"cpuSaturation"`
}
