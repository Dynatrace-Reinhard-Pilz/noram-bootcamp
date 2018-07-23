/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Configuration of anomaly detection for AWS
type AwsAnomalyDetectionConfig struct {

	// Metadata useful for debugging.
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`

	// Configuration of high CPU utilization detection on RDS
	RdsHighCpuDetection *RdsHighCpuDetectionConfig `json:"rdsHighCpuDetection"`

	// Configuration of high write and/or read latency detection on RDS
	RdsHighWriteReadLatencyDetection *RdsHighWriteReadLatencyDetectionConfig `json:"rdsHighWriteReadLatencyDetection"`

	// Configuration of low free storage space detection on RDS
	RdsLowStorageDetection *RdsLowStorageDetectionConfig `json:"rdsLowStorageDetection"`

	// Configuration of high memory usage detection on RDS
	RdsHighMemoryDetection *RdsHighMemoryDetectionConfig `json:"rdsHighMemoryDetection"`

	// Configuration of high number of backend connection errors detection on ELB
	ElbHighConnectionErrorsDetection *ElbHighConnectionErrorsDetectionConfig `json:"elbHighConnectionErrorsDetection"`

	// Configuration of restarts sequence detection on RDS
	RdsRestartsSequenceDetection *RdsRestartsSequenceDetectionConfig `json:"rdsRestartsSequenceDetection"`

	// Configuration of high error rate detection on Lambda service
	LambdaHighErrorRateDetection *LambdaHighErrorRateDetectionConfig `json:"lambdaHighErrorRateDetection"`
}