package api

import (
	"AIMusic-API/global"
	"AIMusic-API/proto"
	"AIMusic-API/utils"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"strconv"
	"time"
)

type Community struct{}

func (c Community) Bulletin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": global.RDB.Get("bulletin").Val(),
		"msg":  nil,
	})
}

func (c Community) CreationTop(ctx *gin.Context) {
	//rsp, err := global.CommunityClient.GetTop(context.Background(), &emptypb.Empty{})
	rsp, err := global.CommunityClients[global.LoadBalancing(len(global.CommunityInstances))].GetTop(context.Background(), &emptypb.Empty{})
	if err != nil {
		zap.S().Errorw("[CreationTop]获取热门作品失败")
		utils.HandleGrpcError2Http(err, ctx)
		return
	}
	type creation struct {
		Id              int64  `json:"id"`
		UserName        string `json:"userName"`
		UserAvatar      string `json:"userAvatar"`
		CompositionName string `json:"compositionName"`
		ForWhat         string `json:"forWhat"`
		Mp3             string `json:"mp3"`
		Likes           int64  `json:"likes"`
		Report          int64  `json:"report"`
		CreationTime    string `json:"creationTime"`
		Detail          string `json:"detail"`
	}
	var result []creation
	for _, v := range rsp.Creation {
		var t creation
		t.Id = v.Id
		t.Likes = v.Likes
		t.CreationTime = time.Unix(int64(v.CreatedTime), 0).Format("2006-01-02")
		t.Mp3 = v.Mp3
		t.CompositionName = v.CompositionName
		t.UserName = v.UserName
		t.UserAvatar = v.UserAvatar
		t.ForWhat = v.ForWhat
		t.Detail = v.Detail
		t.Report = v.Report
		result = append(result, t)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
		"msg":  nil,
	})
}

func (c Community) UserSearch(ctx *gin.Context) {
	rsp, err := global.CommunityClients[global.LoadBalancing(len(global.CommunityInstances))].SearchUserByName(context.Background(), &proto.SearchRequest{Name: ctx.Query("name")})
	if err != nil {
		zap.S().Errorw("[UserSearch]搜索用户失败")
		utils.HandleGrpcError2Http(err, ctx)
		return
	}
	type user struct {
		Nickname string `json:"nickname"`
		Gender   string `json:"gender"`
		Pic      string `json:"pic"`
		Birthday string `json:"birthday"`
		Avatar   string `json:"avatar"`
		Motto    string `json:"motto"`
	}
	var result []user
	for _, v := range rsp.Users {
		var t user
		t.Birthday = v.Birthday
		t.Nickname = v.Nickname
		t.Pic = v.Pic
		t.Gender = v.Gender
		t.Avatar = v.Pic
		t.Motto = v.Motto
		result = append(result, t)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
		"msg":  nil,
	})
}

func (c Community) CreationSearch(ctx *gin.Context) {
	rsp, err := global.CommunityClients[global.LoadBalancing(len(global.CommunityInstances))].SearchCreationByName(context.Background(), &proto.SearchRequest{Name: ctx.Query("name")})
	if err != nil {
		zap.S().Errorw("[CreationSearch]搜索作品失败")
		utils.HandleGrpcError2Http(err, ctx)
		return
	}
	type creation struct {
		Id              int64  `json:"id"`
		Nickname        string `json:"nickname"`
		CompositionName string `json:"compositionName"`
		Detail          string `json:"detail"`
		ForWhat         string `json:"forWhat"`
		Likes           int64  `json:"likes"`
		Report          int64  `json:"report"`
		Mp3             string `json:"mp3"`
		Avatar          string `json:"avatar"`
	}
	var result []creation
	for _, v := range rsp.Creations {
		var t creation
		t.Id = v.Id
		t.Detail = v.Detail
		t.Nickname = v.Nickname
		t.CompositionName = v.CompositionName
		t.ForWhat = v.ForWhat
		t.Likes = v.Likes
		t.Report = v.Report
		t.Mp3 = v.Mp3
		t.Avatar = v.Avatar
		result = append(result, t)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
		"msg":  nil,
	})
}

func (c Community) Attitude(ctx *gin.Context) {
	opt := ctx.PostForm("opt")
	targetStr := ctx.PostForm("target")
	target, _ := strconv.Atoi(targetStr)
	if opt == "0" {
		global.CommunityClients[global.LoadBalancing(len(global.CommunityInstances))].CreationAttitude(context.Background(), &proto.AttitudeCreation{
			Id:  int64(target),
			Opt: 0,
		})
	} else if opt == "1" {
		global.CommunityClients[global.LoadBalancing(len(global.CommunityInstances))].CreationAttitude(context.Background(), &proto.AttitudeCreation{
			Id:  int64(target),
			Opt: 1,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "操作完成",
		"msg":  nil,
	})
}
