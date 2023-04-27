package api

import (
	"AIMusic-API/global"
	"AIMusic-API/proto"
	"AIMusic-API/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Administrator struct{}

func (a Administrator) Usage(c *gin.Context) {
	tracerStr, _ := c.Get("tracer")
	parentSpanStr, _ := c.Get("parentSpan")
	tracer, _ := tracerStr.(opentracing.Tracer)
	parentSpan, _ := parentSpanStr.(opentracing.Span)

	funcName := strings.Split(c.HandlerName(), ".")[1]
	grpcSpan := tracer.StartSpan(
		funcName+"-GRPC",
		opentracing.ChildOf(parentSpan.Context()),
	)
	rsp, err := global.AdministratorClients[global.LoadBalancing(len(global.AdministratorInstances))].GetUsage(context.Background(), &emptypb.Empty{})
	grpcSpan.Finish()

	if err != nil {
		zap.S().Errorw("[Usage]获取使用量失败")
		utils.HandleGrpcError2Http(err, c)
		return
	}

	type node struct {
		Amount int64  `json:"amount"`
		Date   string `json:"date"`
	}
	type target struct {
		Amount int64  `json:"amount"`
		Name   string `json:"name"`
	}
	type ginResponse struct {
		Scratch     []node   `json:"scratch"`
		Composition []node   `json:"composition"`
		ForWhat     []target `json:"forWhat"`
	}

	var result ginResponse
	var scratch []node
	var composition []node
	var forWhat []target

	for _, v := range rsp.ForWhat {
		var t target
		t.Amount = v.Amount
		t.Name = v.Name
		forWhat = append(forWhat, t)
	}
	for _, v := range rsp.Scratch {
		var t node
		t.Date = v.Date
		t.Amount = v.Amount
		scratch = append(scratch, t)
	}
	for _, v := range rsp.Composition {
		var t node
		t.Date = v.Date
		t.Amount = v.Amount
		composition = append(composition, t)
	}
	result.Composition = composition
	result.Scratch = scratch
	result.ForWhat = forWhat

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
		"msg":  nil,
	})
}

func (a Administrator) UserAudit(c *gin.Context) {
	tracerStr, _ := c.Get("tracer")
	parentSpanStr, _ := c.Get("parentSpan")
	tracer, _ := tracerStr.(opentracing.Tracer)
	parentSpan, _ := parentSpanStr.(opentracing.Span)

	funcName := strings.Split(c.HandlerName(), ".")[1]
	grpcSpan := tracer.StartSpan(
		funcName+"-GRPC",
		opentracing.ChildOf(parentSpan.Context()),
	)
	rsp, err := global.AdministratorClients[global.LoadBalancing(len(global.AdministratorInstances))].UserAudit(context.Background(), &emptypb.Empty{})
	grpcSpan.Finish()

	if err != nil {
		zap.S().Errorw("[UserAudit]获取用户审计信息失败")
		utils.HandleGrpcError2Http(err, c)
		return
	}
	type likes struct {
		Mobile   string `json:"mobile"`
		Nickname string `json:"nickname"`
		Likes    int64  `json:"likes"`
		Motto    string `json:"motto"`
		Avatar   string `json:"avatar"`
	}
	type report struct {
		Mobile   string `json:"mobile"`
		Nickname string `json:"nickname"`
		Report   int64  `json:"report"`
		Motto    string `json:"motto"`
		Avatar   string `json:"avatar"`
	}
	type ginResponse struct {
		Likes  []likes  `json:"likes"`
		Report []report `json:"report"`
	}
	var result ginResponse
	var userLikes []likes
	var userReport []report
	for _, v := range rsp.UserLikes {
		var t likes
		t.Nickname = v.Nickname
		t.Mobile = v.Mobile
		t.Likes = v.Likes
		t.Motto = v.Motto
		t.Avatar = v.Avatar
		userLikes = append(userLikes, t)
	}
	for _, v := range rsp.UserReport {
		var t report
		t.Nickname = v.Nickname
		t.Mobile = v.Mobile
		t.Report = v.Report
		t.Motto = v.Motto
		t.Avatar = v.Avatar
		userReport = append(userReport, t)
	}
	result.Likes = userLikes
	result.Report = userReport
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
		"msg":  nil,
	})
}

func (a Administrator) CreationAudit(c *gin.Context) {
	tracerStr, _ := c.Get("tracer")
	parentSpanStr, _ := c.Get("parentSpan")
	tracer, _ := tracerStr.(opentracing.Tracer)
	parentSpan, _ := parentSpanStr.(opentracing.Span)

	funcName := strings.Split(c.HandlerName(), ".")[1]
	grpcSpan := tracer.StartSpan(
		funcName+"-GRPC",
		opentracing.ChildOf(parentSpan.Context()),
	)
	rsp, err := global.AdministratorClients[global.LoadBalancing(len(global.AdministratorInstances))].CreationAudit(context.Background(), &emptypb.Empty{})
	grpcSpan.Finish()

	if err != nil {
		zap.S().Errorw("[CreationAudit]获取作品审计信息失败")
		utils.HandleGrpcError2Http(err, c)
		return
	}
	type likes struct {
		Id      int64  `json:"id"`
		Name    string `json:"name"`
		Likes   int64  `json:"likes"`
		Mp3     string `json:"mp3"`
		Detail  string `json:"detail"`
		Nick    string `json:"nick"`
		Avatar  string `json:"avatar"`
		Note    string `json:"note"`
		Emotion string `json:"emotion"`
	}
	type report struct {
		Id      int64  `json:"id"`
		Name    string `json:"name"`
		Report  int64  `json:"likes"`
		Mp3     string `json:"mp3"`
		Detail  string `json:"detail"`
		Nick    string `json:"nick"`
		Avatar  string `json:"avatar"`
		Note    string `json:"note"`
		Emotion string `json:"emotion"`
	}
	type ginResponse struct {
		Likes  []likes  `json:"likes"`
		Report []report `json:"report"`
	}
	var result ginResponse
	var creationLikes []likes
	var creationReport []report
	for _, v := range rsp.CreationLikes {
		var t likes
		t.Id = v.Id
		t.Name = v.Name
		t.Likes = v.Likes
		t.Mp3 = v.Mp3
		t.Detail = v.Detail
		t.Nick = v.Nick
		t.Avatar = v.Avatar
		creationLikes = append(creationLikes, t)
	}
	for _, v := range rsp.CreationReport {
		var t report
		t.Id = v.Id
		t.Name = v.Name
		t.Report = v.Report
		t.Mp3 = v.Mp3
		t.Detail = v.Detail
		t.Nick = v.Nick
		t.Avatar = v.Avatar
		creationReport = append(creationReport, t)
	}
	result.Likes = creationLikes
	result.Report = creationReport
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
		"msg":  nil,
	})
}

func (a Administrator) UserExecute(c *gin.Context) {
	tracerStr, _ := c.Get("tracer")
	parentSpanStr, _ := c.Get("parentSpan")
	tracer, _ := tracerStr.(opentracing.Tracer)
	parentSpan, _ := parentSpanStr.(opentracing.Span)

	funcName := strings.Split(c.HandlerName(), ".")[1]
	grpcSpan := tracer.StartSpan(
		funcName+"-GRPC",
		opentracing.ChildOf(parentSpan.Context()),
	)
	opt, _ := strconv.Atoi(c.PostForm("opt"))
	rsp, err := global.AdministratorClients[global.LoadBalancing(len(global.AdministratorInstances))].ExecuteUser(context.Background(), &proto.ExecutionUserRequest{
		Mobile: c.PostForm("mobile"),
		Opt:    int32(opt),
	})
	grpcSpan.Finish()

	if err != nil {
		zap.S().Errorw("[UserExecute]用户处理失败")
		utils.HandleGrpcError2Http(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": rsp.Msg,
		"msg":  nil,
	})
}

func (a Administrator) CreationExecute(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("opt"))
	opt, _ := strconv.Atoi(c.PostForm("opt"))

	tracerStr, _ := c.Get("tracer")
	parentSpanStr, _ := c.Get("parentSpan")
	tracer, _ := tracerStr.(opentracing.Tracer)
	parentSpan, _ := parentSpanStr.(opentracing.Span)

	funcName := strings.Split(c.HandlerName(), ".")[1]
	grpcSpan := tracer.StartSpan(
		funcName+"-GRPC",
		opentracing.ChildOf(parentSpan.Context()),
	)
	rsp, err := global.AdministratorClients[global.LoadBalancing(len(global.AdministratorInstances))].ExecuteCreation(context.Background(), &proto.ExecutionCreationRequest{
		Id:  int64(id),
		Opt: int32(opt),
	})
	grpcSpan.Finish()

	if err != nil {
		zap.S().Errorw("[CreationExecute]作品处理失败")
		utils.HandleGrpcError2Http(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": rsp.Msg,
		"msg":  nil,
	})
}

func (a Administrator) Bulletin(c *gin.Context) {
	tracerStr, _ := c.Get("tracer")
	parentSpanStr, _ := c.Get("parentSpan")
	tracer, _ := tracerStr.(opentracing.Tracer)
	parentSpan, _ := parentSpanStr.(opentracing.Span)

	funcName := strings.Split(c.HandlerName(), ".")[1]
	grpcSpan := tracer.StartSpan(
		funcName+"-GRPC",
		opentracing.ChildOf(parentSpan.Context()),
	)
	err := global.RDB.Set("bulletin", c.PostForm("data"), 0).Err()
	grpcSpan.Finish()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "操作完成",
		"msg":  err,
	})
}

func (a Administrator) GetUserFeedback(c *gin.Context) {
	tracerStr, _ := c.Get("tracer")
	parentSpanStr, _ := c.Get("parentSpan")
	tracer, _ := tracerStr.(opentracing.Tracer)
	parentSpan, _ := parentSpanStr.(opentracing.Span)

	funcName := strings.Split(c.HandlerName(), ".")[1]
	grpcSpan := tracer.StartSpan(
		funcName+"-GRPC",
		opentracing.ChildOf(parentSpan.Context()),
	)
	rsp, err := global.AdministratorClients[global.LoadBalancing(len(global.AdministratorInstances))].GetUserFeedback(context.Background(), &emptypb.Empty{})
	grpcSpan.Finish()

	type feedback struct {
		SubmitTime string `json:"time"`
		Mobile     string `json:"mobile"`
		Avatar     string `json:"avatar"`
		Nickname   string `json:"nickname"`
		Advice     string `json:"advice"`
		Evidence   string `json:"evidence"`
		Suspect    string `json:"suspect"`
	}
	var result []feedback
	for _, v := range rsp.UserFeedback {
		var t feedback
		t.SubmitTime = time.Unix(v.Time, 0).Format("2006-01-02 15:04:05")
		t.Mobile = v.Mobile
		t.Avatar = v.Avatar
		t.Nickname = v.Nickname
		t.Advice = v.Advice
		t.Evidence = v.Evidence
		t.Suspect = v.Suspect
		result = append(result, t)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
		"msg":  err,
	})
}
