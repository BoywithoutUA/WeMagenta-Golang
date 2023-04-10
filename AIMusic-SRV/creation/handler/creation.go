package handler

import (
	"context"
	"creation/global"
	"creation/model"
	"creation/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CreationServer struct {
	proto.UnimplementedCreationServer
}

func (c *CreationServer) GetTemplate(context.Context, *emptypb.Empty) (*proto.GetTemplateResponse, error) {
	var templatesDB []model.Template
	global.DB.Find(&templatesDB)
	var templatesProto []*proto.Template
	for _, v := range templatesDB {
		tmp := new(proto.Template)
		tmp.Name = v.Name
		tmp.Type = int64(v.Type)
		tmp.Description = v.Detail
		tmp.Pic = v.Pic
		templatesProto = append(templatesProto, tmp)
	}
	return &proto.GetTemplateResponse{
		Template: templatesProto,
	}, nil
}

func (c *CreationServer) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	var user model.User
	result := global.DB.Where("mobile = ?", req.Mobile).Find(&user)
	if result.Error != nil {
		return &proto.AddResponse{Msg: "获取用户信息失败"}, result.Error
	}
	var composition model.Composition
	composition.CompositionName = req.Name
	composition.CreatorMobile = req.Mobile
	composition.CreatorNickname = user.NickName
	composition.CreatedFor = req.ForWhat
	composition.Detail = req.Detail
	composition.MP3 = req.Mp3
	composition.Status = req.Opt
	composition.Avatar = user.Avatar
	result = global.DB.Create(&composition)
	global.RDB.ZIncrBy("template", 1, req.ForWhat)
	if result.Error != nil {
		return &proto.AddResponse{Msg: "添加作品信息失败"}, result.Error
	}
	return &proto.AddResponse{Msg: "添加作品信息完成"}, nil
}
