package tinify

import (
	"io/ioutil"
)

type Result struct {
	meta map[string][]string
	data []byte
}

//create a new result class
func NewResult(meta map[string][]string, data []byte) *Result {
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
