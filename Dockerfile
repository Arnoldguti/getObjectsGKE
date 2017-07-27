# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# File Author / Maintainer
MAINTAINER Arnol Gutierrez


# Copy the local package files to the container's workspace.
ADD . /go/src/microservices

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install microservices

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/microservices


# Document that the service listens on port 6060.
EXPOSE 6060
