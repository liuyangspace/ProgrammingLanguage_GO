package net

import (
	"net/http"
	"io"
	"context"
	"crypto/tls"
	"net/url"
	"mime/multipart"
)

/*

 */

type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
	Header http.Header
	Body io.ReadCloser
	ContentLength int64
	TransferEncoding []string
	Close bool
	Uncompressed bool
	Trailer http.Header
	Request *http.Request
	TLS *tls.ConnectionState
}
type Request struct {
	Method string
	URL *url.URL
	Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0
	Header http.Header
	Body io.ReadCloser
	GetBody func() (io.ReadCloser, error)
	ContentLength int64
	TransferEncoding []string
	Close bool
	Host string
	Form url.Values
	PostForm url.Values
	MultipartForm *multipart.Form
	Trailer http.Header
	RemoteAddr string
	RequestURI string
	TLS *tls.ConnectionState
	Cancel <-chan struct{}
	Response *Response
	ctx context.Context
}
//client
func Head(url string) (resp *http.Response, err error) { return http.Head(url) }
//server
func ListenAndServe(addr string, handler http.Handler) error { return http.ListenAndServe(addr,handler) }
func (r *Request) FormValue(key string) string { return r.FormValue(key) }
//
func HandleFunc(pattern string, handler func(http.ResponseWriter, *Request)) { http.HandleFunc(pattern,nil) }