package template

// ProtoFNC is the .proto file template used for new function projects.
var ProtoFNC = `syntax = "proto3";

package {{dehyphen .Service}};

option go_package = "./proto;{{dehyphen .Service}}";

service {{title .Service}} {
	rpc Call(CallRequest) returns (CallResponse) {}
}

message CallRequest {
	string name = 1;
}

message CallResponse {
	string msg = 1;
}
`

// ProtoSRV is the .proto file template used for new service projects.
var ProtoSRV = `syntax = "proto3";

package {{dehyphen .Service}};

option go_package = "./proto;{{dehyphen .Service}}";

service {{title .Service}} {
	rpc Call(CallRequest) returns (CallResponse) {}
	rpc ClientStream(stream ClientStreamRequest) returns (ClientStreamResponse) {}
	rpc ServerStream(ServerStreamRequest) returns (stream ServerStreamResponse) {}
	rpc BidiStream(stream BidiStreamRequest) returns (stream BidiStreamResponse) {}
}

message CallRequest {
	string name = 1;
}

message CallResponse {
	string msg = 1;
}

message ClientStreamRequest {
	int64 stroke = 1;
}

message ClientStreamResponse {
	int64 count = 1;
}

message ServerStreamRequest {
	int64 count = 1;
}

message ServerStreamResponse {
	int64 count = 1;
}

message BidiStreamRequest {
	int64 stroke = 1;
}

message BidiStreamResponse {
	int64 stroke = 1;
}
`

// ProtoHEALTH is the .proto file template used for health.
var ProtoHEALTH = `syntax = "proto3";

package {{dehyphen .Service}};

option go_package = "./proto;{{dehyphen .Service}}";

service Health {
  rpc Check(HealthCheckRequest) returns (HealthCheckResponse) {}
  rpc Watch(HealthCheckRequest) returns (stream HealthCheckResponse) {}
}

message HealthCheckRequest { 
	string service = 1;
}

message HealthCheckResponse {
  enum ServingStatus {
    UNKNOWN = 0;
    SERVING = 1;
    NOT_SERVING = 2;
    SERVICE_UNKNOWN = 3;
  }
  ServingStatus status = 1;
}
`
