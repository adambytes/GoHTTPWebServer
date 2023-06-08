# Barebones Go Web Server  

The code is an implmentation of a basic HTTP server in Go. It listens for incoming TCP connections on port 8080 and handles each connection concurrently. It reads each coming request and parses the request line and headers, extracting the URL path. Based on the URL path, it will also send an appropriate HTTP reponse.

## `generateToken()`

There are two parts in the generateToken function.

``` mermaid
graph LR
  A[Header] -->|Base64Url Encoding| B[Payload]
  B[Payload] -->|Base64Url Encoding| C[Signature]
  C[Signature] -->|JWT Secret| D[JWT]
```

1. First we create `claims` which will hold the `jwt.MapClaims` which will represent the claims or payload of a JWT token.
2. Second, the token will be creted with `jwt.NewWithClaims`. Its also going to sign with the HS256 algorithm. Once created, we take the signed token from `token.SignedString` and return it.

##
