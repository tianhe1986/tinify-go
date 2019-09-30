package tinify

import (
	"io/ioutil"
	"net/http"
)

type Source struct {
	Url      string                 // url
	Commands map[string]interface{} // 命令
}

type ResizeParam struct {
	Method string `json:"method"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

// 创建一个新的source
func NewSource(url string, commands map[string]interface{}) *Source {
	source := &Source{
		Url:      url,
		Commands: commands,
	}

	return source
}

// 重新调整大小
func (s *Source) Resize(method string, width int, height int) *Source {
	resizeParam := new(ResizeParam)
	resizeParam.Method = method
	resizeParam.Width = width
	resizeParam.Height = height

	s.Commands["resize"] = resizeParam
	return s
}

// 保存到文件
func (s *Source) ToFile(path string) {
	s.Result().ToFile(path)
}

// 获取结果
func (s *Source) Result() *Result {
	response := GetClient().Request(http.MethodGet, s.Url, s.Commands)
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	return NewResult(response.Header, data)
}
