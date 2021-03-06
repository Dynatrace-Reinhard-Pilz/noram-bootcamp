/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Custom thresholds for slow disk writes and reads detection. If not present (null) then default thresholds will be used.
type DiskLowSpaceThresholds struct {

	// Detect slow disk writes and reads when ead time and write time is higher than N milliseconds.
	WriteAndReadTime int32 `json:"writeAndReadTime"`
}
