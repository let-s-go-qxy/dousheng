package comment

import (
	"errors"
	"github.com/jinzhu/copier"
	repository "tiktok/app/internal/model"
	"tiktok/app/internal/service/user"
	"time"
)

type User struct {
	Id            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
type Comment struct {
	Id         int    `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_date,omitempty"`
}

// GetCommentList查选该视频下的所有评论
func GetCommentList(videoId int) (comments []Comment, vidoeCommentCount int) {

	// 调用model层comment的sql查询语句，根据视频id查询对应id的视频评论
	commentsWithUserid := repository.FindCommentByVideo(videoId)
	for _, commentWithUserid := range commentsWithUserid {
		commentWithUser := Comment{}
		userid, followCount, followerCount, name, isFollow, err := user.UserInfo(commentWithUserid.UserId, commentWithUserid.UserId)
		if err != nil {
			err = errors.New("发表用户不存在: " + err.Error())
		}
		userDao := User{
			Id:            userid,
			Name:          name,
			FollowCount:   followCount,
			FollowerCount: followerCount,
			IsFollow:      isFollow,
		}
		copier.Copy(&commentWithUser, &commentWithUserid)
		commentWithUser.User = userDao
		comments = append(comments, commentWithUser)
	}
	// 得到视频下的评论数
	vidoeCommentCount = len(comments)
	return
}

// 对评论进行创建或者删除
func CommentAction(videoId int, actionType int, content string, commentId int, userId int) (comment repository.Comment, userDao User, err error) {
	// 调用service.userInfo方法查询发表用户信息
	userid, followCount, followerCount, name, isFollow, err := user.UserInfo(userId, userId)
	if err != nil {
		err = errors.New("创建者不存在: " + err.Error())
	}
	userDao = User{
		Id:            userid,
		Name:          name,
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
	}
	// 填装Comment数据
	com := &repository.Comment{
		Id:         commentId,
		UserId:     userId,
		Content:    content,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		VideoId:    videoId,
		Cancel:     actionType,
	}
	if actionType == 1 {
		com, err = repository.CreateComment(com)
		// 评论创建失败
		if err != nil {
			err = errors.New("发表评论失败: " + err.Error())

		}
		comment = repository.FindCommentById(com.Id)
	} else {
		err = repository.DeleteComment(com)
		if err != nil {
			err = errors.New("删除评论失败: " + err.Error())
		}
		comment = repository.FindCommentById(commentId)
	}
	return
}
