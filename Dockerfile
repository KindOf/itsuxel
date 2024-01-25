# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api

# Run the tests in the container
FROM scratch
COPY --from=build-stage /api /api

WORKDIR /

EXPOSE 3000

ENTRYPOINT ["/api"]
