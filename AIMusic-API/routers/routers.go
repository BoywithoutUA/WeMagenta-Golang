package routers

import (
	"AIMusic-API/api"
	"AIMusic-API/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters(r *gin.Engine) {
	r.GET("/consul/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	Administrator := r.Group("/administrator").Use(middlewares.Tracer("Administrator"))
	{
		Administrator.GET("/usage", api.Administrator{}.Usage).Use(middlewares.Sentinel("WarmUp"))
		Administrator.GET("/userAudit", api.Administrator{}.UserAudit).Use(middlewares.Sentinel("CRUD"))
		Administrator.GET("/creationAudit", api.Administrator{}.CreationAudit).Use(middlewares.Sentinel("CRUD"))
		Administrator.POST("/userExecute", api.Administrator{}.UserExecute).Use(middlewares.Sentinel("CRUD"))
		Administrator.POST("/creationExecute", api.Administrator{}.CreationExecute).Use(middlewares.Sentinel("CRUD"))
		Administrator.POST("/bulletin", api.Administrator{}.Bulletin).Use(middlewares.Sentinel("WarmUp"))
		Administrator.GET("/userFeedback", api.Administrator{}.GetUserFeedback).Use(middlewares.Sentinel("CRUD"))
	}

	Community := r.Group("/community").Use(middlewares.Tracer("Community"))
	{
		Community.GET("/bulletin", api.Community{}.Bulletin).Use(middlewares.Sentinel("WarmUp"))
		Community.GET("/creationTop", api.Community{}.CreationTop).Use(middlewares.Sentinel("WarmUp"))
		Community.GET("/userSearch", api.Community{}.UserSearch).Use(middlewares.Sentinel("CRUD"))
		Community.GET("/creationSearch", api.Community{}.CreationSearch).Use(middlewares.Sentinel("CRUD"))
		Community.POST("/attitude", api.Community{}.Attitude).Use(middlewares.Sentinel("CRUD"))
	}

	CreationRouters := r.Group("/creation").Use(middlewares.Tracer("Creation"))
	{
		CreationRouters.POST("/scratch", api.Creation{}.GenerateMusic).Use(middlewares.Sentinel("Scratch"))
		CreationRouters.GET("/template", api.Creation{}.GetTemplate).Use(middlewares.Sentinel("WarmUp"))
		CreationRouters.POST("/add", api.Creation{}.AddCreation).Use(middlewares.Sentinel("CRUD"))
	}

	UserRouters := r.Group("/user").Use(middlewares.Tracer("User"))
	{
		UserRouters.GET("/getInfo", api.User{}.GetInfo).Use(middlewares.Sentinel("WarmUp"))
		UserRouters.GET("/getCreation", api.User{}.GetCreation).Use(middlewares.Sentinel("CRUD"))
		UserRouters.POST("/feedback", api.User{}.Feedback).Use(middlewares.Sentinel("CRUD"))
	}
}
