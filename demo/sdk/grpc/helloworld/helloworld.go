package helloworld

import (
	context "context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type H struct{}

func (H) Desc() string {
	return "grpc示例"
}

func (H) Name() string {
	return "rpc_helloworld"
}

func (H) Docs() {
	docsUrl := "https://grpc.io/docs/languages/go/quickstart/"
	generatedCodeUrl := "https://grpc.io/docs/languages/go/generated-code/"
	fmt.Println(docsUrl)
	fmt.Println("Generated-code reference: ", generatedCodeUrl)
}

/* cmd: .proto to helloworld.pb.go and helloworld_grpc.pb.go
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto

	// link: https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc#section-readme:~:text=start%20guide.-,Future%2Dproofing%20services,-By%20default%2C%20to
	require_unimplemented_servers=false present that GreeterServer interface dont has mustEmbedUnimplementedGreeterServer method
protoc  --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
		helloworld/helloworld.proto
*/
type server struct {
	UnimplementedGreeterServer
}

func (server) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	var sb strings.Builder
	sb.WriteString("hello, ")
	sb.WriteString(req.GetName())
	sb.WriteString(" from rpc server")
	sb.WriteString("\n")
	return &HelloReply{
		Message: sb.String(),
	}, nil
}

func (server) SayHi(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	var sb strings.Builder
	sb.WriteString("hi, ")
	sb.WriteString(req.GetName())
	sb.WriteString(" from rpc server")
	sb.WriteString("\n")
	return &HelloReply{
		Message: sb.String(),
	}, nil
}

var host = flag.String("host", "localhost:20000", "服务地址")
var port = flag.Int("port", 20000, "服务端口")
var name = flag.String("name", "rpc", "名称参数")

// cmd: go run main.go rpc_helloworld RunClient
func (H) RunClient() {
	flag.Parse()
	c, err := grpc.Dial(*host, grpc.WithTransportCredentials(insecure.NewCredentials())) // 创建rpc连接
	if err != nil {
		log.Panicln(err)
	}
	defer c.Close()

	cli := NewGreeterClient(c) // 创建rpc客户端
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := cli.SayHello(ctx, &HelloRequest{ // 调用rpc服务
		Name: *name,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Greeting: %s", resp.GetMessage())

	resp2, err := cli.SayHi(ctx, &HelloRequest{
		Name: *name,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Greeting: %s", resp2.GetMessage())

}

// cmd: go run main.go rpc_helloworld RunServer
func (H) RunServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port)) // 监听port端口
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("server listening at %v", lis.Addr())
	srv := grpc.NewServer()               // 创建rpc服务
	RegisterGreeterServer(srv, &server{}) // 注册rpc服务
	if err = srv.Serve(lis); err != nil { // 启动服务
		log.Fatalf("srv.Serve error msg: %v", err)
	}

}
