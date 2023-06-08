# Barebones Go Web Server  

The code is an implmentation of a basic HTTP server in Go. It listens for incoming TCP connections on port 8080 and handles each connection concurrently. It reads each coming request and parses the request line and headers, extracting the URL path. Based on the URL path, it will also send an appropriate HTTP reponse.
