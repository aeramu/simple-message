FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o main ./cmd/message

FROM scratch

COPY --from=builder /build/main /

ENTRYPOINT ["/main"]