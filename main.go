package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Content struct {
	Body        string
	ContentType string
	Host        string
	Method      string
	Referer     string
	RemoteAddr  string
	RequestURI  string
	URL         string
	UserAgent   string
}

var content = `
Remote address: {{.RemoteAddr}}
Request URI:    {{.RequestURI}}
Referer:        {{.Referer}}
User agent:     {{.UserAgent}}
Host:           {{.Host}}
Content-Type:   {{.ContentType}}
Method:         {{.Method}}
Body:           {{.Body}}
`

func handler(w http.ResponseWriter, r *http.Request) {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	body := bufbody.String()

	trace := Content{
		Body:        body,
		ContentType: r.Header.Get("Content-Type"),
		Host:        r.Host,
		Method:      r.Method,
		Referer:     r.Referer(),
		RemoteAddr:  r.RemoteAddr,
		RequestURI:  r.RequestURI,
		URL:         r.URL.Path,
		UserAgent:   r.UserAgent(),
	}
	t := template.Must(template.New("content").Parse(content))
	err := t.Execute(os.Stdout, trace)
	if err != nil {
		fmt.Println("execute template:", err)
	}

	// Return nothing
	fmt.Fprintf(w, "")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running on 8080 port")

	http.ListenAndServe(":8080", nil)
}
