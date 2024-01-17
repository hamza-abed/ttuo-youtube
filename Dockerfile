FROM golang:1.21.5

RUN go install github.com/cespare/reflex@latest

WORKDIR $GOPATH/src/helloworld

EXPOSE 5000

CMD $GOPATH/bin/helloworld