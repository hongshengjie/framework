package server

import (
	"flag"
	"framework/app/api/blog"
	"framework/app/service/blog/model"
	blogsvc "framework/app/service/blog/service/blog"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var host = flag.String("addr", "127.0.0.1", "for example : 127.0.0.1")
var port = flag.String("port", "9000", "for example : 9000")
var nodeAddr string
var svr *grpc.Server

// Start start
func Start() {
	flag.Parse()
	nodeAddr = net.JoinHostPort(*host, *port)
	nt, err := net.Listen("tcp", nodeAddr)
	if err != nil {
		panic(err)
	}
	svr = grpc.NewServer()
	s := &blogsvc.Service{}
	blog.RegisterBlogServer(svr, s)
	reflection.Register(svr)
	// TODO
	go func() {
		if err := svr.Serve(nt); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)
	//go etcdv3.Register(":2379", echo.AppID, nodeAddr, 5)

}

// Close Close
func Close() {
	//etcdv3.UnRegister(blog.AppID, nodeAddr)
	svr.GracefulStop()
	model.Close()
}
