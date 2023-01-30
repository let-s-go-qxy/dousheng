package global

import "errors"

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

//取消点赞和未取消
var (
	Cancel   = 0 //未取消
	NoCancel = 1 // 取消

)
