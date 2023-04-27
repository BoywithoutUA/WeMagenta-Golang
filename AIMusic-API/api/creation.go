package api

import (
	"AIMusic-API/global"
	"AIMusic-API/proto"
	"AIMusic-API/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"os"
	"strconv"
)

type Creation struct{}

func (c Creation) GetTemplate(ctx *gin.Context) {
	rsp, err := global.CreationClients[global.LoadBalancing(len(global.CreationInstances))].GetTemplate(context.Background(), &emptypb.Empty{})
	if err != nil {
		zap.S().Errorw("[GetTemplate]获取创作模板失败")
		utils.HandleGrpcError2Http(err, ctx)
		return
	}
	type ginResponse struct {
		Name        string `json:"name"`
		Type        int32  `json:"type"`
		Description string `json:"description"`
		Pic         string `json:"pic"`
	}
	var result []ginResponse
	for _, v := range rsp.Template {
		tmp := ginResponse{
			Name:        v.Name,
			Type:        int32(v.Type),
			Description: v.Description,
			Pic:         v.Pic,
		}
		result = append(result, tmp)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
		"msg":  nil,
	})
	zap.S().Debug("[GetTemplate]获取创作模板成功")
}

func (c Creation) AddCreation(ctx *gin.Context) {
	//暂时没有登陆 直接传手机号
	mobile := ctx.PostForm("mobile")
	name := ctx.PostForm("name")
	forWhat := ctx.PostForm("forWhat")
	detail := ctx.PostForm("detail")
	mp3 := ctx.PostForm("mp3")
	optStr := ctx.PostForm("opt")
	chineseNote := ctx.PostForm("chineseNote")
	chineseEmotion := ctx.PostForm("chineseEmotion")
	opt, _ := strconv.Atoi(optStr)
	rsp, err := global.CreationClients[global.LoadBalancing(len(global.CreationInstances))].Add(context.Background(), &proto.AddRequest{
		Mobile:         mobile,
		Name:           name,
		ForWhat:        forWhat,
		Detail:         detail,
		Mp3:            mp3,
		Opt:            int32(opt),
		ChineseNote:    chineseNote,
		ChineseEmotion: chineseEmotion,
	})
	if err != nil {
		zap.S().Errorw("[AddCreation]添加用户作品失败")
		utils.HandleGrpcError2Http(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": rsp.Msg,
		"msg":  nil,
	})
	zap.S().Debug("[AddCreation]添加用户作品成功")
}

// GenerateMusic 生成一段音乐
func (c Creation) GenerateMusic(ctx *gin.Context) {
	mobile := ctx.PostForm("mobile")

	////布置Redis分布式锁等待音乐生成并上传至服务器
	////将当前用户状态标记为0 表示正在通过AI模型生成音乐
	////将当前用户状态标记为1时 表示通过AI模型生成音乐完成
	//err := global.RDB.SetNX(mobile, 0, 0).Err()
	//if err != nil {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"code": 500,
	//		"data": fmt.Sprintf("当前用户%s正在生成音乐中", mobile),
	//		"msg":  fmt.Sprintf("当前用户%s正在生成音乐中", mobile),
	//	})
	//	return
	//}

	reqTime := utils.GetUnix()
	path := fmt.Sprintf("D:\\generations\\%s\\scratch\\%d\\",
		mobile, reqTime)

	os.MkdirAll(path, os.ModePerm)

	err := MQScratch(path, mobile, reqTime)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"data": err.Error(),
			"msg":  err.Error(),
		})
		return
	}

	////检查AI音乐是否生成完成
	//retryTime := 0
	//for {
	//	retryTime++
	//	//休息1s
	//	time.Sleep(1 * time.Second)
	//	if global.RDB.Get(mobile).Val() == "1" || retryTime > 60 {
	//		//释放锁
	//		global.RDB.Del(mobile)
	//		break
	//	}
	//}

	//sshPath := fmt.Sprintf("/AIMusic/generations/%s/scratch/%d/", mobile, reqTime)
	//err = exec.Command("ssh", "root@106.55.143.205", fmt.Sprintf("mkdir -p /AIMusic/generations/%s/scratch/%d/", mobile, reqTime)).Run()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = exec.Command(fmt.Sprintf("scp -r %s root@106.55.143.205:/AIMusic/generations/%s/scratch/%d/",
	//	path+"wav", mobile, reqTime)).Run()
	//if err != nil {
	//	fmt.Println(err)
	//}

	//返回路径给前端
	type Wav struct {
		Left  string `json:"left"`
		Right string `json:"right"`
	}
	var wav Wav
	wav.Left = "ai" + fmt.Sprintf("/%s/scratch/%d/", mobile, reqTime) + "wav/0.wav"
	wav.Right = "ai" + fmt.Sprintf("/%s/scratch/%d/", mobile, reqTime) + "wav/1.wav"
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": wav,
		"msg":  nil,
	})
}

func MQScratch(path string, mobile string, reqTime int64) error {
	global.RDB.Incr("WeMagenta_Scratch_Times")
	err := global.MQCh.Publish(
		"",        // 交换的名字，空表示默认
		"scratch", // 路由键，也是队列的标识
		false,     // 返回值，必填
		false,     // 立即
		amqp.Publishing{
			ContentType: "text/plain",                               // 消息的格式
			Body:        []byte(fmt.Sprintf("%s|%s", path, mobile)), // 发送的消息是字节数组类型
		})
	return err
}
