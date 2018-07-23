/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Configuration of low datastore free space detection
type LowDatastoreSpaceDetectionConfig struct {

	// Is detection enabled.
	Enabled bool `json:"enabled"`

	// Custom thresholds for low datastore free space detection. If not present (null) then default thresholds will be used.
	CustomThresholds *LowDatastoreSpaceThresholds `json:"customThresholds,omitempty"`
}
