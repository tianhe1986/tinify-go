package tinify

import (
	"io/ioutil"
)

type Result struct {
	meta map[string][]string
	data []byte
}

func NewResult(meta map[string][]string, data []byte) *Result {
	result := &Result{
		meta: meta,
		data: data,
	}

	return result
}

func (r *Result) ToFile(path string) {
	err := ioutil.WriteFile(path, r.data, 0666)
	if err != nil {
		return
	}
}
