FROM golang:latest

LABEL maintainer="Clivern <hello@clivern.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]