package routers

import (
	user_ctr "uapply/controller/user"
	"uapply/utils/auth"
	"uapply/utils/cors"
	"uapply/utils/logger"

	"github.com/gin-gonic/gin"
)

const (
	baseUrl = "uapply"
	userUrl = baseUrl + "/user/"
)

func SetRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), cors.Cors()) // 跨域等等

	SetUserRouter(r)
	// 还要 manager、worker 等仿造上面的进行添加，在设计过程中添加任何需要的接口
	return r
}

func SetUserRouter(r *gin.Engine) {
	// 注册
	r.POST(userUrl+"register", user_ctr.Register)
	// 登录
	r.POST(userUrl+"login", user_ctr.Login)

	user := r.Group(userUrl)
	user.Use(auth.JWTAuthUser())
	{
		user.POST("save-cv/commit", user_ctr.CommitCv)

		user.GET("get-cv/complete", user_ctr.GetCv)
	}
}
