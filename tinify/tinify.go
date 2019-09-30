package tinify

import (
	"io/ioutil"
	"net/http"
)

const VERSION = "0.0.1"

var key string = ""    // 使用的Key
var capath string = "" // 使用的capath
var client *Client = nil
var comCount int = 0 // 当前压缩次数

func SetKey(newKey string) {
	key = newKey
}

func GetKey() string {
	return key
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

	client = GetNewClient(key, capath)
	return client
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
