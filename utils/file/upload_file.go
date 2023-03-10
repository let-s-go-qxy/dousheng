package file

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	g "tiktok/app/global"
)

func UploadFile(file []byte, filename string, fileType string) bool {
	var fileSuffix string
	if fileType == "video" {
		fileSuffix = ".mp4"
	} else if fileType == "picture" {
		fileSuffix = ".jpg"
	} else {
		g.Logger.Infof("无法上传" + fileType + "类型文件")
		return false
	}
	//初始化bucket

	err := g.OssBucket.PutObject("video/"+filename+fileSuffix, bytes.NewReader(file))
	if err != nil {
		g.Logger.Infof("上传文件失败" + err.Error())
		return false
	} else {
		return true
	}
}

func UploadAvatar(userID int) bool {

	var fileSuffix string
	fileSuffix = ".jpg"
	strUserID := strconv.Itoa(userID)

	ModUserID := userID % 10
	strModUserID := strconv.Itoa(ModUserID)
	//项目中头像相对路径
	avatarPath := "imgs/" +
		strModUserID +
		"_avatar.jpg"

	//以字节数组的形式读出本地的头像，便于后续上传到云端
	openFile, err := os.Open(avatarPath)
	defer openFile.Close()
	avatarBytes, err := ioutil.ReadAll(openFile)
	if err != nil {
		g.Logger.Infof("上传头像失败" + err.Error())
		return false
	}
	filename := strUserID + "_avatar"
	err = g.OssBucket.PutObject("avatar/"+filename+fileSuffix, bytes.NewReader(avatarBytes))
	if err != nil {
		g.Logger.Infof("上传头像失败" + err.Error())
		return false
	} else {
		return true
	}
}

//func DownloadFile(downloadFileName, filename string, fileType string) bool {
//	bucket.GetObject("video/"+)
//}
