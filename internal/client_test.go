package internal

import (
	"fmt"
	"github.com/lzk97224/getui/public/model"
	"testing"
)

var client *Client

func init() {
	client = &Client{
		appId:        "appid",
		appKey:       "appKey",
		masterSecret: "masterSecret",
	}
}

func TestClient_sign(t *testing.T) {
	c := &Client{
		appId:        "a",
		appKey:       "b",
		masterSecret: "c",
	}
	got, got1 := c.sign()
	fmt.Println(got, got1)
}

func TestClient_Auth(t *testing.T) {
	c := &Client{
		appId:        "appid",
		appKey:       "appKey",
		masterSecret: "masterSecret",
	}
	auth, err := c.Auth()
	fmt.Println(auth, err)
}

func TestClient_DeleteAuth(t *testing.T) {
	c := &Client{
		appId:        "appid",
		appKey:       "appKey",
		masterSecret: "masterSecret",
	}
	err := c.AuthDelete()
	fmt.Println(err)
}
func TestClient_PushSingleByCid(t *testing.T) {
	err := client.PushSingleByCid(
		"ec9d386372f37b299cd801bfa3df8ae81", "ec9d386372f37b299cd801bfa3df8ae8", &model.Notification{
			Title:     "这是标题",
			Body:      "这是我要提示你的内容，快来看看吧。",
			ClickType: "startapp",
		})
	fmt.Println(err)
}

func TestClient_PushBatchByCid(t *testing.T) {
	err := client.PushSingleBatchByCid([]string{"ec9d386372f37b299cd801bfa3df8ae8"}, []string{"ec9d386372f37b299cd801bfa3df8ae8"}, &model.Notification{
		Title:     "这是标题",
		Body:      "这是我要提示你的内容，快来看看吧。",
		ClickType: "startapp",
		Payload:   "这是我要提示你的内容，快来看看吧",
	})
	fmt.Println(err)
}

func TestClient_Test(t *testing.T) {

}
