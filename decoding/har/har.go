package har

import (
	"encoding/json"
	"github.com/yazgazan/backcomp/decoding"
	"io"
	"io/ioutil"
	"time"
)

type HAR struct {
	Log Log `json:"log"`
}

type Log struct {
	Version string
	Creator Creator `json:"creator"`
	Pages   []Page
	Entries []Entry
}

type Creator struct {
	Name    string
	Version string
}

type Page struct {
	StartedDateTime time.Time
	ID              string
	Title           string
	PageTimings     PageTimings `json:"pageTimings"`
}

type PageTimings struct {
	OnContentLoad float64
	OnLoad        float64
}

func Decode(r io.Reader) ([]decoding.Request, error) {
	var har HAR

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &har)
	if err != nil {
		return nil, err
	}

	requests := make([]decoding.Request, len(har.Log.Entries))
	for i, entry := range har.Log.Entries {
		requests[i] = entry.ToRequest()
	}

	return requests, nil
}
