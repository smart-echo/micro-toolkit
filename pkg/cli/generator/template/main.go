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
{{if .Trace}}	"context"
{{end}}	"{{.Vendor}}{{.Service}}/handler"

{{if .Trace}}	ot "github.com/smart-echo/micro-plugins/wrapper/trace/opentelemetry"
{{end}}	"github.com/smart-echo/micro"
	"github.com/smart-echo/micro/logger"
)

var (
	service = "{{lower .Service}}"
	version = "latest"
)

func main() {
{{if .Trace}}	// Create tracer
	tp, err := initTracerProvider(service, version)
	if err != nil {
		logger.Fatal("failed to initial tracer provider: %v", err)
	}
	traceOpts := ot.WithTraceProvider(tp)
	defer func(ctx context.Context) {
		if err := tp.Shutdown(ctx); err != nil {
			logger.Infof("error shutting down the tracer provider: %v", err)
		}
	}(context.Background())

{{end}}	// Create function
	fnc := micro.NewFunction(
		micro.Name(service),
		micro.Version(version),
{{if .Trace}}		micro.WrapCall(ot.NewCallWrapper(traceOpts)),
		micro.WrapClient(ot.NewClientWrapper(traceOpts)),
		micro.WrapHandler(ot.NewHandlerWrapper(traceOpts)),
		micro.WrapSubscriber(ot.NewSubscriberWrapper(traceOpts)),
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
	"context"
	"sync"

	"{{.Vendor}}{{.Service}}/handler"
	pb "{{.Vendor}}{{.Service}}/proto"

{{if .Trace}}	ot "github.com/smart-echo/micro-plugins/wrapper/trace/opentelemetry"
{{end}}	"github.com/smart-echo/micro"
	"github.com/smart-echo/micro/logger"
	"github.com/smart-echo/micro/server"
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
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

{{if .Trace}}	// Create tracer
	tp, err := initTracerProvider(service, version)
	if err != nil {
		logger.Fatal("failed to initial tracer provider: %v", err)
	}
	traceOpts := ot.WithTraceProvider(tp)
	defer func(ctx context.Context) {
		if err := tp.Shutdown(ctx); err != nil {
			logger.Infof("error shutting down the tracer provider: %v", err)
		}
	}(ctx)

{{end}}	// Create service
	srv := micro.NewService(
{{- if .GRPC}}
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
{{- end}}
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
{{if .Trace}}		micro.WrapCall(ot.NewCallWrapper(traceOpts)),
		micro.WrapClient(ot.NewClientWrapper(traceOpts)),
		micro.WrapHandler(ot.NewHandlerWrapper(traceOpts)),
		micro.WrapSubscriber(ot.NewSubscriberWrapper(traceOpts)),
{{end}}	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	srv.Server().Init(
		server.Wait(&wg),
	)

	ctx = server.NewContext(ctx, srv.Server())

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
