FROM golang:1.22.5

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app with the correct architecture
RUN GOOS=linux GOARCH=amd64 go build -o myapp .

# Command to run the executable
CMD ["./myapp"]