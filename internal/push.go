package internal

import (
	"fmt"
	"github.com/lzk97224/getui/public/model"
	"log"
)

type PushSingleReq struct {
	RequestId   string             `json:"request_id,omitempty"`   //请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	GroupName   string             `json:"group_name,omitempty"`   //任务组名。多个消息任务可以用同一个任务组名，后续可根据任务组名查询推送情况（长度限制100字符，且不能含有特殊符号）
	Audience    *model.Audience    `json:"audience,omitempty"`     //推送目标用户
	Settings    *model.Settings    `json:"settings,omitempty"`     //推送条件设置
	PushMessage *model.PushMessage `json:"push_message,omitempty"` //个推推送消息参数
	PushChannel *model.PushChannel `json:"push_channel,omitempty"`
}
type PushSingleResp struct {
	BaseResp
	Data PushSingleData `json:"data"`
}
type PushSingleData map[string]any

func (c *Client) PushSingleByCid(requestId, cid string, notification *model.Notification) error {
	token := c.getToken(c.appId)

	resp := &PushSingleResp{}

	err := PostHeader(c.getUrl(PATH_PUSH_SINGLE_CID), PushSingleReq{
		RequestId: requestId,
		Audience: &model.Audience{
			Cid: []string{cid},
		},
		PushMessage: &model.PushMessage{
			Notification: notification,
		},
		PushChannel: &model.PushChannel{
			Ios: nil,
			Android: &model.Android{Ups: &model.Ups{
				Notification: notification,
			}},
		},
	}, resp, NewHeader().Add("token", token))

	if err != nil {
		return err
	}
	if resp.Code != CODE_SUCCESS {
		return fmt.Errorf("%v", resp.Msg)
	}
	return nil
}

type PushBatchReq struct {
	IsAsync bool            `json:"is_async"`
	MsgList []PushSingleReq `json:"msg_list"`
}

func (c *Client) PushSingleBatchByCid(requestId []string, cid []string, notification *model.Notification) error {
	token := c.getToken(c.appId)
	resp := &BaseResp{}

	if len(requestId) != len(cid) {
		return fmt.Errorf("parrams error")
	}

	pageSize := 200
	requestIdGroup := SliceSplit[string](pageSize, requestId)
	cidGroup := SliceSplit(pageSize, cid)

	for index, ids := range requestIdGroup {
		msgList := make([]PushSingleReq, 0, pageSize)
		for i, id := range ids {
			msgList = append(msgList, PushSingleReq{
				RequestId: id,
				Audience: &model.Audience{
					Cid: []string{cidGroup[index][i]},
				},
				PushMessage: &model.PushMessage{
					Notification: notification,
				},
			})
		}
		err := PostHeader(c.getUrl(PATH_PUSH_SINGLE_BATCH_CID), &PushBatchReq{
			IsAsync: false,
			MsgList: msgList,
		}, resp, NewHeader().Add("token", token))
		if err != nil {
			log.Printf("request err:%v", err)
		}
	}
	return nil
}

type PushBatchCreateMsgReq struct {
	RequestId   string             `json:"request_id,omitempty"`   //请求唯一标识号，10-32位之间；如果request_id重复，会导致消息丢失
	GroupName   string             `json:"group_name,omitempty"`   //任务组名。多个消息任务可以用同一个任务组名，后续可根据任务组名查询推送情况（长度限制100字符，且不能含有特殊符号）
	Settings    *model.Settings    `json:"settings,omitempty"`     //推送条件设置
	PushMessage *model.PushMessage `json:"push_message,omitempty"` //个推推送消息参数
	PushChannel *model.PushChannel `json:"push_channel,omitempty"`
}

type PushBatchCreateMsgResp struct {
	BaseResp
	Data *struct {
		Taskid string `json:"taskid"`
	} `json:"data"`
}

func (c *Client) PushBatchCreateMsg(requestId, groupName string, notification *model.Notification) (string, error) {
	token := c.getToken(c.appId)
	resp := &PushBatchCreateMsgResp{}

	err := PostHeader(c.getUrl(PATH_PUSH_BATCH_CREATE_MSG), &PushBatchCreateMsgReq{
		RequestId: requestId,
		GroupName: groupName,
		PushMessage: &model.PushMessage{
			Notification: notification,
		},
		PushChannel: &model.PushChannel{
			Ios: nil,
			Android: &model.Android{Ups: &model.Ups{
				Notification: notification,
			}},
		},
	}, resp, NewHeader().Add("token", token))

	if err != nil {
		return "", err
	}
	if resp.Code != CODE_SUCCESS {
		return "", fmt.Errorf("%v", resp.Msg)
	}
	return resp.Data.Taskid, nil
}

func (c *Client) PushBatchByCid(taskId string, cid []string) error {
	token := c.getToken(c.appId)
	req := &struct {
		Audience *model.Audience `json:"audience"`
		Taskid   string          `json:"taskid"`
		IsAsync  bool            `json:"is_async"`
	}{
		Audience: &model.Audience{Cid: cid},
		Taskid:   taskId,
		IsAsync:  true,
	}
	resp := &BaseResp{}

	err := PostHeader(c.getUrl(PATH_PUSH_BATCH_BY_CID), req, resp, NewHeader().Add("token", token))

	if err != nil {
		return err
	}
	if resp.Code != CODE_SUCCESS {
		return fmt.Errorf("%v", resp.Msg)
	}
	return nil
}

func (c *Client) PushBatchByAlias(taskId string, alias []string) error {
	token := c.getToken(c.appId)
	req := &struct {
		Audience *model.Audience `json:"audience"`
		Taskid   string          `json:"taskid"`
		IsAsync  bool            `json:"is_async"`
	}{
		Audience: &model.Audience{Alias: alias},
		Taskid:   taskId,
		IsAsync:  true,
	}
	resp := &BaseResp{}

	err := PostHeader(c.getUrl(PATH_PUSH_BATCH_BY_ALIAS), req, resp, NewHeader().Add("token", token))

	if err != nil {
		return err
	}
	if resp.Code != CODE_SUCCESS {
		return fmt.Errorf("%v", resp.Msg)
	}
	return nil
}

func (c *Client) PushCreateMsgAndBatchByAlias(requestId, groupName string, notification *model.Notification, alias []string) (string, error) {
	taskId, err := c.PushBatchCreateMsg(requestId, groupName, notification)
	if err != nil {
		return "", err
	}
	return taskId, c.PushBatchByAlias(taskId, alias)
}

func (c *Client) PushCreateMsgAndBatchByCid(requestId, groupName string, notification *model.Notification, cid []string) (string, error) {
	taskId, err := c.PushBatchCreateMsg(requestId, groupName, notification)
	if err != nil {
		return "", err
	}
	return taskId, c.PushBatchByCid(taskId, cid)
}

func (c *Client) pushBatchCreateTransmission(requestId, groupName string, transmission string) (string, error) {
	token := c.getToken(c.appId)
	resp := &PushBatchCreateMsgResp{}

	err := PostHeader(c.getUrl(PATH_PUSH_BATCH_CREATE_MSG), &PushBatchCreateMsgReq{
		RequestId: requestId,
		GroupName: groupName,
		PushMessage: &model.PushMessage{
			Transmission: transmission,
		},
		PushChannel: &model.PushChannel{
			Ios: nil,
			Android: &model.Android{Ups: &model.Ups{
				Transmission: transmission,
			}},
		},
	}, resp, NewHeader().Add("token", token))

	if err != nil {
		return "", err
	}
	if resp.Code != CODE_SUCCESS {
		return "", fmt.Errorf("%v", resp.Msg)
	}
	return resp.Data.Taskid, nil
}

func (c *Client) PushCreateTransmissionAndBatchByAlias(requestId, groupName string, transmission string, alias []string) (string, error) {
	taskId, err := c.pushBatchCreateTransmission(requestId, groupName, transmission)
	if err != nil {
		return "", err
	}
	return taskId, c.PushBatchByAlias(taskId, alias)
}

func (c *Client) PushCreateTransmissionAndBatchByCid(requestId, groupName string, transmission string, cid []string) (string, error) {
	taskId, err := c.pushBatchCreateTransmission(requestId, groupName, transmission)
	if err != nil {
		return "", err
	}
	return taskId, c.PushBatchByCid(taskId, cid)
}
