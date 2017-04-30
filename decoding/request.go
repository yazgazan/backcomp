package decoding

import (
	"net/http"
)

type Request struct {
	Method string
	URL    string
	Header http.Header
	Body   []byte
}
