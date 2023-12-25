FROM golang:1.21-alpine AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver .

# Dev image
FROM golang:1.21-alpine AS Dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]

# Production image
FROM scratch As Prod

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/apiserver", "/"]

# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]
