package http

import (
	"bytes"
	"net/http"
	"testing"
)

const body = `POST /sgg/v2/api/eventsuggestions/ HTTP/1.1
Host: localhost:8389
Content-Type: application/json
Cache-Control: no-cache

{"evn_name":"Diner","evn_description":"","evn_category":4301,"geocode":"52.3629,4.8925","country":"NL","size":4}`

func TestDecode(t *testing.T) {
	r, err := Decode(bytes.NewBuffer([]byte(body)))
	if err != nil {
		t.Error("Unexpected error decoding request:", err)
		return
	}
	if len(r) != 1 {
		t.Errorf("expected Decode to return 1 request, got %d", len(r))
		return
	}
	req := r[0]

	if req.Method != http.MethodPost {
		t.Errorf("Expected method %s, got %s", http.MethodPost, req.Method)
	}

	wantURL := "/sgg/v2/api/eventsuggestions/"
	if req.URL != wantURL {
		t.Errorf("Expected url to be %q, got %q", wantURL, req.URL)
	}

	wantHeader := http.Header{}
	wantHeader.Set("Content-Type", "application/json")
	wantHeader.Set("Cache-Control", "no-cache")
	if key, want, have, ok := compareHeaders(req.Header, wantHeader); !ok {
		t.Errorf("Expected hearder %s to be %q, got %q", key, want, have)
	}

	wantBody := `{"evn_name":"Diner","evn_description":"","evn_category":4301,"geocode":"52.3629,4.8925","country":"NL","size":4}`
	if string(req.Body) != wantBody {
		t.Errorf("expected body to be %q, got %q", wantBody, body)
	}
}

func compareHeaders(have http.Header, want http.Header) (key, wantVal, haveVal string, ok bool) {
	ok = true
	for k := range want {
		if want.Get(k) != have.Get(k) {
			ok = false
			key = k
			wantVal = want.Get(k)
			haveVal = have.Get(k)
			return
		}
	}
	return
}
