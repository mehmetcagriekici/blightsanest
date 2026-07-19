# ----- build stage -----
FROM golang:1.24 AS build

WORKDIR /app

# 1) Go module files and deps
COPY go.mod go.sum ./
RUN go mod download

# 2) Copy the rest of the source
COPY . .

# 3) Build binaries for each cmd
ENV CGO_ENABLED=0
RUN go build -o /out/server ./cmd/server
RUN go build -o /out/search ./cmd/search
RUN go build -o /out/client ./cmd/client
RUN go build -o /out/migrate ./cmd/migrate

# ----- runtime stage -----
# distroless: no package manager needed, ships CA certs and a non-root user already
FROM gcr.io/distroless/static-debian12:nonroot AS runtime

WORKDIR /app
COPY --from=build /out/server /out/search /out/client /out/migrate ./

USER nonroot:nonroot
