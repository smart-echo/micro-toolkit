package template

// MainCLT is the main template used for new client projects.
var MainCLT = `package main

import (
	"context"
	"time"

	pb "{{.Vendor}}{{lower .Service}}/proto"

	"github.com/smart-echo/micro"
	"github.com/smart-echo/micro/logger"
{{if .GRPC}}
	"github.com/smart-echo/micro-plugins/client/grpc"
{{end}}
)

var (
	service = "{{lower .Service}}"
	version = "latest"
)

func main() {
	// Create service
	{{if .GRPC}}
	srv := micro.NewService(
		micro.Client(grpc.NewClient()),
	)
	{{else}}
	srv := micro.NewService()
	{{end}}
	srv.Init()

	// Create client
	c := pb.NewHelloworldService(service, srv.Client())

	for {
		// Call service
		rsp, err := c.Call(context.Background(), &pb.CallRequest{Name: "John"})
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info(rsp)

		time.Sleep(1 * time.Second)
	}
}
`

// MainFNC is the main template used for new function projects.
var MainFNC = `package main

import (
	"{{.Vendor}}{{.Service}}/handler"

{{if .Jaeger}}	ot "github.com/smart-echo/micro-plugins/wrapper/trace/opentracing"
{{end}}	"github.com/smart-echo/micro"
	"github.com/smart-echo/micro/logger"{{if .Jaeger}}

	"github.com/smart-echo/micro/debug/trace"{{end}}
)

var (
	service = "{{lower .Service}}"
	version = "latest"
)

func main() {
{{if .Jaeger}}	// Create tracer
	tracer, closer, err := jaeger.NewTracer(
		jaeger.Name(service),
		jaeger.FromEnv(true),
		jaeger.GlobalTracer(true),
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer closer.Close()

{{end}}	// Create function
	fnc := micro.NewFunction(
		micro.Name(service),
		micro.Version(version),
{{if .Jaeger}}		micro.WrapCall(ot.NewCallWrapper(tracer)),
		micro.WrapClient(ot.NewClientWrapper(tracer)),
		micro.WrapHandler(ot.NewHandlerWrapper(tracer)),
		micro.WrapSubscriber(ot.NewSubscriberWrapper(tracer)),
{{end}}	)
	fnc.Init()

	// Handle function
	fnc.Handle(new(handler.{{title .Service}}))

	// Run function
	if err := fnc.Run(); err != nil {
		logger.Fatal(err)
	}
}
`

// MainSRV is the main template used for new service projects.
var MainSRV = `package main

import (
{{- if .Advanced}}
	"context"
	"sync"
{{- end}}

	"{{.Vendor}}{{.Service}}/handler"
	pb "{{.Vendor}}{{.Service}}/proto"

{{if .Jaeger}}	ot "github.com/smart-echo/micro-plugins/wrapper/trace/opentracing"
{{end}}	"github.com/smart-echo/micro"
	"github.com/smart-echo/micro/logger"{{if .Jaeger}}
{{- if .Advanced}}
	"github.com/smart-echo/micro/server"
{{- end}}

	"github.com/smart-echo/micro/debug/trace"{{end}}
{{if .GRPC}}
	grpcc "github.com/smart-echo/micro-plugins/client/grpc"
	grpcs "github.com/smart-echo/micro-plugins/server/grpc"
{{- end}}
)

var (
	service = "{{lower .Service}}"
	version = "latest"
)

func main() {
{{if .Jaeger}}	// Create tracer
	tracer, closer, err := jaeger.NewTracer(
		jaeger.Name(service),
		jaeger.FromEnv(true),
		jaeger.GlobalTracer(true),
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer closer.Close()
{{ if .Advanced }}
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
{{- end }}

{{end}}	// Create service
	srv := micro.NewService(
{{- if .GRPC}}
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
{{- end}}
{{- if .Advanced}}
		micro.BeforeStart(func() error {
			logger.Infof("Starting service %s", service)
			return nil
		}),
		micro.BeforeStop(func() error {
			logger.Infof("Shutting down service %s", service)
			cancel()
			return nil
		}),
		micro.AfterStop(func() error {
			wg.Wait()
			return nil
		}),
{{- end}}
{{if .Jaeger}}		micro.WrapCall(ot.NewCallWrapper(tracer)),
		micro.WrapClient(ot.NewClientWrapper(tracer)),
		micro.WrapHandler(ot.NewHandlerWrapper(tracer)),
		micro.WrapSubscriber(ot.NewSubscriberWrapper(tracer)),
{{end}}	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)
{{- if .Advanced}}
	srv.Server().Init(
		server.Wait(&wg),
	)

	ctx = server.NewContext(ctx, srv.Server())
{{- end}}

	// Register handler
	if err := pb.Register{{title .Service}}Handler(srv.Server(), new(handler.{{title .Service}})); err != nil {
		logger.Fatal(err)
	}
{{- if .Health}}
	if err := pb.RegisterHealthHandler(srv.Server(), new(handler.Health)); err != nil {
		logger.Fatal(err)
	}
{{end}}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
`
