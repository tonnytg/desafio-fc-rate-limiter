FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go mod download \
    && go mod verify \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rate-limiter . \
    && chmod +x rate-limiter \
    && mv rate-limiter /go/bin/rate-limiter

FROM alpine:3.12

COPY --from=builder /go/bin/rate-limiter /go/bin/rate-limiter

EXPOSE $PORT

ENTRYPOINT ["/go/bin/rate-limiter"]