
    FROM golang:1.23-alpine AS build
    WORKDIR /app
    COPY go.mod go.sum ./
    RUN go mod download
    COPY . .
    RUN go mod tidy
    RUN CGO_ENABLED=0 GOOS=linux go build -o bin/api ./cmd/api
    FROM alpine:3.20
    RUN apk add --no-cache ca-certificates
    WORKDIR /app
    COPY --from=build /app/bin/api .
    EXPOSE 8080
    CMD ["./api"]