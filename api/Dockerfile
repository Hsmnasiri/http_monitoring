FROM golang:1.16-buster as build

ENV GOPROXY http://proxy.golang.org
# ENV GO111MODULE=on


WORKDIR /app

COPY go.mod ./
COPY go.sum ./

# RUN rm $GOPATH/go.mod

RUN go mod download

COPY . .

# RUN mkdir bin

RUN go build -o /docker-gs-ping


FROM alpine
RUN apk add --no-cache bash

WORKDIR /app

COPY --from=build /app .

EXPOSE 3000

CMD exec /app
