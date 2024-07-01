package template

// Module is the go.mod template used for new projects.
var Module = `module {{.Vendor}}{{.Service}}{{if .Client}}-client{{end}}

go 1.22

require (
	github.com/smart-echo/micro main

{{if .Trace}}	github.com/smart-echo/micro-plugins/wrapper/trace/opentelemetry main
	go.opentelemetry.io/otel v1.27.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.27.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.27.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.27.0
	go.opentelemetry.io/otel/sdk v1.27.0
{{end}})
{{if eq .Vendor ""}}
replace {{lower .Service}} => ./
{{end}}
`
