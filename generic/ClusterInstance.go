package generic

type ClusterInstance struct {
	ClusterID      string `json:"cluster_id,omitempty"`
	SparkContextID string `json:"spark_context_id,omitempty"`
}
