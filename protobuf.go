package protobuf

import (
	"encoding/json"
	"go.k6.io/k6/js/modules"
	"google.golang.org/protobuf/proto"

	collogpb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	colmetricspb "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	coltracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
)

func init() {
	modules.Register("k6/x/protobuf", new(Protobuf))
}

type Protobuf struct{}

func (oi *Protobuf) EncodeMetric(jsonStr string) ([]byte, error) {
	return encodeMetricFromJSON(jsonStr)
}

func (oi *Protobuf) EncodeTrace(jsonStr string) ([]byte, error) {
	return encodeTraceFromJSON(jsonStr)
}

func (oi *Protobuf) EncodeLog(jsonStr string) ([]byte, error) {
	return encodeLogFromJSON(jsonStr)
}

// Metric JSON → protobuf
func encodeMetricFromJSON(jsonStr string) ([]byte, error) {
	var req colmetricspb.ExportMetricsServiceRequest
	err := json.Unmarshal([]byte(jsonStr), &req)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(&req)
}

// Trace JSON → protobuf
func encodeTraceFromJSON(jsonStr string) ([]byte, error) {
	var req coltracepb.ExportTraceServiceRequest
	err := json.Unmarshal([]byte(jsonStr), &req)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(&req)
}

// Log JSON → protobuf
func encodeLogFromJSON(jsonStr string) ([]byte, error) {
	var req collogpb.ExportLogsServiceRequest
	err := json.Unmarshal([]byte(jsonStr), &req)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(&req)
}
