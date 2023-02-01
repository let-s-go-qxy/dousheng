package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	g "tiktok/app/global"
	repository "tiktok/app/internal/model"
	"tiktok/manifest/ossRelated"
	utils "tiktok/utils/file"
)

func PublishVideo(userID int, title string, fileBytes []byte) (success bool) {

	//雪花算法生成fileID
	node, _ := utils.NewWorker(1)
	randomId := node.GetId()
	fileID := fmt.Sprintf("%v", randomId)
	success = true

	if !utils.UploadFile(fileBytes, fileID, "video") {
		success = false
	}

	// 通过ffmpeg截取视频第一帧为视频封面
	//videoURL := common.OSSPreURL + fileID + ".mp4"
	videoName := fileID + ".mp4"
	pictureName := fileID + ".jpg"

	//封面图和视频在本地的保存路径
	picturePath := ossRelated.LocalFolderPath + pictureName
	videoPath := ossRelated.LocalFolderPath + videoName

	//将上传的文件流的形式以mp4的形式保存到本地，并将视频的第一帧作为封面图导出到，picturePath下
	ioutil.WriteFile(videoPath, fileBytes, 0666)
	cmd := exec.Command(ossRelated.FfmpegPath, "-i", videoPath, "-y", "-f", "image2", "-ss", "1", picturePath)
	//buf := new(bytes.Buffer)
	//cmd.Stdout = buf
	cmd.Run()

	//以字节数组的形式读出本地的封面图，便于后续上传到云端
	openFile, err := os.Open(picturePath)
	defer openFile.Close()
	if err != nil {
		g.Logger.Infof("打开picture文件时发生了错误")
	}
	pictureBytes, err := ioutil.ReadAll(openFile)
	if err != nil {
		g.Logger.Infof("读取picture文件时发生了错误")
	}

	// 将视频封面上传至OSS中
	if !utils.UploadFile(pictureBytes, fileID, "picture") {
		success = false
	}
	if success {
		if repository.VideoDao.PublishVideo(userID, title, fileID) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// GetVideoListByIdList 根据视频ID列表查询视频列表,按照点赞时间顺序
func GetVideoListByIdList(videoIdList []int) (videoList []repository.Video) {
	for _, videoId := range videoIdList {
		video := repository.Video{}
		g.MysqlDB.Table("videos").Where("id = ?", videoId).Take(&video)
		videoList = append(videoList, video)
	}
	return
}
