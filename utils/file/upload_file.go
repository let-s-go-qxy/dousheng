package file

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sync"
	g "tiktok/app/global"
	"tiktok/manifest/ossRelated"
)

var bucket *oss.Bucket
var once sync.Once

// OSSInit 初始化，将ConnQuery与数据库绑定
func OSSInit() {
	once.Do(func() {
		// 连接OSS账户
		client, err := oss.New(ossRelated.EndPoint, ossRelated.AccessKeyID, ossRelated.AccessKeySecret)

		if err != nil {
			g.Logger.Infof("连接OSS账户失败" + err.Error())
		} else { // OSS账户连接成功

			// 连接存储空间
			bucket, err = client.Bucket("camp-dou-sheng")
			if err != nil {
				g.Logger.Infof("连接存储空间失败" + err.Error())
			} else { // 存储空间连接成功
				g.Logger.Infof("OSS初始化完成")
			}
		}
	})
}

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
	err := bucket.PutObject("video/"+filename+fileSuffix, bytes.NewReader(file))
	if err != nil {
		g.Logger.Infof("上传文件失败" + err.Error())
		return false
	} else {
		return true
	}
}

//func DownloadFile(downloadFileName, filename string, fileType string) bool {
//	bucket.GetObject("video/"+)
//}
