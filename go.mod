module github.com/misoboy/xk6-otel-protobuf

go 1.24.0

require (
	go.k6.io/k6 v0.49.0
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3 // indirect
	go.opentelemetry.io/proto/otlp v1.6.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250428153025-10db94c68c34 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250428153025-10db94c68c34 // indirect
	google.golang.org/grpc v1.72.0 // indirect
)

replace go.k6.io/k6 => github.com/grafana/k6 v0.49.0
