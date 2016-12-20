# Pull base image
FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/merqure/backend

# Install our dependencies
RUN go get github.com/gorilla/mux  

# Install backend binary globally within container 
RUN go install github.com/merqure/backend

# Set binary as entrypoint
ENTRYPOINT /go/bin/backend

# Expose default port (8080)
EXPOSE 8080