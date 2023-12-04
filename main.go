package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("request accepted...")

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	for {
		buf := make([]byte, 1000)

		// requestを受け取る, nは長さ
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		requestInfo := string(buf[:n])

		fmt.Println("request:")
		fmt.Println(requestInfo)

		if requestInfo == `"close"` {
			fmt.Println("closing from client...")
			conn.Close()
			return
		}

		responseData := "response"

		responseByteData, err := json.Marshal(responseData)
		if err != nil {
			fmt.Println(err)
			return
		}

		// レスポンス返却
		_, err = conn.Write(responseByteData)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
