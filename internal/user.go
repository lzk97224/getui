package internal

import (
	"fmt"
	"github.com/lzk97224/getui/public/model"
)

type UserBindAliasReq struct {
	DataList []*model.UserBindAlias `json:"data_list"`
}

func (c *Client) UserBindAlias(dataList []*model.UserBindAlias) error {
	token := c.getToken(c.appId)
	resp := &BaseResp{}

	err := PostHeader(c.getUrl(PATH_USER_ALIAS), &UserBindAliasReq{
		DataList: dataList,
	}, resp, NewHeader().Add("token", token))

	if err != nil {
		return err
	}
	if resp.Code != CODE_SUCCESS {
		return fmt.Errorf("%v", resp.Msg)
	}
	return nil
}

type UserFindAliasByCidResp struct {
	BaseResp
	Data *UserFindAliasByCidData `json:"data"`
}

type UserFindAliasByCidData struct {
	Alias string `json:"alias"`
}

func (c *Client) UserFindAliasByCid(cid string) (string, error) {
	token := c.getToken(c.appId)
	resp := &UserFindAliasByCidResp{}

	err := GetHeader(c.getUrl(PATH_USER_ALIAS_CID, cid), resp, NewHeader().Add("token", token))

	if err != nil {
		return "", err
	}
	if resp.Code != CODE_SUCCESS {
		return "", fmt.Errorf("%v", resp.Msg)
	}
	return resp.Data.Alias, nil
}

type UserFindCidByAliasResp struct {
	BaseResp
	Data *UserFindCidByAliasData `json:"data"`
}

type UserFindCidByAliasData struct {
	Cid []string `json:"cid"`
}

func (c *Client) UserFindCidByAlias(alias string) ([]string, error) {
	token := c.getToken(c.appId)
	resp := &UserFindCidByAliasResp{}

	err := GetHeader(c.getUrl(PATH_USER_CID_ALIAS, alias), resp, NewHeader().Add("token", token))

	if err != nil {
		return nil, err
	}
	if resp.Code != CODE_SUCCESS {
		return nil, fmt.Errorf("%v", resp.Msg)
	}
	return resp.Data.Cid, nil
}
