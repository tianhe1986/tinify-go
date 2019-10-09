package tinify

import (
	"io/ioutil"
	"net/http"
)

type Result struct {
	meta http.Header
	data []byte
}

//create a new result class
func NewResult(meta http.Header, data []byte) *Result {
	result := &Result{
		meta: meta,
		data: data,
	}

	return result
}

// write to file
func (r *Result) ToFile(path string) error {
	return ioutil.WriteFile(path, r.data, 0666)
}

func (r *Result) ToBuffer() []byte {
	return r.data
}

func (r *Result) GetMeta() http.Header {
	return r.meta
}
