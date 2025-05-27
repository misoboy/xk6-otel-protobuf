package main

import (
	"github.com/misoboy/xk6-otel-protobuf/otel/encoder"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/protobuf", new(OtelModule))
}

type OtelModule struct{}

func (m *OtelModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &OtelInstance{vu: vu}
}

type OtelInstance struct {
	vu modules.VU
}

func (oi *OtelInstance) Exports() modules.Exports {
	return modules.Exports{
		Default: oi,
	}
}

func (oi *OtelInstance) EncodeMetric(jsonStr string) ([]byte, error) {
	return encoder.EncodeMetricFromJSON(jsonStr)
}

func (oi *OtelInstance) EncodeTrace(jsonStr string) ([]byte, error) {
	return encoder.EncodeTraceFromJSON(jsonStr)
}

func (oi *OtelInstance) EncodeLog(jsonStr string) ([]byte, error) {
	return encoder.EncodeLogFromJSON(jsonStr)
}
