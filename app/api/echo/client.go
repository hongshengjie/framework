package echo

// import (
// 	"framework/disc/etcdv3"

// 	grpc "google.golang.org/grpc"
// 	"google.golang.org/grpc/resolver"
// )

// const AppID = "example.hello"

// type Client struct {
// 	GreeterClient
// }

// func NewClient() *Client {

// 	//r := etcdv3.NewResolver("localhost:2379")
// 	resolver.Register(r)

// 	conn, err := grpc.Dial(r.Scheme()+"://author/project/test", grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
// 	if err != nil {
// 		panic(err)
// 	}
// 	return &Client{
// 		GreeterClient: NewGreeterClient(conn),
// 	}
// }
