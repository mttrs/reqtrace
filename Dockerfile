# For build
FROM golang:alpine AS build

# Install dep required to build
# Need to run `docker build --no-cache .` to update those dependencies
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

ARG WORK_DIR=/go/src/github.com/mttrs/reqtrace/

# Gopkg.toml and Gopkg.lock lists project dependencies
# These layers are only re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml $WORK_DIR
WORKDIR $WORK_DIR
# Install library dependencies
RUN dep ensure -vendor-only

# Copy all project and build it
# This layer is rebuilt when ever a file has changed in the project directory
COPY . $WORK_DIR
RUN cd $WORK_DIR && go build -o reqtrace

# For run
FROM alpine
ARG BUILD_DIR=/go/src/github.com/mttrs/reqtrace/
WORKDIR /app
COPY --from=build $BUILD_DIR/reqtrace /app/

# Run the image as a non-root user
RUN adduser -D myuser
USER myuser

CMD ["./reqtrace"]
