FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/app .

RUN chmod +x ./app

EXPOSE 8080

ENTRYPOINT ["./app"]

