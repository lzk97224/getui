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

type StoreFunc func(appId, token string)
type GetFunc func(appId string) string

type Client struct {
	User         *User
	Push         *Push
	Auth         *Auth
	appId        string
	appKey       string
	masterSecret string
	storeToken   StoreFunc
	getToken     GetFunc
}

func NewClient(appId, appKey, masterSecret string, storeToken StoreFunc, getToken GetFunc) *Client {
	c := &Client{
		appId:        appId,
		appKey:       appKey,
		masterSecret: masterSecret,
		storeToken:   storeToken,
	}
	c.User = &User{Client: c}
	c.Push = &Push{Client: c}
	c.Auth = &Auth{Client: c}
	c.getToken = func(appId string) string {
		token := getToken(appId)
		if len(token) <= 0 {
			auth, err := c.Auth.Auth()
			if err != nil {
				return ""
			}
			token = auth.Token
			c.storeToken(c.appId, auth.Token)
		}
		return token
	}
	return c
}

func (c *Client) getUrl(paths ...string) string {
	return BASE_URL + path.Join(path.Join("/", c.appId), path.Join(paths...))
}

func (c *Client) sign() (string, string) {
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))

	original := c.appKey + timestamp + c.masterSecret

	hash := sha256.New()
	hash.Write([]byte(original))
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x", sum), timestamp
}
