package main

import (
	"AIMusic-API/global"
	"AIMusic-API/middlewares"
	"AIMusic-API/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"./zap.log"}
	return cfg.Build()
}

func main() {
	global.Init() //初始化logger和config

	// Update service instances initially
	global.UpdateServiceInstances()
	global.UpdateGRPCConn()

	// Set up a ticker to update service instances periodically
	//ticker := time.NewTicker(30 * time.Minute)
	//go func() {
	//	for range ticker.C {
	//		global.UpdateServiceInstances()
	//	}
	//}()
	// Call your services using the updated serviceInstances list and your load balancing strategy

	//logger, err := NewLogger()
	//if err != nil {
	//	panic(err)
	//}
	//su := logger.Sugar()
	//defer su.Sync() //flush buffer

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.StaticFS("/ai", http.Dir("D:\\generations"))
	routers.InitRouters(r)
	err := r.Run(fmt.Sprintf(":%d", global.Config.API.Port))
	if err != nil {
		zap.S().Panic("启动服务器失败：", err.Error())
	}

	//midiFirst, err := ioutil.ReadFile("./midis/1.mid")
	//midiSecond, err := ioutil.ReadFile("./midis/2.mid")
	//
	//conn, err := grpcSRV.Dial("127.0.0.1:50051", grpcSRV.WithInsecure())
	//if err != nil {
	//	panic(err)
	//}
	//
	//c := Splice.NewAIMusicSpliceClient(conn)
	//response, err := c.MidiSplice(context.Background(), &Splice.InputFile{
	//	FirstName:     "1.mid",
	//	SecondName:    "2.mid",
	//	FirstContent:  midiFirst,
	//	SecondContent: midiSecond,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = ioutil.WriteFile("./midis/3.mid", response.Spliced, 0644)
	//if err != nil {
	//	panic(err)
	//}
}
