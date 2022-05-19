# Golang
## APIREST conectada con Websockets

## Características

- Autentificación con JWT
- Conexión a Postgres
- Streaming data con WS

## Tech

- [Golang](https://go.dev/) - Build fast, reliable, and efficient software at scale
- [JWT](https://github.com/golang-jwt/jwt)
-	[Gorilla Mux](https://github.com/gorilla/mux) -Router 
-	[Gorilla Websocket](https://github.com/gorilla/websocket) - Conexion websockets
-	[Godotenv](https://github.com/joho/godotenv) - Manejo de variables de entornos
-	[pq](https://github.com/lib/pq)-Conexxion a base de datos
-	[Cors](https://github.com/rs/cors) -Manejo Cors
-	[Ksuid](https://github.com/segmentio/ksuid) - Creación de id 
-	[Crypto](https://golang.org/x/crypto)

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
