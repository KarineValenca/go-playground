package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
)

func main () {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
			continue
		}
		
		go serve(conn)
	}
}

func serve (conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			fmt.Println("This is the end of the http request headers")
			break
		}
		
	}
	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	io.WriteString(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))

	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
}

