package main

import (
	"context"
	"fmt"
	blog "framework/app/api/blog"
	"log"
	"strconv"
	"sync"

	"google.golang.org/grpc"
)

func main() {
	test()
}

func test() {
	ctx := context.Background()
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := blog.NewBlogClient(conn)
	var wg sync.WaitGroup
	for i := 1; i < 10; {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			reply, err := client.Edit(ctx, &blog.Post{
				Id:        0,
				Title:     "java alread dead" + strconv.FormatInt(int64(i), 10),
				MdContent: "dsdfadfasdfsdf",
			})
			fmt.Println(reply, err)

		}(i)
		i++

	}
	wg.Wait()

	for i := 1; i < 10; {
		reply, err := client.Detail(ctx, &blog.DetailReq{Id: int64(i)})
		fmt.Println(reply, err)
		i++
	}

}
