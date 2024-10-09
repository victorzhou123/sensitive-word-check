FROM golang:1.21 as BUILDER

# build binary
COPY . /go/src/github.com/victorzhou123/sensitive-word-check
RUN cd /go/src/github.com/victorzhou123/sensitive-word-check && GO111MODULE=on CGO_ENABLED=0 go build

# copy binary config and utils
FROM alpine:latest
WORKDIR /opt/app/

COPY  --from=BUILDER /go/src/github.com/victorzhou123/sensitive-word-check/sensitive-word-check /opt/app

ENTRYPOINT ["/opt/app/sensitive-word-check"]
