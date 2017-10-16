# reqtrace

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/mttrs/reqtrace/tree/master)

## Install
```
% go get github.com/mttrs/reqtrace
```

## Build & Run
```
% cd $GOPATH/src/github.com/mttrs/reqtrace
% go build
% ./reqtrace
Running on 8080 port
```

### Request trace sample
```
% curl -X POST -d '{"drink": "coffee"}' -H "Content-Type: application/json" http://localhost:8080?name=me

POST /?name=me HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 19
Content-Type: application/json
User-Agent: curl/7.54.0

{"drink": "coffee"}
```
