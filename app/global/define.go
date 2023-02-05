package global

import (
	"errors"
)

// 状态码
var (
	StatusCodeOk   int32 = 0 // 响应状态码 - 成功
	StatusCodeFail int32 = 1 // 响应状态码 - 一般失败
)

// 错误大全
var (
	// 数据库方面错误
	ErrDbCreateUniqueKeyRepeatedly error = errors.New("ErrDbCreateUniqueKeyRepeatedly") // 重复创建按了应该唯一的key的一条记录
)

// 点赞和取消点赞
var (
	CancelFavoriteAction        = 0 //取消点赞
	FavoriteAction              = 1 //点赞
	RequestCancelFavoriteAction = 2
)

var (
	MessageSendEvent = 1 //发送消息
)
