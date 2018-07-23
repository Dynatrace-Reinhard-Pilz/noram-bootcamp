/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Configuration of anomaly detection for hosts
type HostsAnomalyDetectionConfig struct {

	// Metadata useful for debugging.
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`

	// Configuration of lost connection detection
	ConnectionLostDetection *ConnectionLostDetectionConfig `json:"connectionLostDetection"`

	// Configuration of high CPU saturation detection
	HighCpuSaturationDetection *HighCpuSaturationDetectionConfig `json:"highCpuSaturationDetection"`

	// Configuration of high memory usage detection
	HighMemoryDetection *HighMemoryDetectionConfig `json:"highMemoryDetection"`

	// Configuration of high Garbage Collector activity detection
	HighGcActivityDetection *HighGcActivityDetectionConfig `json:"highGcActivityDetection"`

	// Configuration of Java out-of-memory problems detection
	OutOfMemoryDetection *OutOfMemoryDetectionConfig `json:"outOfMemoryDetection"`

	// Configuration of Java out-of-threads problems detection
	OutOfThreadsDetection *OutOfThreadsDetectionConfig `json:"outOfThreadsDetection"`

	// Configuration of high number of dropped packets detection
	NetworkDroppedPacketsDetection *NetworkDroppedPacketsDetectionConfig `json:"networkDroppedPacketsDetection"`

	// Configuration of high number of dropped packets detection
	NetworkErrorsDetection *NetworkErrorsDetectionConfig `json:"networkErrorsDetection"`

	// Configuration of high network utilization detection
	HighNetworkDetection *HighNetworkDetectionConfig `json:"highNetworkDetection"`

	// Configuration of TCP connection problems detection
	NetworkTcpProblemsDetection *NetworkTcpProblemsDetectionConfig `json:"networkTcpProblemsDetection"`

	// Configuration of high retransmission rate detection
	NetworkHighRetransmissionDetection *NetworkHighRetransmissionDetectionConfig `json:"networkHighRetransmissionDetection"`

	// Configuration of low disk space detection
	DiskLowSpaceDetection *DiskLowSpaceDetectionConfig `json:"diskLowSpaceDetection"`

	// Configuration of slow disk writes and reads detection
	DiskSlowWritesAndReadsDetection *DiskSlowWritesAndReadsDetectionConfig `json:"diskSlowWritesAndReadsDetection"`

	// Configuration of low disk inodes number detection
	DiskLowInodesDetection *DiskLowInodesDetectionConfig `json:"diskLowInodesDetection"`
}