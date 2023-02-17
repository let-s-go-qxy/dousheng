package boot

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sync"
	g "tiktok/app/global"
	"tiktok/manifest/ossRelated"
)

var bucket *oss.Bucket
var once sync.Once

// OSSInit 初始化
func OSSInit() {
	once.Do(func() {
		// 连接OSS账户
		client, err := oss.New(ossRelated.EndPoint, ossRelated.AccessKeyID, ossRelated.AccessKeySecret)

		if err != nil {
			g.Logger.Infof("连接OSS账户失败" + err.Error())
		} else { // OSS账户连接成功

			// 连接存储空间
			bucket, err = client.Bucket(ossRelated.BucketName)
			if err != nil {
				g.Logger.Infof("连接存储空间失败" + err.Error())
			} else { // 存储空间连接成功
				g.Logger.Infof("OSS初始化完成")
			}
		}
	})
	g.OssBucket = bucket

}
