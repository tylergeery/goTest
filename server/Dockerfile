# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD ./server /go/src/github.com/gofun-server

# Build the gofun go app
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/gofun-server

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/gofun-server

# Document that the service listens on port 8080.
EXPOSE 8080
