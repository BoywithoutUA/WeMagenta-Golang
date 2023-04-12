package middlewares

import (
	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var ruleMap = make(map[string]*flow.Rule)

//初始化一些规则
func init() {
	err := api.InitDefault()
	if err != nil {
		log.Fatalf("Sentinel init failed: %v", err)
	}

	// QPS限流规则
	ruleMap["Scratch"] = &flow.Rule{
		Resource:               "Scratch",
		TokenCalculateStrategy: flow.Direct,
		ControlBehavior:        flow.Reject,
		Threshold:              1,
		StatIntervalInMs:       2000,
	}

	ruleMap["CRUD"] = &flow.Rule{
		Resource:               "CRUD",
		TokenCalculateStrategy: flow.Direct,
		ControlBehavior:        flow.Reject,
		Threshold:              2, // QPS限制为100
		StatIntervalInMs:       1000,
	}

	// Warm Up冷启动限流规则
	ruleMap["WarmUp"] = &flow.Rule{
		Resource:               "WarmUp",
		TokenCalculateStrategy: flow.WarmUp,
		ControlBehavior:        flow.Reject,
		Threshold:              10,
		StatIntervalInMs:       1000,
		WarmUpPeriodSec:        5, // 预热时间为5秒
		WarmUpColdFactor:       5, // 冷启动因子为5
	}
	// 将规则加载到Sentinel
	rules := make([]*flow.Rule, 0, len(ruleMap))
	for _, rule := range ruleMap {
		rules = append(rules, rule)
	}
	_, err = flow.LoadRules(rules)
	if err != nil {
		log.Printf("failed to load rules: %v\n", err)
	}
}

func Sentinel(resourceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		e, b := api.Entry(resourceName,
			api.WithResourceType(base.ResTypeWeb), api.WithTrafficType(base.Inbound))
		if b != nil {
			// 请求被 Sentinel 限流，降级或者熔断
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code": 500,
				"data": "Sentinel Block",
				"msg":  "Sentinel Block",
			})
			c.Abort()
			return
		}
		// 请求正常，继续执行后续中间件和处理函数
		c.Next()
		// 释放资源
		e.Exit()
	}

}
