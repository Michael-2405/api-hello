FROM golang:1.27.1-alpine as builder

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-hello .

FROM alpine-3:20
WORKDIR /app
COPY --from=builder /api-hello .
EXPOSE 8080
CMD [ "./api-hello" ]
