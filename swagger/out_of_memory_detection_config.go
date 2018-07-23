/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Configuration of Java out-of-memory problems detection
type OutOfMemoryDetectionConfig struct {

	// Is detection enabled.
	Enabled bool `json:"enabled"`

	// Custom thresholds for Java out-of-memory detection. If not present (null) then default thresholds will be used.
	CustomThresholds *OutOfMemoryThresholds `json:"customThresholds,omitempty"`
}