package har

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/yazgazan/backcomp/decoding"
)

type testFile struct {
	har      string
	expected string
}

type testData struct {
	name     string
	har      HAR
	expected []decoding.Request
}

var testFiles = []testFile{
	{har: "test_data/golang.org.har", expected: "test_data/golang.org.json"},
}

func TestDecode(t *testing.T) {
	for _, test := range testFiles {
		var expectations []decoding.Request
		f, err := os.Open(test.har)
		if err != nil {
			t.Errorf("Unexpected error opening %q: %s", test.har, err)
			continue
		}

		err = unmarshalFile(test.expected, &expectations)
		if err != nil {
			t.Errorf("Unexpected error unmarshaling %q: %s", test.expected, err)
			continue
		}

		reqs, err := Decode(f)
		if err != nil {
			t.Errorf("Unexpected error decoding %q: %s", test.har, err)
			f.Close()
			continue
		}

		if !reflect.DeepEqual(expectations, reqs) {
			t.Errorf("%q: Expectations not met", test.har)
		}
	}
}

func unmarshalFile(fname string, v interface{}) error {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}
