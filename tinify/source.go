package tinify

import (
	"io/ioutil"
	"net/http"
)

type Source struct {
	Url      string                 // url
	Commands map[string]interface{} // 命令
}

// 创建一个新的source
func NewSource(url string, commands map[string]interface{}) *Source {
	source := &Source{
		Url:      url,
		Commands: commands,
	}

	return source
}

// 保存到文件
func (s *Source) ToFile(path string) {
	s.Result().ToFile(path)
}

func (s *Source) Result() *Result {
	response := GetClient().Request(http.MethodGet, s.Url, s.Commands)
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	return NewResult(response.Header, data)
}
