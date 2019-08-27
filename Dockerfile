FROM golang:latest

RUN mkdir -p /go/src/crud-golang

WORKDIR  /go/src/crud-golang

COPY . /go/src/crud-golang

RUN go install crud-golang

CMD /go/bin/crud-golang

EXPOSE 8080