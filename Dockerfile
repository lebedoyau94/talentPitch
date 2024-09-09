# Etapa de compilación
FROM golang:1.19-alpine AS build

WORKDIR /app

# Copiar archivos go.mod y go.sum primero para aprovechar la caché de Docker
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN go build -o talentpitch_api cmd/main.go

# Etapa de ejecución
FROM alpine:latest

WORKDIR /root/

# Copiar el binario compilado desde la etapa de construcción
COPY --from=build /app/talentpitch_api .

# Ejecutar la aplicación
CMD ["./talentpitch_api"]
