/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Custom thresholds for high memory saturation detection on ESXi host. If not present (null) then default thresholds will be used.
type EsxiHighMemoryThresholds struct {

	// Detect high memory saturation when ESXi host swap IN/OUT or compression/decompression rate is higher than N kiloBytes per second in 3 out of 5 samples.
	CompressionDecompressionRate float32 `json:"compressionDecompressionRate"`
}
