/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type WebApplicationConfig struct {

	// Metadata useful for debugging.
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`

	// The unique id of the web application, automatically generated by the dynatrace server.
	Identifier string `json:"identifier,omitempty"`

	// The name of the Web Application
	Name string `json:"name"`

	// Real User Monitoring enabled/disabled.
	RealUserMonitoringEnabled bool `json:"realUserMonitoringEnabled"`

	// User session percentage for cost and traffic control (decimal number with three fractional digits).
	CostControlUserSessionPercentage float32 `json:"costControlUserSessionPercentage"`

	// The Key Performance Metric for load actions
	LoadActionKeyPerformanceMetric string `json:"loadActionKeyPerformanceMetric"`

	// The Key Performance Metric for xhr actions
	XhrActionKeyPerformanceMetric string `json:"xhrActionKeyPerformanceMetric"`

	// The apdex settings for load actions
	LoadActionApdexSettings *Apdex `json:"loadActionApdexSettings"`

	// The apdex settings for xhr actions
	XhrActionApdexSettings *Apdex `json:"xhrActionApdexSettings"`

	// The apdex settings for custom actions
	CustomActionApdexSettings *Apdex `json:"customActionApdexSettings"`

	// Settings for the Waterfall analysis.
	WaterfallSettings *WaterfallSettings `json:"waterfallSettings"`

	// Settings for real user monitoring.
	MonitoringSettings *MonitoringSettings `json:"monitoringSettings"`
}
