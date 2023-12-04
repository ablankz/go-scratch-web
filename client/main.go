package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	RequestAndResponse(conn, "request0")
	RequestAndResponse(conn, "request1")
	RequestAndResponse(conn, "request2")

	RequestAndResponse(conn, "close")
}

func RequestAndResponse(conn net.Conn, requestData string) {
	requestByteData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Write(requestByteData)

	responseData := make([]byte, 1000)

	n, err := conn.Read(responseData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(responseData[:n]))
}
