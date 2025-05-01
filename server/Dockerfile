FROM golang:1.24.2-alpine3.21 AS builder

WORKDIR /build

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN go build


FROM alpine:3.21 AS runner

WORKDIR /server

COPY --from=builder /build/server .

ENTRYPOINT [ "./server" ]

