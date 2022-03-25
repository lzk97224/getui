package getui

import (
	"crypto/sha256"
	"fmt"
	"github.com/lzk97224/getui/internal"
	"log"
	"path"
	"time"
)

type BaseResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Client struct {
	AppId        string
	AppKey       string
	MasterSecret string
}

func (c *Client) getUrl(paths ...string) string {
	return internal.BASE_URL + path.Join(path.Join("/", c.AppId), path.Join(paths...))
}

func (c *Client) sign() (string, string) {
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())

	original := c.AppKey + timestamp + c.MasterSecret

	hash := sha256.New()
	hash.Write([]byte(original))
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x", sum), timestamp
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

func (c *Client) Auth() (*AuthRespData, error) {
	sign, i := c.sign()

	resp := &AuthResp{}

	err := internal.Post(c.getUrl(internal.PATH_AUTH), AuthReq{
		Sign:      sign,
		Timestamp: i,
		Appkey:    c.AppKey,
	}, resp)

	if err != nil {
		return nil, err
	}
	if resp.Code != internal.CODE_SUCCESS {
		return nil, fmt.Errorf("%v", resp.Msg)
	}
	return &resp.Data, nil
}

func (c *Client) DeleteAuth(token string) error {

	resp := &BaseResp{}

	err := internal.Delete(c.getUrl(internal.PATH_AUTH, token), struct{}{}, resp)

	if err != nil {
		return err
	}
	if resp.Code != internal.CODE_SUCCESS {
		return fmt.Errorf("%v", resp.Msg)
	}
	return nil
}

type PushSingleReq struct {
	RequestId   string       `json:"request_id,omitempty"`   //请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	GroupName   string       `json:"group_name,omitempty"`   //任务组名。多个消息任务可以用同一个任务组名，后续可根据任务组名查询推送情况（长度限制100字符，且不能含有特殊符号）
	Audience    *Audience    `json:"audience,omitempty"`     //推送目标用户
	Settings    *Settings    `json:"settings,omitempty"`     //推送条件设置
	PushMessage *PushMessage `json:"push_message,omitempty"` //个推推送消息参数
	PushChannel *PushChannel `json:"push_channel,omitempty"`
}
type PushSingleResp struct {
	BaseResp
	Data PushSingleData `json:"data"`
}
type PushSingleData map[string]any

func (c *Client) PushSingleByCid(token string, requestId, cid string, notification *Notification) error {

	resp := &PushSingleResp{}

	err := internal.PostHeader(c.getUrl(internal.PATH_PUSH_SINGLE_CID), PushSingleReq{
		RequestId: requestId,
		Audience: &Audience{
			Cid: []string{cid},
		},
		PushMessage: &PushMessage{
			Notification: notification,
		},
		PushChannel: &PushChannel{
			Ios: nil,
			Android: &Android{Ups: &Ups{
				Notification: notification,
			}},
		},
	}, resp, internal.NewHeader().Add("token", token))

	if err != nil {
		return err
	}
	if resp.Code != internal.CODE_SUCCESS {
		return fmt.Errorf("%v", resp.Msg)
	}
	return nil
}

type PushBatchReq struct {
	IsAsync bool            `json:"is_async"`
	MsgList []PushSingleReq `json:"msg_list"`
}

func (c *Client) PushSingleBatchByCid(token string, requestId []string, cid []string, notification *Notification) error {

	resp := &BaseResp{}

	if len(requestId) != len(cid) {
		return fmt.Errorf("parrams error")
	}

	pageSize := 200
	requestIdGroup := internal.SliceSplit[string](pageSize, requestId)
	cidGroup := internal.SliceSplit(pageSize, cid)

	for index, ids := range requestIdGroup {
		msgList := make([]PushSingleReq, 0, pageSize)
		for i, id := range ids {
			msgList = append(msgList, PushSingleReq{
				RequestId: id,
				Audience: &Audience{
					Cid: []string{cidGroup[index][i]},
				},
				PushMessage: &PushMessage{
					Notification: notification,
				},
			})
		}
		err := internal.PostHeader(c.getUrl(internal.PATH_PUSH_SINGLE_BATCH_CID), &PushBatchReq{
			IsAsync: false,
			MsgList: msgList,
		}, resp, internal.NewHeader().Add("token", token))
		if err != nil {
			log.Printf("request err:%v", err)
		}
	}
	return nil
}

type PushBatchCreateMsgReq struct {
	RequestId   string       `json:"request_id,omitempty"`   //请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	GroupName   string       `json:"group_name,omitempty"`   //任务组名。多个消息任务可以用同一个任务组名，后续可根据任务组名查询推送情况（长度限制100字符，且不能含有特殊符号）
	Settings    *Settings    `json:"settings,omitempty"`     //推送条件设置
	PushMessage *PushMessage `json:"push_message,omitempty"` //个推推送消息参数
	PushChannel *PushChannel `json:"push_channel,omitempty"`
}

type PushBatchCreateMsgResp struct {
	BaseResp
	Taskid string `json:"taskid"`
}

func (c *Client) PushBatchCreateMsg(token string, requestId, groupName string, notification *Notification) (string, error) {
	resp := &PushBatchCreateMsgResp{}

	err := internal.PostHeader(c.getUrl(internal.PATH_PUSH_BATCH_CREATE_MSG), &PushBatchCreateMsgReq{
		RequestId: requestId,
		GroupName: groupName,
		PushMessage: &PushMessage{
			Notification: notification,
		},
		PushChannel: &PushChannel{
			Ios: nil,
			Android: &Android{Ups: &Ups{
				Notification: notification,
			}},
		},
	}, resp, internal.NewHeader().Add("token", token))

	if err != nil {
		return "", err
	}
	if resp.Code != internal.CODE_SUCCESS {
		return "", fmt.Errorf("%v", resp.Msg)
	}
	return resp.Taskid, nil
}

func (c *Client) PushBatchByCid(token string, taskId string, cid []string) error {
	req := &struct {
		Audience *Audience `json:"audience"`
		Taskid   string    `json:"taskid"`
		IsAsync  bool      `json:"is_async"`
	}{
		Audience: &Audience{Cid: cid},
		Taskid:   taskId,
		IsAsync:  true,
	}
	resp := &BaseResp{}

	err := internal.PostHeader(c.getUrl(internal.PATH_PUSH_BATCH_BY_CID), req, resp, internal.NewHeader().Add("token", token))

	if err != nil {
		return err
	}
	if resp.Code != internal.CODE_SUCCESS {
		return fmt.Errorf("%v", resp.Msg)
	}
	return nil
}

func (c *Client) PushBatchByAlias(token string, taskId string, alias []string) error {
	req := &struct {
		Audience *Audience `json:"audience"`
		Taskid   string    `json:"taskid"`
		IsAsync  bool      `json:"is_async"`
	}{
		Audience: &Audience{Alias: alias},
		Taskid:   taskId,
		IsAsync:  true,
	}
	resp := &BaseResp{}

	err := internal.PostHeader(c.getUrl(internal.PATH_PUSH_BATCH_BY_ALIAS), req, resp, internal.NewHeader().Add("token", token))

	if err != nil {
		return err
	}
	if resp.Code != internal.CODE_SUCCESS {
		return fmt.Errorf("%v", resp.Msg)
	}
	return nil
}

func (c *Client) PushCreateMsgAndBatchByAlias(token string, requestId, groupName string, notification *Notification, alias []string) (string, error) {
	taskId, err := c.PushBatchCreateMsg(token, requestId, groupName, notification)
	if err != nil {
		return "", err
	}
	return taskId, c.PushBatchByAlias(token, taskId, alias)
}

func (c *Client) PushCreateMsgAndBatchByCid(token string, requestId, groupName string, notification *Notification, cid []string) (string, error) {
	taskId, err := c.PushBatchCreateMsg(token, requestId, groupName, notification)
	if err != nil {
		return "", err
	}
	return taskId, c.PushBatchByCid(token, taskId, cid)
}
