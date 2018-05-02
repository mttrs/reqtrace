# For build
FROM golang:alpine AS build

ADD . /src
RUN cd /src && go build -o reqtrace

# For run
FROM alpine
WORKDIR /app
COPY --from=build /src/reqtrace /app/
CMD ["./reqtrace"]
