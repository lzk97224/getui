package internal

const (
	BASE_URL = "https://restapi.getui.com/v2"
)

const (
	PATH_AUTH                    = "/auth"                    //授权接口
	PATH_PUSH_SINGLE_CID         = "/push/single/cid"         //向单个用户推送消息，可根据cid指定用户
	PATH_PUSH_SINGLE_BATCH_CID   = "/push/single/batch/cid"   //批量发送单推消息，每个cid用户的推送内容都不同的情况下，使用此接口，可提升推送效率。
	PATH_PUSH_SINGLE_BATCH_ALIAS = "/push/single/batch/alias" //批量发送单推消息，在给每个别名用户的推送内容都不同的情况下，可以使用此接口
	PATH_PUSH_BATCH_CREATE_MSG   = "/push/list/message"       //此接口用来创建消息体，并返回taskid，为批量推的前置步骤
	PATH_PUSH_BATCH_BY_CID       = "/push/list/cid"           //对列表中所有cid进行消息推送
	PATH_PUSH_BATCH_BY_ALIAS     = "/push/list/alias"         //对列表中所有别名进行消息推送。调
	PATH_USER_ALIAS              = "/user/alias"              //绑定别名 ;一个cid只能绑定一个别名，若已绑定过别名的cid再次绑定新别名，则前一个别名会自动解绑，并绑定新别名。
	PATH_USER_ALIAS_CID          = "/user/alias/cid"          //通过传入的cid查询对应的别名信息
	PATH_USER_CID_ALIAS          = "/user/cid/alias"          //通过传入的别名查询对应的cid信息
	PATH_BATCH_UNBIND_ALIAS      = "/user/alias"              //批量解除别名与cid的关系
	PATH_UNBIND_BY_ALIAS         = "/user/alias"              //解绑所有与该别名绑定的cid
	PATH_BIND_TAG_USER_TO_TAGS   = "/user/custom_tag/cid"     //一个用户绑定一批标签
	PATH_BIND_TAG_USERS_TO_TAG   = "/user/custom_tag/batch"   //一批用户绑定一个标签
	PATH_UNBIND_TAG_USERS_OF_TAG = "/user/custom_tag/batch"   //一批用户解绑一个标签
	PATH_QUERY_USER_TAG_OF_TAG   = "/user/custom_tag/cid"     //根据cid查询用户标签列表
	PATH_ADD_BLACK               = "/user/black/cid"          //根据cid查询用户标签列表
)

const (
	CODE_SUCCESS = 0
)
