package internal

import (
	"fmt"
	"github.com/lzk97224/getui/public/model"
	"testing"
)

var client *Client

func init() {
	client = &Client{
		appId:        "A8DXHioQAH6O1YugIeK9e2",
		appKey:       "DM08GyByoO6QTTETV0DXY3",
		masterSecret: "P5TlL65c6ZAkqC22mzCGB2",
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
		appId:        "A8DXHioQAH6O1YugIeK9e2",
		appKey:       "DM08GyByoO6QTTETV0DXY3",
		masterSecret: "P5TlL65c6ZAkqC22mzCGB2",
	}
	auth, err := c.Auth()
	fmt.Println(auth, err)
}

func TestClient_DeleteAuth(t *testing.T) {
	c := &Client{
		appId:        "A8DXHioQAH6O1YugIeK9e2",
		appKey:       "DM08GyByoO6QTTETV0DXY3",
		masterSecret: "P5TlL65c6ZAkqC22mzCGB2",
	}
	err := c.AuthDelete("f3a67e46872907cf5cc11694d40fd483395132aef2b059c87ba729174e2026b6")
	fmt.Println(err)
}
func TestClient_PushSingleByCid(t *testing.T) {
	err := client.PushSingleByCid("3865f9358473b088a0b00aab2ce82f1bb99c1e2c0d11d31fc25c1db9d705a83d",
		"ec9d386372f37b299cd801bfa3df8ae81", "ec9d386372f37b299cd801bfa3df8ae8", &model.Notification{
			Title:     "这是标题",
			Body:      "这是我要提示你的内容，快来看看吧。",
			ClickType: "startapp",
		})
	fmt.Println(err)
}

func TestClient_PushBatchByCid(t *testing.T) {
	err := client.PushSingleBatchByCid("3865f9358473b088a0b00aab2ce82f1bb99c1e2c0d11d31fc25c1db9d705a83d", []string{"ec9d386372f37b299cd801bfa3df8ae8"}, []string{"ec9d386372f37b299cd801bfa3df8ae8"}, &model.Notification{
		Title:     "这是标题",
		Body:      "这是我要提示你的内容，快来看看吧。",
		ClickType: "startapp",
		Payload:   "这是我要提示你的内容，快来看看吧",
	})
	fmt.Println(err)
}

func TestClient_Test(t *testing.T) {

}
