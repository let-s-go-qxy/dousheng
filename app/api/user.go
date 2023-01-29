package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"gorm.io/gorm"
	"strconv"
	g "tiktok/app/global"
	"tiktok/app/internal/model"
	"tiktok/app/internal/service"
	utils2 "tiktok/utils"
	"time"
)

type UserLoginResponse struct {
	Response
	UserId int    `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

// UserInfo 获取用户详情
func UserInfo(c context.Context, ctx *app.RequestContext) {
	user := new(model.User)
	user.Id, _ = strconv.Atoi(ctx.Query("user_id"))
	err := g.MysqlDB.First(&user).Error
	if err != nil {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  "查询失败: " + err.Error(),
		})
	}
	ctx.JSON(consts.StatusOK, UserResponse{
		Response: Response{
			StatusCode: g.StatusCodeOk,
		},
		User: User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   0, // TODO 等待数据库设计稳定后
			FollowerCount: 0,
			IsFollow:      false,
		},
	})
}

// UserRegister 用户注册
func UserRegister(c context.Context, ctx *app.RequestContext) {
	name := ctx.Query("username")
	pw := ctx.Query("password")
	// TODO 看一下password在app那边是不是限制6位以上
	if len(name) > 32 || len(pw) > 32 || name == "" || len(pw) < 6 {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  "账号或密码不符合要求",
		})
		return
	}
	salt := strconv.Itoa(int(time.Now().UnixNano()))
	pw = utils2.GetMd5Str(pw + salt)
	// 填装数据
	user := &model.User{
		Name:     name,
		Password: pw,
		Salt:     salt,
	}
	_, err := model.CreateUser(user)
	// 创建失败
	if err != nil {
		if err == g.ErrDbCreateUniqueKeyRepeatedly {
			ctx.JSON(consts.StatusOK, Response{
				StatusCode: g.StatusCodeFail,
				StatusMsg:  "User already exist",
			})
			return
		}
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  "创建用户失败: " + err.Error(),
		})
		return
	}
	ctx.JSON(consts.StatusOK, UserLoginResponse{
		Response: Response{
			StatusCode: g.StatusCodeOk,
		},
		UserId: user.Id,
		Token:  service.GenerateToken(*user),
	})
}

// UserLogin 用户登录
func UserLogin(c context.Context, ctx *app.RequestContext) {
	name := ctx.Query("username")
	pw := ctx.Query("password")
	if len(name) > 32 || len(pw) > 32 || name == "" || len(pw) < 6 {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  "账号或密码不符合要求",
		})
		return
	}
	user, err := model.GetUserByName(name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(consts.StatusOK, Response{
				StatusCode: g.StatusCodeFail,
				StatusMsg:  "账号或密码错误",
			})
			return
		}
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  "查询失败: " + err.Error(),
		})
		return
	}
	if utils2.GetMd5Str(pw+user.Salt) != user.Password {
		ctx.JSON(consts.StatusOK, Response{
			StatusCode: g.StatusCodeFail,
			StatusMsg:  "账号或密码错误",
		})
		return
	}
	ctx.JSON(consts.StatusOK, UserLoginResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserId: user.Id,
		Token:  service.GenerateToken(*user),
	})
}
