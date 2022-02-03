FROM golang:1.16-buster as build

ENV GOPROXY http://proxy.golang.org

WORKDIR /api

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /bin/app

FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

COPY --from=build /bin/app .

EXPOSE 3000

CMD exec /bin/app
