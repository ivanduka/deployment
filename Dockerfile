###################################################
# Building the app
###################################################
FROM golang:1.15.6-alpine3.12 AS build_base
RUN apk add upx

# Set the Current Working Directory inside the container
WORKDIR /tmp/deployment

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/deployment.app . && upx ./out/deployment.app

###################################################
# Building the healthcheck utility
###################################################
FROM golang:1.15.6-alpine3.12 AS build_base_healthcheck
RUN apk add upx

# Set the Current Working Directory inside the container
WORKDIR /tmp/healthcheck

COPY ./healthcheck/. .

# Build the Go app with 'no debugging info' flags
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/healthcheck.app . && upx ./out/healthcheck.app

###################################################
# Putting it all together in a new container
###################################################

# Start fresh from a smaller image
FROM alpine:3.12

COPY --from=build_base_healthcheck /tmp/healthcheck/out/healthcheck.app /app/healthcheck.app
HEALTHCHECK CMD /app/healthcheck.app

COPY --from=build_base /tmp/deployment/out/deployment.app /app/deployment.app

# This container exposes port 8080 to the outside world
EXPOSE 3333

# Run the binary program produced by `go build`
CMD ["/app/deployment.app"]
