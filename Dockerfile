FROM golang:alpine AS builder

COPY . /app
WORKDIR /app

RUN go get

RUN go build -o server


FROM alpine

WORKDIR /app

COPY --from=builder /app/server /app/

CMD ["/app/server"]
