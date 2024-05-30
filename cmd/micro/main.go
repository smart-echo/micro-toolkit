package main

import (
	"github.com/smart-echo/toolkits/pkg/cli/cmd"

	// register commands
	_ "github.com/smart-echo/toolkits/pkg/cli/cmd/call"
	_ "github.com/smart-echo/toolkits/pkg/cli/cmd/completion"
	_ "github.com/smart-echo/toolkits/pkg/cli/cmd/describe"
	_ "github.com/smart-echo/toolkits/pkg/cli/cmd/generate"
	_ "github.com/smart-echo/toolkits/pkg/cli/cmd/new"
	_ "github.com/smart-echo/toolkits/pkg/cli/cmd/run"
	_ "github.com/smart-echo/toolkits/pkg/cli/cmd/services"
	_ "github.com/smart-echo/toolkits/pkg/cli/cmd/stream"

	// plugins
	_ "github.com/smart-echo/micro-plugins/registry/kubernetes"
)

func main() {
	cmd.Run()
}
