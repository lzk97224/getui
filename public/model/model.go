package model

type Audience struct {
	Cid           []string `json:"cid,omitempty"`
	Alias         []string `json:"alias,omitempty"`
	Tag           []string `json:"tag,omitempty"`
	FastCustomTag string   `json:"fast_custom_tag,omitempty"`
	All           string   `json:"all,omitempty"`
}
type Settings struct {
	Ttl          int       `json:"ttl,omitempty"`      //消息离线时间设置，单位毫秒，-1表示不设离线，-1 ～ 3 * 24 * 3600 * 1000(3天)之间
	Strategy     *Strategy `json:"strategy,omitempty"` //厂商通道策略 https://docs.getui.com/getui/server/rest_v2/common_args/?id=doc-title-5
	Speed        int       `json:"speed,omitempty"`
	ScheduleTime int64     `json:"schedule_time,omitempty"`
}
type Strategy struct {
	Default int `json:"default,omitempty"`
}
type PushMessage struct {
	Duration     string        `json:"duration,omitempty"`     //手机端通知展示时间段，格式为毫秒时间戳段，两个时间的时间差必须大于10分钟，例如："1590547347000-1590633747000"
	Transmission string        `json:"transmission,omitempty"` //纯透传消息内容，安卓和iOS均支持，与notification、revoke 三选一，都填写时报错，长度 ≤ 3072
	Notification *Notification `json:"notification,omitempty"` //通知消息内容，仅支持安卓系统，iOS系统不展示个推通知消息，与transmission、revoke三选一，都填写时报错
	Revoke       *Revoke       `json:"revoke,omitempty"`
}
type Revoke struct {
	OldTaskId string `json:"old_task_id,omitempty"` //	需要撤回的taskId
	Force     bool   `json:"force,omitempty"`       //在没有找到对应的taskId，是否把对应appId下所有的通知都撤回
}
type Notification struct {
	Title        string `json:"title,omitempty"`         //通知消息标题，长度 ≤ 50
	Body         string `json:"body,omitempty"`          //通知消息内容，长度 ≤ 256
	BigText      string `json:"big_text,omitempty"`      //长文本消息内容，通知消息+长文本样式，与big_image二选一，两个都填写时报错，长度 ≤ 512
	BigImage     string `json:"big_image,omitempty"`     //大图的URL地址，通知消息+大图样式， 与big_text二选一，两个都填写时报错，长度 ≤ 1024
	Logo         string `json:"logo,omitempty"`          //通知的图标名称，包含后缀名（需要在客户端开发时嵌入），如“push.png”，长度 ≤ 64
	LogoUrl      string `json:"logo_url,omitempty"`      //通知图标URL地址，长度 ≤ 256
	ChannelId    string `json:"channel_id,omitempty"`    //通知渠道id，长度 ≤ 64
	ChannelName  string `json:"channel_name,omitempty"`  //通知渠道名称，长度 ≤ 64
	ChannelLevel int    `json:"channel_level,omitempty"` //设置通知渠道重要性（可以控制响铃，震动，浮动，闪灯等等）
	ClickType    string `json:"click_type,omitempty"`    //点击通知后续动作，
	Intent       string `json:"intent,omitempty"`        //点击通知打开应用特定页面，长度 ≤ 4096; 示例：intent://com.getui.push/detail?#Intent;scheme=gtpushscheme;launchFlags=0x4000000; package=com.getui.demo;component=com.getui.demo/com.getui.demo.DemoActivity;S.payload=payloadStr;end
	Url          string `json:"url,omitempty"`           //点击通知打开链接，长度 ≤ 1024
	Payload      string `json:"payload,omitempty"`       //点击通知时，附加自定义透传消息，长度 ≤ 3072
	NotifyId     int    `json:"notify_id,omitempty"`     //覆盖任务时会使用到该字段，两条消息的notify_id相同，新的消息会覆盖老的消息，范围：0-2147483647
	RingName     string `json:"ring_name,omitempty"`     //自定义铃声，请填写文件名，不包含后缀名(需要在客户端开发时嵌入)，个推通道下发有效 客户端SDK最低要求 2.14.0.0
	BadgeAddNum  int    `json:"badge_add_num,omitempty"` //角标, 必须大于0, 个推通道下发有效
	ThreadId     string `json:"thread_id,omitempty"`     //消息折叠分组，设置成相同thread_id的消息会被折叠（仅支持个推渠道下发的安卓消息）。目前与iOS的thread_id设置无关，安卓和iOS需要分别设置。
}
type PushChannel struct {
	Ios     *Ios     `json:"ios,omitempty"`
	Android *Android `json:"android,omitempty"`
}
type Android struct {
	Ups *Ups `json:"ups,omitempty"`
}
type Ups struct {
	Transmission string        `json:"transmission,omitempty"` //纯透传消息内容，安卓和iOS均支持，与notification、revoke 三选一，都填写时报错，长度 ≤ 3072
	Notification *Notification `json:"notification,omitempty"` //通知消息内容，仅支持安卓系统，iOS系统不展示个推通知消息，与transmission、revoke三选一，都填写时报错
	Revoke       *Revoke       `json:"revoke,omitempty"`
}
type Ios struct {
	Type           string       `json:"type,omitempty"`
	Aps            *Aps         `json:"aps,omitempty"`
	AutoBadge      string       `json:"auto_badge,omitempty"`
	Payload        string       `json:"payload,omitempty"`
	Multimedia     []Multimedia `json:"multimedia,omitempty"`
	ApnsCollapseId string       `json:"apns-collapse-id,omitempty"`
}
type Multimedia struct {
	Url      string `json:"url,omitempty"`
	Type     int    `json:"type,omitempty"`
	OnlyWifi bool   `json:"only_wifi,omitempty"`
}
type Aps struct {
	Alert            *Alert `json:"alert,omitempty"`
	ContentAvailable int    `json:"content-available,omitempty"`
	Sound            string `json:"sound,omitempty"`
	Category         string `json:"category,omitempty"`
	ThreadId         string `json:"thread-id,omitempty"`
}

type Alert struct {
	Title           string   `json:"title,omitempty"`
	Body            string   `json:"body,omitempty"`
	ActionLocKey    string   `json:"action_loc_key,omitempty"`
	LocKey          string   `json:"loc_key,omitempty"`
	LocArgs         []string `json:"loc_args,omitempty"`
	LaunchImage     string   `json:"launch_image,omitempty"`
	TitleLocKey     string   `json:"title_loc_key,omitempty"`
	TitleLocArgs    []string `json:"title_loc_args,omitempty"`
	Subtitle        string   `json:"subtitle,omitempty"`
	SubtitleLocKey  string   `json:"subtitle_loc_key,omitempty"`
	SubtitleLocArgs []string `json:"subtitle_loc_args,omitempty"`
}

type UserBindAlias struct {
	Cid   string `json:"cid"`
	Alias string `json:"alias"`
}
