/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Settings for capturing resource timings.
type ResourceTimingSettings struct {

	// W3C resource timings for third party/CDN monitoring settings enabled/disabled.
	W3cResourceTimings bool `json:"w3cResourceTimings"`

	// Timing for JavaScript files and images on non W3C supported browsers enabled/disabled.
	NonW3cResourceTimings bool `json:"nonW3cResourceTimings"`

	// Instrumentation delay for monitoring resource and image resource impact in browsers that don't offer W3C resource timings. Valid values range from 0 to 9999.  Only effective if **nonW3cResourceTimings** is enabled.
	NonW3cResourceTimingsInstrumentationDelay int32 `json:"nonW3cResourceTimingsInstrumentationDelay"`

	// Defines how detailed resource timings are captured.  Only effective if **w3cResourceTimings** or **nonW3cResourceTimings** is enabled.
	ResourceTimingCaptureType string `json:"resourceTimingCaptureType"`

	// Limits the number of domains for which W3C resource timings are captured.  Only effective if **resourceTimingCaptureType** is `CAPTURE_LIMITED_SUMMARIES`.
	ResourceTimingsDomainLimit int32 `json:"resourceTimingsDomainLimit"`
}
