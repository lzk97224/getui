package internal

import (
	"crypto/sha256"
	"fmt"
	"path"
	"time"
)

type BaseResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Client struct {
	appId        string
	appKey       string
	masterSecret string
}

func NewClient(appId, appKey, masterSecret string) *Client {
	return &Client{
		appId:        appId,
		appKey:       appKey,
		masterSecret: masterSecret,
	}
}

func (c *Client) getUrl(paths ...string) string {
	return BASE_URL + path.Join(path.Join("/", c.appId), path.Join(paths...))
}

func (c *Client) sign() (string, string) {
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())

	original := c.appKey + timestamp + c.masterSecret

	hash := sha256.New()
	hash.Write([]byte(original))
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x", sum), timestamp
}
