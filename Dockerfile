FROM golang:1.8

## Create a directory and Add Code
RUN mkdir -p /go/src/github.com/orvice/go-socket.io-bench
WORKDIR /go/src/github.com/orvice/go-socket.io-bench
ADD .  /go/src/github.com/orvice/go-socket.io-bench

# Download and install any required third party dependencies into the container.
RUN go-wrapper download
RUN go-wrapper install

# Now tell Docker what command to run when the container starts
CMD ["go-wrapper", "run"]