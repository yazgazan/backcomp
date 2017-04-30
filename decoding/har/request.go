package har

import (
	"github.com/yazgazan/backcomp/decoding"
	"net/http"
)

type Request struct {
	Method      string
	URL         string
	HTTPVersion string
	HeaderSize  int64
	BodySize    int64
	Headers     []KV
	QueryString []KV
	Cookies     []Cookie
	PostData    Content
	Content     Content `json:"content"`
}

type KV struct {
	Name  string
	Value string
}

type Cookie struct {
	Name     string
	Value    string
	HTTPOnly bool
	Secure   bool
	// Expires (type ?)
}

func (r Request) ToRequest() decoding.Request {
	return decoding.Request{
		Method: r.Method,
		URL:    r.URL,
		Header: kvToHeader(r.Headers),
		Body:   []byte(r.Content.Text),
	}
}

func kvToHeader(kvs []KV) http.Header {
	header := http.Header{}

	for _, kv := range kvs {
		header.Set(kv.Name, kv.Value)
	}

	return header
}
