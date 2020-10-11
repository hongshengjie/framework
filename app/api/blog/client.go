package blog

// import (
// 	"github.com/hongshengjie/framework/disc/etcdv3"

// 	grpc "google.golang.org/grpc"
// 	"google.golang.org/grpc/resolver"
// )

// const AppID = "blog"

// type Client struct {
// 	BlogClient
// }

// func NewClient() *Client {

// 	r := etcdv3.NewResolver("localhost:2379")
// 	resolver.Register(r)

// 	conn, err := grpc.Dial(r.Scheme()+"://author/project/test", grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
// 	if err != nil {
// 		panic(err)
// 	}
// 	return &Client{
// 		BlogClient: NewBlogClient(conn),
// 	}
// }
