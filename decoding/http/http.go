package http

import (
	"bufio"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/yazgazan/backcomp/decoding"
)

func Decode(r io.Reader) ([]decoding.Request, error) {
	b := bufio.NewReader(r)
	req, err := http.ReadRequest(b)
	if err != nil {
		return nil, err
	}

	// http.ReadRequest doesn't read the body, reading the left-overs
	body, err := ioutil.ReadAll(b)
	return []decoding.Request{
		{
			Method: req.Method,
			URL:    req.URL.String(),
			Header: req.Header,
			Body:   body,
		},
	}, nil
}
