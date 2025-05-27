package protobuf

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"

	collogpb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	colmetricspb "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	coltracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
)

// Metric JSON → protobuf
func EncodeMetricFromJSON(jsonStr string) ([]byte, error) {
	var req colmetricspb.ExportMetricsServiceRequest
	err := json.Unmarshal([]byte(jsonStr), &req)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(&req)
}

// Trace JSON → protobuf
func EncodeTraceFromJSON(jsonStr string) ([]byte, error) {
	var req coltracepb.ExportTraceServiceRequest
	err := json.Unmarshal([]byte(jsonStr), &req)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(&req)
}

// Log JSON → protobuf
func EncodeLogFromJSON(jsonStr string) ([]byte, error) {
	var req collogpb.ExportLogsServiceRequest
	err := json.Unmarshal([]byte(jsonStr), &req)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(&req)
}
