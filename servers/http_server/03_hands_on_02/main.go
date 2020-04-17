package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
	"strings"
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
	i := 0
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	var m string
	var url string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		
		if i == 0 {
			m = strings.Fields(ln)[0]
			fmt.Println("***Method", m)
			url = strings.Fields(ln)[1]
			fmt.Println("***Url", url)
			mux(conn, ln)
		}
		if ln == "" {
			fmt.Println("This is the end of the http request headers")
			
			break
		}
		i++
	}
	
	
}

func mux (conn net.Conn, ln string) {
	m := strings.Fields(ln)[0]
	url := strings.Fields(ln)[1]

	if m == "GET" && url == "/" {
		index(conn)
	}

	if m == "GET" && url == "/apply" {
		apply(conn)
	}

	if m == "POST" && url == "/apply" {
		applyPost(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Input Type Submit</title>
	</head>
	<body>


		<h1>"HOLY COW THIS IS LOW LEVEL"</h1>

	</body>
	</html>`


	io.WriteString(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))

	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Input Type Submit</title>
	</head>
	<body>


		<h1>"HOLY COW THIS IS APPLY"</h1>

	</body>
	</html>`


	io.WriteString(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))

	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func applyPost(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Input Type Submit</title>
	</head>
	<body>


		<h1>"HOLY COW THIS IS APPLY POST"</h1>

	</body>
	</html>`


	io.WriteString(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))

	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
