# ----- build stage -----
FROM golang:1.24 AS build

WORKDIR /app

# 1) Go module files and deps
COPY go.mod go.sum ./
RUN go mod download

# 2) Copy the rest of the source
COPY . .

# 3) Build binaries for each cmd
RUN go build -o server ./cmd/server
RUN go build -o search ./cmd/search
RUN go build -o client ./cmd/client

# Default command (overridden by docker-compose per service)
CMD ["/bin/server"]