# golang image where workspace (GOPATH) configured at /go
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/srodman7689/learning_go/web_development_in_go/chapter9/taskmanager/

# Setting up working directory
WORKDIR /go/src/github.com/srodman7689/learning_go/web_development_in_go/chapter9/taskmanager/

# Get godeps for managing and restoring dependencies
RUN go get github.com/tools/godep

# Restore godep dependencies
Run godep restore

# Build the taskmanager command inside the container.
RUN go install github.com/srodman7689/learning_go/web_development_in_go/chapter9/taskmanager

# Run the taskmanager command when the container starts.
ENTRYPOINT /go/bin/taskmanager

# Service listens on port 8080
EXPOSE 8080
