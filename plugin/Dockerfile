FROM golang:1.19-alpine3.17
WORKDIR /go/src/myplugin
ADD ./ /go/src/myplugin

RUN go build -o /myplugin /go/src/myplugin/main.go
CMD ["/myplugin"]
