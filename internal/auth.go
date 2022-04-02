package internal

import "fmt"

type Auth struct {
	*Client
}

type AuthReq struct {
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	Appkey    string `json:"appkey"`
}

type AuthResp struct {
	BaseResp
	Data AuthRespData `json:"data"`
}
type AuthRespData struct {
	ExpireTime string `json:"expire_time"`
	Token      string `json:"token"`
}

func (c *Auth) Auth() (*AuthRespData, error) {
	sign, i := c.sign()

	resp := &AuthResp{}

	err := Post(c.getUrl(PATH_AUTH), AuthReq{
		Sign:      sign,
		Timestamp: i,
		Appkey:    c.appKey,
	}, resp)

	if err != nil {
		return nil, err
	}
	if resp.Code != CODE_SUCCESS {
		return nil, fmt.Errorf("%v", resp.Msg)
	}
	return &resp.Data, nil
}

func (c *Auth) Delete() error {
	token := c.getToken(c.appId)
	resp := &BaseResp{}

	err := Delete(c.getUrl(PATH_AUTH, token), struct{}{}, resp)

	if err != nil {
		return err
	}
	if resp.Code != CODE_SUCCESS {
		return fmt.Errorf("%v", resp.Msg)
	}
	return nil
}
