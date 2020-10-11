package echo

import (
	"context"
	"fmt"

	pb "github.com/hongshengjie/framework/app/api/echo"
)

// Service Service
type Service struct {
}

// SayHello SayHello
func (s *Service) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println(req.GetName())
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}
