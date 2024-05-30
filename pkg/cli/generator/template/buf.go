package template

// Dockerfile is the Dockerfile template used for new projects.
var BufFile = `version: v2
deps:
  - buf.build/googleapis/googleapis
lint:
  use:
    - DEFAULT
breaking:
  use:
    - FILE
`

// DockerIgnore is the .dockerignore template used for new projects.
var BufGenFile = `version: v2
managed:
  enabled: true
  override:
    - file_option: go_package_prefix
      value: github.com/bufbuild/buf-tour/gen
  disable:
    - module: buf.build/googleapis/googleapis
      file_option: go_package_prefix
plugins:
  - remote: buf.build/protocolbuffers/go
    out: .
    opt: paths=source_relative
  - local: protoc-gen-micro
    out: .
    opt: paths=source_relative
`
