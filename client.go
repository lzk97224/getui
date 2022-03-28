package getui

import "github.com/lzk97224/getui/internal"

type Client struct {
	*internal.Client
}

func NewClient(appId, appKey, masterSecret string,
	storeToken internal.StoreFunc, getToken internal.GetFunc) *Client {
	return &Client{
		Client: internal.NewClient(appId, appKey, masterSecret, storeToken, getToken),
	}
}
