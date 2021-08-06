# GENERATE THE OPENAPI SPECIFICATION
FROM quay.io/goswagger/swagger:latest as swagger

# Copy source files from host to the container
COPY . .

# Create and move to a working folder
WORKDIR /output

# Generate the swagger specification
RUN swagger generate spec -o ./swagger.json

# BUILD THE GO APPLICATION
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /go/src/app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go get -d -v ./...
RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest

ENTRYPOINT ["app"]

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/app/app /usr/local/bin
COPY --from=swagger /output/swagger.json /usr/local/bin