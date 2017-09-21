FROM golang:1.9-alpine3.6

COPY . /go/src/github.com/dollarshaveclub/sysctl-write-docker
RUN go install github.com/dollarshaveclub/sysctl-write-docker

FROM alpine:3.6

COPY --from=0 /go/bin/sysctl-write-docker /bin/
CMD /bin/sysctl-write-docker