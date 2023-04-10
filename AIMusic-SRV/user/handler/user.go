package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"user/global"
	"user/model"
	"user/proto"
)

type UserServer struct {
	proto.UnimplementedUserServer
}

//func Model2UserResponse(user model.User) proto.RegisterInfo {
//	//在grpc的message中字符应有默认值 不能随意赋值为nil
//	rsp := proto.RegisterInfo{
//		Id:       user.ID,
//		UserName: user.UserName,
//		Password: user.Password,
//		Mobile:   user.Mobile,
//		NickName: user.NickName,
//		Birthday: user.Birthday,
//		Gender:   user.Gender,
//		Role:     user.Role,
//	}
//	return rsp
//}

//// Register 注册新用户
//func (s *UserServer) Register(ctx context.Context, req *proto.RegisterInfo) (*emptypb.Empty, error) {
//	//查询账号是否已经存在
//	var userCheck model.User
//	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&userCheck)
//	if result.RowsAffected != 0 {
//		return nil, status.Error(codes.AlreadyExists, "手机号已存在")
//	}
//	user := model.User{
//		Mobile:   req.Mobile,
//		NickName: req.NickName,
//		Birthday: req.Birthday,
//		Gender:   req.Gender,
//		Role:     req.Role,
//	}
//	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 48, HashFunction: sha512.New}
//	salt, encodedPwd := password.Encode(req.Password, options)
//	user.Password = fmt.Sprintf("pbkdf2-sha512$%s$%s", salt, encodedPwd)
//	result = global.DB.Create(&user)
//	if result.Error != nil {
//		return nil, status.Errorf(codes.Internal, result.Error.Error())
//	}
//	return &empty.Empty{}, nil
//}

func (u *UserServer) GetInfo(ctx context.Context, req *proto.GetRequest) (*proto.InfoResponse, error) {
	var user model.User
	result := global.DB.Where("mobile = ?", req.Mobile).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "获取用户信息失败"+result.Error.Error())
	}

	var genderStr string
	if user.Gender == 0 {
		genderStr = "男"
	} else if user.Gender == 1 {
		genderStr = "女"
	} else {
		genderStr = "错误性别类型"
	}

	return &proto.InfoResponse{
		Nickname: user.NickName,
		Motto:    user.Motto,
		Avatar:   user.Avatar,
		Gender:   genderStr,
		Birthday: user.Birthday,
		Likes:    user.Likes,
	}, nil
}

func (u *UserServer) GetCreation(ctx context.Context, req *proto.GetRequest) (*proto.CreationResponse, error) {
	var compositions []model.Composition
	result := global.DB.Where("creator_mobile = ?", req.Mobile).Find(&compositions)
	if result.Error != nil {
		return nil, status.Errorf(codes.NotFound, "获取用户作品失败"+result.Error.Error())
	}
	var creations []*proto.UserCreation
	for _, v := range compositions {
		tmp := new(proto.UserCreation)
		tmp.Name = v.CompositionName
		tmp.CreationTime = v.CreatedAt.String()
		tmp.Status = v.Status
		tmp.Detail = v.Detail
		tmp.ForWhat = v.CreatedFor
		tmp.LastModified = v.UpdatedAt.String()
		tmp.Mp3 = v.MP3
		creations = append(creations, tmp)
	}
	return &proto.CreationResponse{
		Creation: creations,
	}, nil
}

func (u *UserServer) Feedback(ctx context.Context, req *proto.FeedbackRequest) (*emptypb.Empty, error) {
	var feedback model.Feedback
	feedback.Advice = req.Advice
	feedback.Avatar = req.Avatar
	feedback.Target = req.Suspect
	feedback.Evidence = req.Evidence
	feedback.UserMobile = req.Mobile
	feedback.UserNickName = req.Nickname
	global.DB.Create(&feedback)
	return nil, nil
}

//Update 更新用户信息
//func (s *UserServer) Update(ctx context.Context, req *proto.UpdateInfo) (*emptypb.Empty, error) {
//	var user model.User
//	result := global.DB.First(&user, req.Mobile)
//	if result.RowsAffected == 0 {
//		return nil, status.Errorf(codes.NotFound, "用户不存在")
//	}
//	user.Mobile = req.Mobile
//	user.NickName = req.NickName
//	user.Birthday = req.Birthday
//	user.Gender = req.Gender
//	result = global.DB.Save(&user)
//	if result.Error != nil {
//		return nil, status.Error(codes.Internal, result.Error.Error())
//	}
//	return &empty.Empty{}, nil
//}

//// Authenticate 用户登录鉴权
//func (s *UserServer) Authenticate(ctx context.Context, req *proto.AuthenticationInfo) (*proto.AuthenticationResponse, error) {
//	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 48, HashFunction: sha512.New}
//	passwordInfo := strings.Split(req.EncryptedPassword, "$")
//	check := password.Verify(req.Password, passwordInfo[1], passwordInfo[2], options)
//	return &proto.AuthenticationResponse{Success: check}, nil
//}
