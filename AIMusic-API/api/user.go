package api

import (
	"AIMusic-API/global"
	"AIMusic-API/proto"
	"AIMusic-API/utils"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type User struct{}

func (u User) GetInfo(ctx *gin.Context) {
	//暂时没有登陆 直接传手机号
	mobile := ctx.Query("mobile")
	rsp, err := global.UserClients[global.LoadBalancing(len(global.UserInstances))].GetInfo(context.Background(), &proto.GetRequest{Mobile: mobile})
	if err != nil {
		zap.S().Errorw("[GetInfo]获取用户信息失败")
		utils.HandleGrpcError2Http(err, ctx)
		return
	}
	type ginResponse struct {
		Nickname string `json:"nickname,omitempty"`
		Avatar   string `json:"avatar,omitempty"`
		Motto    string `json:"motto,omitempty"`
		Gender   string `json:"gender,omitempty"`
		Birthday string `json:"birthday,omitempty"`
		Likes    int64  `json:"likes,omitempty"`
	}
	var result ginResponse
	result.Likes = rsp.Likes
	result.Avatar = rsp.Avatar
	result.Gender = rsp.Gender
	result.Birthday = rsp.Birthday
	result.Motto = rsp.Motto
	result.Nickname = rsp.Nickname
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
		"msg":  nil,
	})
	zap.S().Debug("[Mobile:%d]获取用户信息成功", mobile)
}

func (u User) GetCreation(ctx *gin.Context) {
	//暂时没有登陆 直接传手机号
	mobile := ctx.Query("mobile")
	rsp, err := global.UserClients[global.LoadBalancing(len(global.UserInstances))].GetCreation(context.Background(), &proto.GetRequest{Mobile: mobile})
	if err != nil {
		zap.S().Errorw("[GetCreation]获取用户作品失败")
		utils.HandleGrpcError2Http(err, ctx)
		return
	}
	type Finalize struct {
		Name         string `json:"name"`
		ForWhat      string `json:"forWhat"`
		CreationTime string `json:"creationTime"`
		Status       string `json:"status"`
		Detail       string `json:"detail"`
		Mp3          string `json:"mp3"`
	}
	type Draft struct {
		Name         string `json:"name"`
		ForWhat      string `json:"forWhat"`
		CreationTime string `json:"creationTime"`
		LastModified string `json:"lastModified"`
		Detail       string `json:"detail"`
		Mp3          string `json:"mp3"`
	}
	type ginResponse struct {
		Finalize []Finalize `json:"finalize"`
		Draft    []Draft    `json:"draft"`
	}
	var draft []Draft
	var finalize []Finalize
	var result ginResponse
	for _, v := range rsp.Creation {
		if v.Status == 0 {
			tmp := Draft{
				Name:         v.Name,
				ForWhat:      v.ForWhat,
				CreationTime: v.CreationTime,
				LastModified: v.LastModified,
				Detail:       v.Detail,
				Mp3:          v.Mp3,
			}
			draft = append(draft, tmp)
		} else if v.Status == 1 {
			tmp := Finalize{
				Name:         v.Name,
				ForWhat:      v.ForWhat,
				CreationTime: v.CreationTime,
				Status:       "公开",
				Detail:       v.Detail,
				Mp3:          v.Mp3,
			}
			finalize = append(finalize, tmp)
		} else if v.Status == 2 {
			tmp := Finalize{
				Name:         v.Name,
				ForWhat:      v.ForWhat,
				CreationTime: v.CreationTime,
				Status:       "私密",
				Detail:       v.Detail,
				Mp3:          v.Mp3,
			}
			finalize = append(finalize, tmp)

		}
	}
	result.Draft = draft
	result.Finalize = finalize
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
		"msg":  nil,
	})
	zap.S().Debug("[Mobile:%d]获取用户作品成功", mobile)
}

func (u User) Feedback(ctx *gin.Context) {
	global.UserClients[global.LoadBalancing(len(global.UserInstances))].Feedback(context.Background(), &proto.FeedbackRequest{
		Mobile:   ctx.PostForm("mobile"),
		Suspect:  ctx.PostForm("target"),
		Evidence: ctx.PostForm("reason"),
		Advice:   ctx.PostForm("advice"),
		Avatar:   ctx.PostForm("avatar"),
		Nickname: ctx.PostForm("nickname"),
	})
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "操作完成",
		"msg":  nil,
	})
}
