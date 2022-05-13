package internal

import (
	"fmt"
	"github.com/lzk97224/getui/getui"
)

type User struct {
	*Client
}

type UserBindAliasReq struct {
	DataList []*getui.UserBindAlias `json:"data_list"`
}

//BindAlias 绑定别名
func (c *User) BindAlias(dataList []*getui.UserBindAlias) error {
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

//FindAliasByCid 根据cid查询别名
func (c *User) FindAliasByCid(cid string) (string, error) {
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

//FindCidByAlias 根据别名查询cid
func (c *User) FindCidByAlias(alias string) ([]string, error) {
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

//BatchUnbindAlias 批量解绑别名
func (c *User) BatchUnbindAlias(dataList []*getui.UserBindAlias) error {
	token := c.getToken(c.appId)
	resp := &BaseResp{}

	err := DeleteHeader(c.getUrl(PATH_BATCH_UNBIND_ALIAS), &UserBindAliasReq{
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

//UnbindByAlias 解绑所有与该别名绑定的cid
func (c *User) UnbindByAlias(alias string) error {
	token := c.getToken(c.appId)
	resp := &BaseResp{}

	err := DeleteHeader(c.getUrl(PATH_UNBIND_BY_ALIAS, alias), nil, resp, NewHeader().Add("token", token))

	if err != nil {
		return err
	}
	if resp.Code != CODE_SUCCESS {
		return fmt.Errorf("%v", resp.Msg)
	}
	return nil
}

//BindUserToTags 一个用户绑定一批标签
//此接口对单个cid有频控限制，每天只能修改一次，最多设置100个标签；单个标签长度最大为32字符，标签总长度最大为512个字符，申请修改请点击右侧“技术咨询”了解详情 。
func (c *User) BindUserToTags(cid string, tags []string) error {
	token := c.getToken(c.appId)
	resp := &BaseResp{}

	req := struct {
		CustomTag []string `json:"custom_tag"`
	}{
		CustomTag: tags,
	}

	err := PostHeader(c.getUrl(PATH_BIND_TAG_USER_TO_TAGS, cid), req, resp, NewHeader().Add("token", token))

	if err != nil {
		return err
	}
	if resp.Code != CODE_SUCCESS {
		return fmt.Errorf("%v", resp.Msg)
	}
	return nil
}

//BindUsersToTag 一批用户绑定一个标签，此接口为增量
//此接口有频次控制(每分钟最多100次，每天最多10000次)，申请修改请点击右侧“技术咨询”了解详情
func (c *User) BindUsersToTag(tag string, cids []string) error {
	token := c.getToken(c.appId)
	resp := &BaseResp{}

	req := struct {
		Cid []string `json:"cid"`
	}{
		Cid: cids,
	}

	err := PostHeader(c.getUrl(PATH_BIND_TAG_USERS_TO_TAG, tag), req, resp, NewHeader().Add("token", token))

	if err != nil {
		return err
	}
	if resp.Code != CODE_SUCCESS {
		return fmt.Errorf("%v", resp.Msg)
	}
	return nil
}

//https://docs.getui.com/getui/server/rest_v2/user/
