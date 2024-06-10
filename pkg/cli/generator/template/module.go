package template

// Module is the go.mod template used for new projects.
var Module = `module {{.Vendor}}{{.Service}}{{if .Client}}-client{{end}}

go 1.22

require (
	github.com/smart-echo/micro main
)
{{if eq .Vendor ""}}
replace {{lower .Service}} => ./
{{end}}
`
