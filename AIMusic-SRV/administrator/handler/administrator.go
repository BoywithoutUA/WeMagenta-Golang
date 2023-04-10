package handler

import (
	"administrator/global"
	"administrator/model"
	"administrator/proto"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type AdministratorServer struct {
	proto.UnimplementedAdministratorServer
}

func (a *AdministratorServer) GetUsage(context.Context, *emptypb.Empty) (*proto.Usage, error) {
	var usage []model.Usage
	global.DB.Where("add_time > ?", time.Now().AddDate(0, 0, -8)).Order("add_time asc").Find(&usage)
	var scratch []*proto.UsageScratch
	var composition []*proto.UsageComposition
	for _, v := range usage {
		tmpScratch := new(proto.UsageScratch)
		tmpComposition := new(proto.UsageComposition)
		tmpScratch.Amount = v.Scratch
		tmpComposition.Amount = v.Composition
		tmpScratch.Date = v.CreatedAt.Format("2006-01-02")
		tmpComposition.Date = v.CreatedAt.Format("2006-01-02")
		scratch = append(scratch, tmpScratch)
		composition = append(composition, tmpComposition)
	}
	var forWhat []*proto.ForWhat
	result := global.RDB.ZRevRangeWithScores("template", 0, 7).Val()
	for _, v := range result {
		t := new(proto.ForWhat)
		t.Amount = int64(v.Score)
		t.Name = v.Member.(string)
		forWhat = append(forWhat, t)
	}
	return &proto.Usage{
		Scratch:     scratch,
		Composition: composition,
		ForWhat:     forWhat,
	}, nil
}
func (a *AdministratorServer) ExecuteUser(ctx context.Context, req *proto.ExecutionUserRequest) (*proto.ExecutionResponse, error) {
	var user model.User
	global.DB.Where("mobile = ?", req.Mobile).First(&user)
	user.Banned = req.Opt
	global.DB.Save(&user)
	return &proto.ExecutionResponse{Msg: "操作完成"}, nil
}
func (a *AdministratorServer) ExecuteCreation(ctx context.Context, req *proto.ExecutionCreationRequest) (*proto.ExecutionResponse, error) {
	var composition model.Composition
	global.DB.Where("id = ?", req.Id).First(&composition)
	if req.Opt == 1 {
		composition.Status = 3
	} else if req.Opt == 0 {
		composition.Status = 2
	}
	global.DB.Save(&composition)
	return &proto.ExecutionResponse{Msg: "操作完成"}, nil
}
func (a *AdministratorServer) EditBulletin(ctx context.Context, req *proto.EditBulletinMsg) (*proto.EditBulletinMsg, error) {
	global.RDB.Set("bulletin", "欢迎来到WeMagenta", 0)
	return &proto.EditBulletinMsg{Msg: "操作完成"}, nil
}
func (a *AdministratorServer) UserAudit(context.Context, *emptypb.Empty) (*proto.UserAuditResponse, error) {
	var userLikes []model.User
	var userReport []model.User
	global.DB.Limit(10).Order("likes desc").Find(&userLikes)
	global.DB.Limit(10).Order("report desc").Find(&userReport)
	var likesUser []*proto.TopLikesUser
	var reportUser []*proto.TopReportUser
	for _, v := range userLikes {
		tmp := new(proto.TopLikesUser)
		tmp.Mobile = v.Mobile
		tmp.Nickname = v.NickName
		tmp.Likes = v.Likes
		tmp.Avatar = v.Avatar
		tmp.Motto = v.Motto
		likesUser = append(likesUser, tmp)
	}
	for _, v := range userReport {
		tmp := new(proto.TopReportUser)
		tmp.Mobile = v.Mobile
		tmp.Nickname = v.NickName
		tmp.Report = v.Report
		tmp.Avatar = v.Avatar
		tmp.Motto = v.Motto
		reportUser = append(reportUser, tmp)
	}
	return &proto.UserAuditResponse{
		UserLikes:  likesUser,
		UserReport: reportUser,
	}, nil
}
func (a *AdministratorServer) CreationAudit(context.Context, *emptypb.Empty) (*proto.CreationAuditResponse, error) {
	var compositionLikes []model.Composition
	var compositionReport []model.Composition
	global.DB.Limit(10).Order("likes desc").Find(&compositionLikes)
	global.DB.Limit(10).Order("report desc").Find(&compositionReport)
	var likesCreation []*proto.TopLikesCreation
	var reportCreation []*proto.TopReportCreation
	for _, v := range compositionLikes {
		tmp := new(proto.TopLikesCreation)
		tmp.Id = int64(v.ID)
		tmp.Name = v.CompositionName
		tmp.Likes = v.Likes
		tmp.Mp3 = v.MP3
		tmp.Detail = v.Detail
		tmp.Nick = v.CreatorNickname
		tmp.Avatar = v.Avatar
		likesCreation = append(likesCreation, tmp)
	}
	for _, v := range compositionReport {
		tmp := new(proto.TopReportCreation)
		tmp.Id = int64(v.ID)
		tmp.Name = v.CompositionName
		tmp.Report = v.Report
		tmp.Mp3 = v.MP3
		tmp.Detail = v.Detail
		tmp.Avatar = v.Avatar
		reportCreation = append(reportCreation, tmp)
	}
	return &proto.CreationAuditResponse{
		CreationLikes:  likesCreation,
		CreationReport: reportCreation,
	}, nil
}

func (a *AdministratorServer) GetUserFeedback(context.Context, *emptypb.Empty) (*proto.UserFeedbackResponse, error) {
	var feedback []model.Feedback
	global.DB.Order("add_time desc").Find(&feedback)
	var result []*proto.UserFeedback
	for _, v := range feedback {
		t := new(proto.UserFeedback)
		t.Evidence = v.Evidence
		t.Suspect = v.Target
		t.Time = v.CreatedAt.Unix()
		t.Advice = v.Advice
		t.Avatar = v.Avatar
		t.Nickname = v.UserNickName
		t.Mobile = v.UserMobile
		result = append(result, t)
	}
	return &proto.UserFeedbackResponse{UserFeedback: result}, nil
}
