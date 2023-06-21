# Build the application from source
FROM golang:latest AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY activities ./activities
COPY rest ./rest
RUN mkdir -p /tmp/activities
RUN CGO_ENABLED=0 GOOS=linux go build -o /grmn-server

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /grmn-server /grmn-server

EXPOSE 8088

USER nonroot:nonroot

ENTRYPOINT ["/grmn-server]
