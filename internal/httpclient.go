package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Header struct {
	headers http.Header
}

func NewHeader() *Header {
	return &Header{headers: http.Header{}}
}
func (h *Header) Add(key, value string) *Header {
	h.headers.Add(key, value)
	return h
}

func GetHeader(url string, respBody any, header *Header) error {
	return baseRequest(http.MethodGet, url, struct{}{}, respBody, header)
}

func PostHeader(url string, reqBody any, respBody any, header *Header) error {
	return baseRequest(http.MethodPost, url, reqBody, respBody, header)
}

func DeleteHeader(url string, reqBody any, respBody any, header *Header) error {
	return baseRequest(http.MethodDelete, url, reqBody, respBody, header)
}

func Post(url string, reqBody any, respBody any) error {
	return baseRequest(http.MethodPost, url, reqBody, respBody, nil)
}

func Delete(url string, reqBody any, respBody any) error {
	return baseRequest(http.MethodDelete, url, reqBody, respBody, nil)
}

func baseRequest(method, url string, reqBody any, respBody any, header *Header) error {
	//创建客户端实例
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	//创建请求实例
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if header != nil {
		for key, value := range header.headers {
			req.Header[key] = value
		}
	}

	req.Header.Add("Charset", "UTF-8")
	req.Header.Add("Content-Type", "application/json")

	//发起请求
	resp, err := client.Do(req)
	if err != nil {
		errLog(err, url, method, string(body), "")
		return err
	}

	defer resp.Body.Close()

	//读取响应
	result, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		errLog(err, url, method, string(body), string(result))
		return err
	}

	err = json.Unmarshal(result, respBody)
	if err != nil {
		log.Printf("response json fail ,err:%v", err)
		errLog(err, url, method, string(body), string(result))
		return err
	}
	return nil
}

func errLog(err error, url string, method string, body string, result string) {
	log.Printf("request getui fail ,err:%v", err)
	log.Printf("request url:%v", url)
	log.Printf("request method:%v", method)
	log.Printf("request body:%v", body)
	log.Printf("request result:%v", result)
}
