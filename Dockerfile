FROM golang:1.22 AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y nodejs npm

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN npm ci
RUN make build

FROM alpine:latest
WORKDIR /root/

RUN apk add --no-cache curl

COPY --from=builder /app/bin/main .

EXPOSE 8080

CMD ["./main"]
