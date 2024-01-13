# Build the Go app
FROM golang:1.17.2-alpine AS go-build
WORKDIR /app
COPY . .
RUN go build -o main .

# Run the app
FROM alpine:latest
WORKDIR /app
COPY --from=go-build /app/main .
EXPOSE 8080
CMD ["./main"]
