package getui

const (
	ClickTypeIntent        = "intent"         //打开应用内特定页面，
	ClickTypeUrl           = "url"            //打开网页地址，
	ClickTypePayload       = "payload"        //自定义消息内容启动应用
	ClickTypePayloadCustom = "payload_custom" //自定义消息内容不启动应用
	ClickTypeStartapp      = "startapp"       //打开应用首页，
	ClickTypeNone          = "none"           //纯通知，无后续动作
)
