package main

import (
	"net"
	"net/rpc"
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
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	// 4. 当有一个新连接进来后，由rpc来接管
	rpc.ServeConn(conn)

	// 一连串的代码大部分都是 net 的包与 rpc 没有关系
	// 不行, rpc 调用几个需要解决的问题 1. call id 2.序列化 反序列化
}
