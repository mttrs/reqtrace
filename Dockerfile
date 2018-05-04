# For build
FROM golang:alpine AS build

ADD . /src
RUN cd /src && go build -o reqtrace

# For run
FROM alpine
WORKDIR /app
COPY --from=build /src/reqtrace /app/

# Run the image as a non-root user
RUN adduser -D myuser
USER myuser

CMD ["./reqtrace"]
