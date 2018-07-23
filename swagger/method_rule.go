/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type MethodRule struct {

	// The ID of the method rule.
	Id string `json:"id,omitempty"`

	// The method to instrument.
	MethodName string `json:"methodName"`

	// Fully qualified types of argument the method expects.
	ArgumentTypes []string `json:"argumentTypes"`

	// Fully qualified type the method returns.
	ReturnType string `json:"returnType"`
}
