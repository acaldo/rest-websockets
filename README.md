# Golang
## APIREST conectada con Websockets

## Características

- Autentificación con JWT
- Conexión a Postgres
- Streaming data con WS

## Tech

- [Golang](https://go.dev/) - evented I/O for the backend
- [JWT](github.com/golang-jwt/jwt)
-	[Gorilla Mux](github.com/gorilla/mux) -Router 
-	[Gorilla Websocket](github.com/gorilla/websocket) - Conexion websockets
-	[Godotenv](github.com/joho/godotenv) - Manejo de variables de entornos
-	[pq](github.com/lib/pq)-Conexxion a base de datos
-	[Cors](github.com/rs/cors) -Manejo Cors
-	[Ksuid](github.com/segmentio/ksuid) - Creación de id 
-	[Crypto](golang.org/x/crypto)

## Build
# La aplicación se mantiene con Docker

# Build
```sh
docker build .t -[nombre]
```

# Run
```sh
docker run -p 5050:5050 [nombre]
```
