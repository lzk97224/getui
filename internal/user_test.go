package internal

import (
	"fmt"
	"testing"
)

func TestUser_UnbindByAlias(t *testing.T) {
	client := NewClient("0ugRpERCPX7L19GDzM5OZ4", "SUJrLwOCklAg1m3FTniDi", "aAy7QpGj4C8WfR3WrVUH98", func(appId, token string) {
		return
	}, func(appId string) string {
		return ""
	})
	fmt.Println(client.User.UnbindByAlias("sladjflajsdklf"))
}
