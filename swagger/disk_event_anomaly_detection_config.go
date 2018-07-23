/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DiskEventAnomalyDetectionConfig struct {

	// Metadata useful for debugging.
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`

	// The ID of the disk event.
	Id string `json:"id,omitempty"`

	// The name of the disk event.
	Name string `json:"name"`

	// Whether the disk event is enabled.
	Enabled bool `json:"enabled"`

	// The metric to monitor.
	Metric string `json:"metric"`

	// The event threshold. If the metric is `LowDiskSpace` or `LowInodes`, this is a percentage. For `ReadTimeExceeding` or `WriteTimeExceeding`, milliseconds.
	Threshold float64 `json:"threshold"`

	// The minimum number of violating samples for an event to be triggered. Must not exceed `samples`.
	ViolatingSamples int32 `json:"violatingSamples"`

	// The number of samples that are inspected.
	Samples int32 `json:"samples"`

	// Only apply this rule to disks matching the given name.
	DiskNameFilter *DiskNameFilterDto `json:"diskNameFilter,omitempty"`

	// Only apply this rule to hosts having these tags.
	TagFilters []TagFilterDto `json:"tagFilters,omitempty"`
}
