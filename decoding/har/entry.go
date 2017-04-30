package har

import (
	"github.com/yazgazan/backcomp/decoding"
	"time"
)

type Entry struct {
	StartedDateTime time.Time
	Time            float64
	ServerIPAddress string
	Connection      string
	PageRef         string
	Request         Request  `json:"request"`
	Response        Response `json:"response"`
	Cache           Cache    `json:"cache"`
	Timings         Timings  `json:"timings"`
}

type Response struct {
	Status       int
	StatusText   string
	HTTPVersion  string
	RedirectURL  string
	HeadersSize  int64
	BodySize     int64
	TransferSize int64 `json:"_transferSize"`
	Headers      []KV
	Cookies      []Cookie
	Content      Content `json:"content"`
}

type Content struct {
	Size        int64
	MimeType    string
	Text        string
	Compression int
}

type Cache struct {
}

type Timings struct {
	Blocked float64
	DNS     float64
	Connect float64
	Send    float64
	Wait    float64
	Receive float64
	SSL     float64
}

func (e Entry) ToRequest() decoding.Request {
	return e.Request.ToRequest()
}
