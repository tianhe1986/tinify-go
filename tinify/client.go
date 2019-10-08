package tinify

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const API_ENDPOINT = "https://api.tinify.com"

//TODO 证书使用， 代理地址使用
type Client struct {
	key    string         // 使用的api key
	capool *x509.CertPool // 使用的证书 pool
	proxy  string         // 使用的代理地址
}

func GetNewClient(key string, capath string, proxy string) *Client {
	c := new(Client)
	c.key = key
	c.proxy = proxy

	if capath != "" {
		cacert, _ := ioutil.ReadFile(capath)
		c.capool = x509.NewCertPool()
		c.capool.AppendCertsFromPEM(cacert)
	}

	return c
}

// 发起请求，返回 response
func (c *Client) Request(method string, requestUrl string, body interface{}) *http.Response {

	//url 格式化处理
	if !strings.HasPrefix(requestUrl, "https:") {
		requestUrl = API_ENDPOINT + requestUrl
	}

	var reader io.Reader

	//body体处理，根据传入的body类型进行封装
	var jsonHeader bool = false
	switch body.(type) {
	case []byte: //读取的文件数据，将整个数据作为body直接发送
		reader = bytes.NewReader(body.([]byte))
	case map[string]interface{}: //作为json处理
		jsonHeader = true
		jsonStr, _ := json.Marshal(body)
		reader = bytes.NewReader(jsonStr)
	}

	//发起请求
	var transport *http.Transport
	if c.proxy == "" { //没有代理
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: c.capool,
			},
		}
	} else {
		proxyUrl, _ := url.Parse(c.proxy)
		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
			TLSClientConfig: &tls.Config{
				RootCAs: c.capool,
			},
		}
	}

	client := &http.Client{
		Transport: transport,
	}

	req, _ := http.NewRequest(method, requestUrl, reader)

	if jsonHeader {
		req.Header.Set("Content-Type", "application/json")
	}

	req.SetBasicAuth("api", c.key)

	response, _ := client.Do(req)

	countStr := response.Header.Get("compression-count")
	if countStr != "" {
		count, err := strconv.Atoi(countStr)
		if err == nil {
			setComCount(count)
		}
	}

	return response
}
