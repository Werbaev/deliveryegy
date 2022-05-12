FROM golang:1.16 as builder


#
RUN mkdir -p $GOPATH/src/gitlab.udevs.io/delever/delever_delivery
WORKDIR $GOPATH/src/gitlab.udevs.io/delever/delever_delivery

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/delever_delivery /



FROM alpine
COPY --from=builder delever_delivery .
ENTRYPOINT ["/delever_delivery"]




