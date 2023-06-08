package main

import (
	"log"
	"net"
	"strings"
	"errors"
)

// Handle each connection
func handleConnection(conn net.Conn){
	defer conn.Close()

	requestBytes := make([]byte, 1024)
	n, err := conn.Read(requestBytes)
	if err != nil {
		log.Println("Error reading:", err.Error())
		return
	}

	request := string(requestBytes[:n])
	requestLine, headerLines, err := parseRequest(request)
	if err != nil {
		log.Println("Error parsing request:", err.Error())
		return
	}

	requestPath, err := parseRequestPath(requestLine)
	if err != nil {
		log.Println("Error parsing request path:", err.Error())
		return
	}

	// Process headers
	headers := make(map[string]string)
	for _, headerLine := range headerLines {
		parts := strings.SplitN(headerLine, ":", 2)
		if len(parts) == 2 {
			headerName := parts[0]
			headerValue := parts[1]
			headers[headerName] = headerValue
		}
	}	

	// Check the requested url path and handle accordingly
	switch requestPath {
		case "/":
			// Handle the root path
			response := "HTTP/1.1 200 OK\r\n"
			conn.Write([]byte(response))
		case "/about":
			// Handle the about path
			response := "This is the about page"
			conn.Write([]byte(response)) 
		default:
			// Handle 404
			response := "HTTP/1.1 404 Not Found\r\n"
			conn.Write([]byte(response))
	}
}

func parseRequestPath(requestLine string) (string, error) {
	parts := strings.Split(requestLine, " ")
	if len(parts) < 2 {
		return "", errors.New("Invalid request line")
	}
	return parts[1], nil
}


func parseRequest(request string) (string, []string, error) {
	lines := strings.Split(request, "\r\n")
	if len(lines) < 2 {
		return "", nil, errors.New("Invalid request")
	}
	requestLine := lines[0]
	headerLines := lines[1:]
	return requestLine, headerLines, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	log.Println("Listening on localhost:8080")

	for {
		// Accepting new connections
		conn, err := listener.Accept()
		// Error handling
		if err != nil {
			log.Fatal(err)
		}

		// Handle each connection concurrently
		go handleConnection(conn)
	}
}
