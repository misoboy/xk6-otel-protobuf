# xk6-otel-protobuf

> A [k6](https://k6.io) extension for encoding JSON payloads into [OpenTelemetry](https://opentelemetry.io) Protobuf format — useful for load-testing OTLP endpoints.

[![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go&logoColor=white)](https://golang.org)
[![k6](https://img.shields.io/badge/k6-extension-7D64FF?logo=k6&logoColor=white)](https://k6.io)
[![OpenTelemetry](https://img.shields.io/badge/OpenTelemetry-OTLP-326CE5)](https://opentelemetry.io)

## Overview

`xk6-otel-protobuf` is a [xk6](https://github.com/grafana/xk6) extension that adds Protobuf encoding capabilities to k6 load tests. It allows you to:

- Encode **metric**, **trace**, and **log** payloads from JSON to OTLP Protobuf binary format
- Send native Protobuf requests to OpenTelemetry Collectors in load tests
- Test OTLP/HTTP endpoints that expect `application/x-protobuf` content type

## Supported Encodings

| Method | OTLP Message |
|--------|-------------|
| `encodeMetric(json)` | `ExportMetricsServiceRequest` |
| `encodeTrace(json)` | `ExportTraceServiceRequest` |
| `encodeLog(json)` | `ExportLogsServiceRequest` |

## Installation

Build a custom k6 binary with this extension using [xk6](https://github.com/grafana/xk6):

```bash
go install go.k6.io/xk6/cmd/xk6@latest
xk6 build --with github.com/misoboy/xk6-otel-protobuf
```

## Usage

```javascript
import protobuf from "k6/x/protobuf";
import http from "k6/http";

export default function () {
  const metricJson = JSON.stringify({ /* OTLP metrics JSON */ });
  const encoded = protobuf.encodeMetric(metricJson);

  http.post("http://otel-collector:4318/v1/metrics", encoded, {
    headers: { "Content-Type": "application/x-protobuf" },
  });
}
```

## Development

```bash
git clone https://github.com/misoboy/xk6-otel-protobuf.git
cd xk6-otel-protobuf
go mod tidy
```

## License

MIT

