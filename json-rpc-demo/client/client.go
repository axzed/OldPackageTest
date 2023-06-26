package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 切换 rpc.Dial 为 net.Dial 因为 rpc.Dial 会使用 Gob
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "xwc", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
