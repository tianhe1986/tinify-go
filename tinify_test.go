package tinifytest

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/tianhe1986/tinify-go/tinify"
)

// 这个测试没有什么用， 就是写着检查下是否正确
func TestKey(t *testing.T) {
	var key string = "ok your key"
	tinify.SetKey(key)
	assert := assert.New(t)
	// assert equality
	assert.Equal(key, tinify.GetKey(), "they should be equal")
}

// 自定义返回处理，由于需要设置header，因此自行封装一层
func NewStringHeaderResponder(status int, headers map[string]string, body string) httpmock.Responder {
	return func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(status, body)

		for key, value := range headers {
			resp.Header.Set(key, value)
		}
		return resp, nil
	}
}

// 根据buffer生成source， 用mock的方式处理
func TestSource(t *testing.T) {
	assert := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// 由于未使用default Transport，因此需要这样处理
	tinify.SetMockClientFun(httpmock.ActivateNonDefault)
	defer tinify.SetMockClientFun(nil)

	var location string = "uuu happy"
	mockResponseHeader := make(map[string]string)
	mockResponseHeader["location"] = location

	httpmock.RegisterResponder("POST", "https://api.tinify.com/shrink", NewStringHeaderResponder(200, mockResponseHeader, ""))

	var buffer []byte = make([]byte, 1)
	var source *tinify.Source = tinify.FromBuffer(buffer)

	assert.Equal(location, source.Url, "location should be the value set")

	info := httpmock.GetCallCountInfo()
	assert.Equal(1, info["POST https://api.tinify.com/shrink"], "url should be called only once")
}

// 压缩处理，检查发送的头和获取到的结果
func TestCompress(t *testing.T) {

}

// resize处理， 检查发送的头和获取到的结果
func TestResize(t *testing.T) {

}
