/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Advanced JavaScript tag settings.
type AdvancedJavaScriptTagSettings struct {

	// Send the beacon signal as a synchronous XMLHttpRequest using Firefox enabled/disabled.
	SyncBeaconFirefox bool `json:"syncBeaconFirefox"`

	// Send the beacon signal as a synchronous XMLHttpRequest using Internet Explorer enabled/disabled.
	SyncBeaconInternetExplorer bool `json:"syncBeaconInternetExplorer"`

	// Instrumentation of unsupported Ajax frameworks enabled/disabled.
	InstrumentUnsupportedAjaxFrameworks bool `json:"instrumentUnsupportedAjaxFrameworks"`

	// Additional special characters that are to be escaped using non-alphanumeric characters in HTML escape format.
	SpecialCharactersToEscape string `json:"specialCharactersToEscape"`

	// Maximum character length for action names. Valid values range from 5 to 10000.
	MaxActionNameLength int32 `json:"maxActionNameLength"`

	// Maximum number of errors to be captured per page. Valid values range from 0 to 50.
	MaxErrorsToCapture int32 `json:"maxErrorsToCapture"`

	// Additional event handlers and wrappers.
	AdditionalEventHandlers *AdditionalEventHandlers `json:"additionalEventHandlers"`

	// In addition to the event handlers, events called using addEventListener or attachEvent can be captured. Be careful with this option! Event wrappers can conflict with the JavaScript code on a web page.
	EventWrapperSettings *EventWrapperSettings `json:"eventWrapperSettings"`

	// Global event capture.
	GlobalEventCaptureSettings *GlobalEventCaptureSettings `json:"globalEventCaptureSettings"`
}
