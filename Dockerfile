# For build
FROM golang:alpine AS build

ARG WORK_DIR=/go/src/github.com/mttrs/reqtrace/
WORKDIR $WORK_DIR
COPY . $WORK_DIR
RUN cd $WORK_DIR && go build -o reqtrace

# For run
FROM alpine
ARG BUILD_DIR=/go/src/github.com/mttrs/reqtrace/
WORKDIR /app
COPY --from=build $BUILD_DIR/reqtrace /app/

# Run the image as a non-root user
RUN adduser -D user
USER user

CMD ["./reqtrace"]
