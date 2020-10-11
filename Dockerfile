FROM golang:1.10

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/nullc0rp/go-ip2proxy-api

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8443

# Run the executable
CMD ["go-ip2proxy-api"]

