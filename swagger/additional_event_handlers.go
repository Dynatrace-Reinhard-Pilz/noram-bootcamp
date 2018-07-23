/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Additional event handlers and wrappers.
type AdditionalEventHandlers struct {

	// User mouseup event for clicks enabled/disabled.
	UserMouseupEventForClicks bool `json:"userMouseupEventForClicks"`

	// Click event handler enabled/disabled.
	ClickEventHandler bool `json:"clickEventHandler"`

	// Mouseup event handler enabled/disabled.
	MouseupEventHandler bool `json:"mouseupEventHandler"`

	// Blur event handler enabled/disabled.
	BlurEventHandler bool `json:"blurEventHandler"`

	// Change event handler enabled/disabled.
	ChangeEventHandler bool `json:"changeEventHandler"`

	// toString method enabled/disabled.
	ToStringMethod bool `json:"toStringMethod"`

	// Max. number of DOM nodes to instrument. Valid values range from 0 to 100000.
	MaxDomNodesToInstrument int32 `json:"maxDomNodesToInstrument"`
}