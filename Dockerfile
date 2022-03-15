FROM golang:1.18-rc

WORKDIR /usr/src/app

COPY go.* ./

RUN go mod download

COPY ./ ./
