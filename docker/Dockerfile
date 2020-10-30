FROM golang:1.13 as builder
MAINTAINER system@ajira.com

COPY go.mod /go/src/ajiiranetservice/
COPY go.sum /go/src/ajiiranetservice/
RUN cd /go/src/ajiiranetservice/ && go mod download

ADD . / /go/src/ajiiranetservice/

WORKDIR /go/src/ajiiranetservice
RUN go mod tidy && \
    go get -d -v 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ajiiranetservice .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir -p /ajira/ajiiranetservice
WORKDIR /ajira/ajiiranetservice
COPY --from=builder /go/src/ajiiranetservice/ .
RUN cd /ajira/ajiiranetservice && rm *.go
CMD ["/bin/sh","-c","/ajira/ajiiranetservice/ajiiranetservice"]