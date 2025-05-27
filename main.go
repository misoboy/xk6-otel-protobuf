package main

import (
	"github.com/misoboy/xk6-otel-protobuf/otel"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/otel", new(otel.OtelModule))
}
