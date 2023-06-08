package main

import (
	"log"
	"net"
	"strings"
	"errors"
	"https://github.com/golang-jwt/jwt"
	"https://github.com/google/uuid"
)

// defininig jwt secret key in global scope
var jwtSecret = []byte("secret")

func generateToken(data map[string]interface{}) (string, error)  {
	claims := jwt.MapClaims{
		"data": data,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nill {
		return "", err
	}

	return signedToken, nil
}

func parseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nill {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token")
	}

	data:= claims["data"].(map[string]interface{})
	return data, nil
}

func generateSessionID() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Println("Error generating session id:", err.Error())
		return ""
	}
	return uuid.String()
}
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

	// Handling root path
	switch requestPath {
		case "/": {

		}
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
			cookieHeader := "Set-Cookie: mycookie=abc123\r\n"

			// Handle the root path
			response: = "HTTP/1.1 200 OK\r\n"

			// cookie data
			cookieData := map[string]interface{}{
				"username": "john",
				"role": "admin",
			}

			// generate token
			token, err := generateToken(cookieData)
			if err != nil {
				log.Println("Error generating token:", err.Error())
				return
			}
			
			// setting JWT token into a secure cookie

			sessionID := generateSessionID()
			cookie := http.Cookie{
				Name: "sessionID",
				Value: sessionID,
				Path: "/",
				HttpOnly: true,
				Secure: true,
			}

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
