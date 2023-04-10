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
		Administrator.GET("/usage", api.Administrator{}.Usage)
		Administrator.GET("/userAudit", api.Administrator{}.UserAudit)
		Administrator.GET("/creationAudit", api.Administrator{}.CreationAudit)
		Administrator.POST("/userExecute", api.Administrator{}.UserExecute)
		Administrator.POST("/creationExecute", api.Administrator{}.CreationExecute)
		Administrator.POST("/bulletin", api.Administrator{}.Bulletin)
		Administrator.GET("/userFeedback", api.Administrator{}.GetUserFeedback)
	}

	Community := r.Group("/community").Use(middlewares.Tracer("Community"))
	{
		Community.GET("/bulletin", api.Community{}.Bulletin)
		Community.GET("/creationTop", api.Community{}.CreationTop)
		Community.GET("/userSearch", api.Community{}.UserSearch)
		Community.GET("/creationSearch", api.Community{}.CreationSearch)
		Community.POST("/attitude", api.Community{}.Attitude)
	}

	CreationRouters := r.Group("/creation").Use(middlewares.Tracer("Creation"))
	{
		CreationRouters.POST("/scratch", api.Creation{}.GenerateMusic)
		CreationRouters.GET("/template", api.Creation{}.GetTemplate)
		CreationRouters.POST("/add", api.Creation{}.AddCreation)
	}

	UserRouters := r.Group("/user").Use(middlewares.Tracer("User"))
	{
		UserRouters.GET("/getInfo", api.User{}.GetInfo)
		UserRouters.GET("/getCreation", api.User{}.GetCreation)
		UserRouters.POST("/feedback", api.User{}.Feedback)
	}
}
