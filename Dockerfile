# syntax=docker/dockerfile:1

FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /ringcentral-permahooks

EXPOSE 8080

CMD [ "/ringcentral-permahooks" ]
