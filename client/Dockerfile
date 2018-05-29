FROM golang:1.10.2-stretch

WORKDIR /go/src/github/cdupree/time-api

COPY main.go .

RUN go build && echo 'built'

ENTRYPOINT ["./time-api"]



