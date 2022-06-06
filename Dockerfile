FROM golang:1.18-alpine

WORKDIR /usr/src
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY * ./

RUN go build -o /usr/bin/app

CMD [ "/usr/bin/app" ]