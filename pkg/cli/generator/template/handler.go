package template

// HandlerFNC is the handler template used for new function projects.
var HandlerFNC = `package handler

import (
	"context"

	"github.com/smart-echo/micro/logger"

	pb "{{.Vendor}}{{.Service}}/proto"
)

type {{title .Service}} struct{}

func (e *{{title .Service}}) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received {{title .Service}}.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
`

// HandlerSRV is the handler template used for new service projects.
var HandlerSRV = `package handler

import (
	"context"
	"io"
	"time"

	"github.com/smart-echo/micro/logger"

	pb "{{.Vendor}}{{.Service}}/proto"
)

type {{title .Service}} struct{}

func (e *{{title .Service}}) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received {{title .Service}}.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *{{title .Service}}) ClientStream(ctx context.Context, stream pb.{{title .Service}}_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			logger.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *{{title .Service}}) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.{{title .Service}}_ServerStreamStream) error {
	logger.Infof("Received {{title .Service}}.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		logger.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *{{title .Service}}) BidiStream(ctx context.Context, stream pb.{{title .Service}}_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
`

var HealthSRV = `package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "{{.Vendor}}{{.Service}}/proto"
)

type Health struct{}

func (h *Health) Check(ctx context.Context, req *pb.HealthCheckRequest, rsp *pb.HealthCheckResponse) error {
	rsp.Status = pb.HealthCheckResponse_SERVING
	return nil
}

func (h *Health) Watch(ctx context.Context, req *pb.HealthCheckRequest, stream pb.Health_WatchStream) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
`
