package handler

import (
	"community/global"
	"community/model"
	"community/proto"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommunityServer struct {
	proto.UnimplementedCommunityServer
}

func bindCreationStruct(dst []*proto.CommunityCreation, src []model.Composition) {
	for _, v := range src {
		tmp := new(proto.CommunityCreation)
		tmp.Id = int64(v.ID)
		tmp.ForWhat = v.CreatedFor
		tmp.UserName = v.CreatorNickname
		tmp.UserAvatar = v.Avatar
		tmp.CompositionName = v.CompositionName
		tmp.Mp3 = v.MP3
		tmp.Likes = v.Likes
		tmp.Report = v.Report
		tmp.Detail = v.Detail
		tmp.ChineseNote = v.ChineseNote
		tmp.ChineseEmotion = v.ChineseEmotion
		tmp.CreatedTime = uint64(v.CreatedAt.Unix())
		dst = append(dst, tmp)
	}
}

func (c *CommunityServer) GetTop(context.Context, *emptypb.Empty) (*proto.TopCreation, error) {
	var compositionLikes []model.Composition
	var compositionLikesGong []model.Composition
	var compositionLikesShang []model.Composition
	var compositionLikesJue []model.Composition
	var compositionLikesZhi []model.Composition
	var compositionLikesYu []model.Composition
	global.DB.Order("likes desc").Where("status = 1").Find(&compositionLikes).Limit(10)
	global.DB.Order("likes desc").Where("status = 1 AND chinese_note = gong").Find(&compositionLikesGong).Limit(10)
	global.DB.Order("likes desc").Where("status = 1 AND chinese_note = shang").Find(&compositionLikesShang).Limit(10)
	global.DB.Order("likes desc").Where("status = 1 AND chinese_note = jue").Find(&compositionLikesJue).Limit(10)
	global.DB.Order("likes desc").Where("status = 1 AND chinese_note = zhi").Find(&compositionLikesZhi).Limit(10)
	global.DB.Order("likes desc").Where("status = 1 AND chinese_note = yu").Find(&compositionLikesYu).Limit(10)
	var communityCreation []*proto.CommunityCreation
	var communityCreationGong []*proto.CommunityCreation
	var communityCreationShang []*proto.CommunityCreation
	var communityCreationJue []*proto.CommunityCreation
	var communityCreationZhi []*proto.CommunityCreation
	var communityCreationYu []*proto.CommunityCreation
	bindCreationStruct(communityCreation, compositionLikes)
	bindCreationStruct(communityCreationGong, compositionLikesGong)
	bindCreationStruct(communityCreationShang, compositionLikesShang)
	bindCreationStruct(communityCreationJue, compositionLikesJue)
	bindCreationStruct(communityCreationZhi, compositionLikesZhi)
	bindCreationStruct(communityCreationYu, compositionLikesYu)
	return &proto.TopCreation{
		Creation:      communityCreation,
		CreationGong:  communityCreationGong,
		CreationShang: communityCreationShang,
		CreationJue:   communityCreationJue,
		CreationZhi:   communityCreationZhi,
		CreationYu:    communityCreationYu}, nil
}

func (c *CommunityServer) GetBulletin(context.Context, *emptypb.Empty) (*proto.Bulletin, error) {
	bulletin := global.RDB.Get("bulletin").Val()
	return &proto.Bulletin{Msg: bulletin}, nil
}

func (c *CommunityServer) SearchUserByName(ctx context.Context, req *proto.SearchRequest) (*proto.SearchUserResponse, error) {
	var user []model.User
	req.Name = "%" + req.Name + "%"
	global.DB.Where("nick_name LIKE ?", req.Name).Find(&user)
	var userPb []*proto.SearchUser
	for _, v := range user {
		tmp := new(proto.SearchUser)
		tmp.Nickname = v.NickName
		tmp.Pic = v.Avatar
		tmp.Birthday = v.Birthday
		tmp.Motto = v.Motto
		//用户作品由user-GetCreation获得
		userPb = append(userPb, tmp)
	}
	return &proto.SearchUserResponse{Users: userPb}, nil
}

func (c *CommunityServer) SearchCreationByName(ctx context.Context, req *proto.SearchRequest) (*proto.SearchCreationResponse, error) {
	var creation []model.Composition
	req.Name = "%" + req.Name + "%"
	global.DB.Where("composition_name LIKE ? AND status = 2", req.Name).Find(&creation)
	var creationPb []*proto.PublicCreation
	for _, v := range creation {
		tmp := new(proto.PublicCreation)
		tmp.Id = int64(v.ID)
		tmp.Report = v.Report
		tmp.Likes = v.Likes
		tmp.Nickname = v.CreatorNickname
		tmp.Mp3 = v.MP3
		tmp.ForWhat = v.CreatedFor
		tmp.CompositionName = v.CompositionName
		tmp.Detail = v.Detail
		tmp.Avatar = v.Avatar
		tmp.ChineseNote = v.ChineseNote
		tmp.ChineseEmotion = v.ChineseEmotion
		creationPb = append(creationPb, tmp)
	}
	return &proto.SearchCreationResponse{Creations: creationPb}, nil
}

func (c *CommunityServer) CreationAttitude(ctx context.Context, req *proto.AttitudeCreation) (*proto.AttitudeResponse, error) {
	var composition model.Composition
	composition.ID = int(req.Id)
	global.DB.First(&composition)
	if req.Opt == 1 {
		composition.Likes = composition.Likes + 1
	} else if req.Opt == 0 {
		composition.Report = composition.Report + 1
	}
	global.DB.Save(&composition)
	return &proto.AttitudeResponse{Msg: "操作完成"}, nil
}
