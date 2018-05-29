FROM golang:1.10.2-stretch

WORKDIR /go/src/github/cdupree/time-api

COPY main.go .

RUN go get github.com/gorilla/mux && go build && echo 'built'

ENTRYPOINT ./time-api



