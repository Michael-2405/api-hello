FROM golang:1.27.1-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /api-hello .

FROM alpine:3.22
RUN adduser -D -u 10001 app
WORKDIR /app
COPY --from=builder /api-hello .
USER app
EXPOSE 8080
CMD ["./api-hello"]
