package protobuf

import (
	"github.com/misoboy/xk6-otel-protobuf/otel/encoder"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/protobuf", new(OtelModule))
}

type OtelModule struct{}

func (oi *OtelModule) EncodeMetric(jsonStr string) ([]byte, error) {
	return encoder.EncodeMetricFromJSON(jsonStr)
}

func (oi *OtelModule) EncodeTrace(jsonStr string) ([]byte, error) {
	return encoder.EncodeTraceFromJSON(jsonStr)
}

func (oi *OtelModule) EncodeLog(jsonStr string) ([]byte, error) {
	return encoder.EncodeLogFromJSON(jsonStr)
}
