package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Content struct {
	Body          string
	ContentType   string
	Host          string
	Method        string
	Referer       string
	RemoteAddr    string
	RequestURI    string
	URL           string
	UserAgent     string
	XForwardedFor string
}

var content = `
Remote address:  {{.RemoteAddr}}
Request URI:     {{.RequestURI}}
Referer:         {{.Referer}}
User agent:      {{.UserAgent}}
Host:            {{.Host}}
Content-Type:    {{.ContentType}}
Method:          {{.Method}}
Body:            {{.Body}}
X-Forwarded-For: {{.XForwardedFor}}
`

func handler(w http.ResponseWriter, r *http.Request) {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	body := bufbody.String()

	trace := Content{
		Body:          body,
		ContentType:   r.Header.Get("Content-Type"),
		Host:          r.Host,
		Method:        r.Method,
		Referer:       r.Referer(),
		RemoteAddr:    r.RemoteAddr,
		RequestURI:    r.RequestURI,
		URL:           r.URL.Path,
		UserAgent:     r.UserAgent(),
		XForwardedFor: r.Header.Get("X-Forwarded-For"),
	}
	t := template.Must(template.New("content").Parse(content))
	err := t.Execute(os.Stdout, trace)
	if err != nil {
		fmt.Println("execute template:", err)
	}

	// Return nothing
	fmt.Fprintf(w, "kk")
}

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Running on:", port, "...")

	http.ListenAndServe(":"+port, nil)
}
