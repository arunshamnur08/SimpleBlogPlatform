FROM golang:1.21.0-alpine

# Creat New Directory inside the container
RUN mkdir /app

ADD . /app

# Set the Current Working Directory inside the container

WORKDIR /app/cmd/

# Build the Go app
RUN go build -o simple-blog-platform .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary`

ENTRYPOINT ["./simple-blog-platform"]