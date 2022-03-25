package getui

import "github.com/lzk97224/getui/internal"

//type Client struct {
//	internal.Client
//}

func NewClient(appId, appKey, masterSecret string) *internal.Client {
	return internal.NewClient(appId, appKey, masterSecret)
}
