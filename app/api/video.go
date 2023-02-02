package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	g "tiktok/app/global"
	"tiktok/app/internal/service/video"
	"tiktok/utils/common"
	"tiktok/utils/msg"
)

func PublishVideo(c context.Context, ctx *app.RequestContext) {
	title := ctx.PostForm("title")
	data, err := ctx.FormFile("data")
	userID, success := ctx.Get("user_id")
	if !success {
		ctx.JSON(http.StatusOK,
			common.Response{
				StatusCode: -1,
				StatusMsg:  msg.TokenParameterAcquisitionError,
			})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK,
			common.Response{
				StatusCode: -1,
				StatusMsg:  msg.PublishVideoFailedMsg,
			})
		return
	}

	fileHandle, err1 := data.Open() //打开上传文件
	if err1 != nil {
		g.Logger.Infof("打开文件失败" + err1.Error())
	}

	// 闭包处理错误
	defer func(fileHandle multipart.File) {
		err := fileHandle.Close()
		if err != nil {
			g.Logger.Infof("关闭文件错误" + err.Error())
		}
	}(fileHandle)

	fileByte, err2 := ioutil.ReadAll(fileHandle)
	if err2 != nil {
		g.Logger.Infof("读取文件错误" + err2.Error())
	}

	if video.PublishVideo(userID.(int), title, fileByte) {
		ctx.JSON(http.StatusOK,
			common.Response{
				StatusCode: 0,
				StatusMsg:  msg.PublishVideoSuccessMsg,
			})
	} else {
		ctx.JSON(http.StatusOK,
			common.Response{
				StatusCode: -1,
				StatusMsg:  msg.PublishVideoFailedMsg,
			})
	}
}
