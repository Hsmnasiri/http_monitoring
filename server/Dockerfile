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


# FROM alpine
# RUN apk add --no-cache bash

# WORKDIR /app

# COPY --from=build /app /app

# RUN chmod +x /app/main

# ENTRYPOINT ["/app/main"]

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /docker-gs-ping /docker-gs-ping

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-ping"]