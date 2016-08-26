# reqtrace

## Example
```
% go build
% ./reqtrace

% curl -X POST -d '{"drink": "coffee"}' -H "Content-Type: application/json" http://localhost:8080?name=me

Remote address: [::1]:49782
Request URI:    /?name=me
Referer:
User agent:     curl/7.43.0
Host:           localhost:8080
Content-Type:   application/json
Method:         POST
Body:           {"drink": "coffee"}

```
