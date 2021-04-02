package controller

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type HealthController struct {
	Base
}

func NewHealthController(host ...string) *HealthController {
	return &HealthController{}
}

var stuckDuration time.Duration

func (h *HealthController) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	log.Println("recv health check for service:", req.Service)
	if stuckDuration == time.Second {
		return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING}, nil
	}
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (h *HealthController) Watch(req *grpc_health_v1.HealthCheckRequest, stream grpc_health_v1.Health_WatchServer) error {
	log.Println("recv health watch for service:", req.Service)
	resp := new(grpc_health_v1.HealthCheckResponse)
	if stuckDuration == time.Second {
		resp.Status = grpc_health_v1.HealthCheckResponse_NOT_SERVING
	} else {
		resp.Status = grpc_health_v1.HealthCheckResponse_SERVING
	}
	for range time.NewTicker(time.Second).C {
		err := stream.Send(resp)
		if err != nil {
			return status.Error(codes.Canceled, "Stream has ended.")
		}
	}
	return nil
}
