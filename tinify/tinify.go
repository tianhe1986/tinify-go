package tinify

import (
	"io/ioutil"
	"net/http"
	"path"
	"runtime"
)

const VERSION = "0.0.1"

var key string = ""    // 使用的Key
var capath string = "" // 使用的capath
var proxy string = ""  // 使用的代理url
var client *Client = nil
var comCount int = 0                       // 当前压缩次数
var mockClientFun func(*http.Client) = nil // 用于mock测试处理

func SetMockClientFun(f func(*http.Client)) {
	mockClientFun = f
}

func SetKey(newKey string) {
	key = newKey
}

func GetKey() string {
	return key
}

func SetProxy(newProxy string) {
	proxy = newProxy
}

func GetProxy() string {
	return proxy
}

func setComCount(val int) {
	comCount = val
}

func GetComCount() int {
	return comCount
}

// client单例
func GetClient() *Client {
	if client != nil {
		return client
	}

	tempPath := capath
	if tempPath == "" {
		_, tempfilename, _, _ := runtime.Caller(1)
		tempPath = path.Join(path.Dir(tempfilename), "../data/cacert.pem")
	}

	client = GetNewClient(key, tempPath, proxy)
	return client
}

func SetClient(c *Client) {
	client = c
}

//从文件创建一个source
func FromFile(path string) *Source {
	buffer, _ := ioutil.ReadFile(path)
	return FromBuffer(buffer)
}

//从buffer创建一个source
func FromBuffer(buffer []byte) *Source {
	response := GetClient().Request(http.MethodPost, "/shrink", buffer)
	defer response.Body.Close()
	return NewSource(response.Header.Get("location"), make(map[string]interface{}))
}

//从url创建一个source
func FromUrl(url string) *Source {
	var body map[string]interface{} = make(map[string]interface{})
	body["source"] = map[string]string{"url": url}

	response := GetClient().Request(http.MethodPost, "/shrink", body)
	defer response.Body.Close()
	return NewSource(response.Header.Get("location"), make(map[string]interface{}))
}
