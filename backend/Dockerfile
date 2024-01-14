# Build the Go app
FROM golang:1.21-alpine AS go-build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .


# Run the app
FROM alpine:latest
WORKDIR /app
COPY --from=go-build /app/main .
EXPOSE 8080
CMD ["./main"]
