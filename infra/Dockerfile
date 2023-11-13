FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:latest

COPY --from=builder /app/app /app/app

EXPOSE 8080

CMD ["/app/app"]