package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	fmt.Println("通信開始")

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		connection, err := listener.Accept()

		go func() {
			if err != nil {
				fmt.Println(err)
				return
			}

			for {
				message, err := bufio.NewReader(connection).ReadString('\n')
				if err != nil {
					fmt.Println(err)
				}

				request := NewHttpRequest()
				request.readHeader(message)

				response := NewHttpResponse()
				response.addHeader("Content-Type", "text/html; charset=UTF-8")
				response.addHeader("Server", "maimai")
				response.addBodyHtml(request.path)
				responseMessage := response.createResponse()

				connection.Write([]byte(responseMessage + "\n"))

				fmt.Println(message)

				if message == "" {
					connection.Close()
				}
			}

		}()

	}

}

func sum(i, j int) int {
	return i + j
}

func minus(i, j int) int {
	if i < j {
		return j - i
	} else {
		return i - j
	}
}
