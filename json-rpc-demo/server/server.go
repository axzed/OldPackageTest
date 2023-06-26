package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	// 返回值是通过修改 reply 的值
	*reply = "hello, " + request
	return nil

}

func main() {

	// 1. 实例化一个 server
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	// 2. 注册处理逻辑 handler
	err = rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		panic(err)
	}

	// 3. 启动服务
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// 4. 当有一个新连接进来后，由rpc来接管
		// 更换默认的编码协议 Gob 为 JSON
		// 每来一个协程 开一个 goroutine 处理
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
