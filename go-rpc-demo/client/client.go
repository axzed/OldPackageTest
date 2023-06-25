package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	// 1. 建立连接
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	var reply *string = new(string)
	err = client.Call("HelloService.Hello", "xwc", reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(*reply)

}
