# reqtrace

## Example
```
% go build
% ./reqtrace
Running on 8080 port
```

### Send a request
```
% curl -X POST -d '{"drink": "coffee"}' -H "Content-Type: application/json" http://localhost:8080?name=me

Remote address: [::1]:49782
Request URI:    /?name=me
Referer:
User agent:     curl/7.43.0
Host:           localhost:8080
Content-Type:   application/json
Method:         POST
Body:           {"drink": "coffee"}


# postman
http://localhost:8080?drink=beer&type=ipa

Remote address: [::1]:62717
Request URI:    /?drink=beer&type=ipa
Referer:
User agent:     Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36
Host:           localhost:8080
Content-Type:   application/json
Method:         GET
Body:
```
