package tinify

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// 这个测试没有什么用， 就是写着检查下是否正确
func TestKey(t *testing.T) {
	var key string = "ok your key"
	SetKey(key)
	assert := assert.New(t)
	// assert equality
	assert.Equal(key, GetKey(), "they should be equal")
}

// 自定义返回处理，由于需要设置header，因此自行封装一层
func newStringHeaderResponder(status int, headers map[string]string, body string) httpmock.Responder {
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
	SetMockClientFun(httpmock.ActivateNonDefault)
	defer SetMockClientFun(nil)

	var location string = "uuu happy"
	mockResponseHeader := make(map[string]string)
	mockResponseHeader["location"] = location

	httpmock.RegisterResponder("POST", "https://api.tinify.com/shrink", newStringHeaderResponder(200, mockResponseHeader, ""))

	var buffer []byte = make([]byte, 1)
	var source *Source = FromBuffer(buffer)

	assert.Equal(location, source.Url, "location should be the value set")

	assert.Equal(1, httpmock.GetTotalCallCount(), "request should be only once")
	info := httpmock.GetCallCountInfo()
	assert.Equal(1, info["POST https://api.tinify.com/shrink"], "shrink url should be called only once")
}

// 压缩处理，检查发送的头和获取到的结果
func TestCompress(t *testing.T) {
	assert := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// 由于未使用default Transport，因此需要这样处理
	SetMockClientFun(httpmock.ActivateNonDefault)
	defer SetMockClientFun(nil)

	var location string = "/ojbk.png"
	mockShrinkHeader := make(map[string]string)
	mockShrinkHeader["location"] = location

	//第一次请求shrink，返回图片URL
	httpmock.RegisterResponder("POST", "https://api.tinify.com/shrink", newStringHeaderResponder(200, mockShrinkHeader, ""))
	//第二次请求图片URL，返回数据
	var imgStr string = "your test passed"
	httpmock.RegisterResponder("GET", "https://api.tinify.com"+location, httpmock.NewStringResponder(200, imgStr))

	var buffer []byte = make([]byte, 1)
	result := FromBuffer(buffer).Result()
	assert.IsType(new(Result), result, "should be a Result pointer instance")

	var resultStr string = string(result.ToBuffer())
	assert.Equal(imgStr, resultStr, "image data should be the same")

	assert.Equal(2, httpmock.GetTotalCallCount(), "request should be 2")
	info := httpmock.GetCallCountInfo()
	assert.Equal(1, info["POST https://api.tinify.com/shrink"], "shrink url should be called only once")
	assert.Equal(1, info["GET https://api.tinify.com"+location], "image url should be called only once")
}

// resize处理， 检查发送的头和获取到的结果
func TestResize(t *testing.T) {

}
