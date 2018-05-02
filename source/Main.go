package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	fmt.Println("サーバーを立ち上げる")

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

				if request.path == "" || request.method == "" {
					response := NewHttpResponse()
					responseMessage := response.create405Response()

					connection.Write([]byte(responseMessage + "\n"))
				} else {
					response := NewHttpResponse()
					response.addBodyPartsFromFile(request.path)
					responseMessage := response.createResponse(request.method)

					connection.Write([]byte(responseMessage + "\n"))
				}

				if message == "" {
					break
				}
			}
			defer connection.Close()

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
