package routers

import (
	depa_ctr "uapply/controller/department"
	orga_ctr "uapply/controller/organization"
	user_ctr "uapply/controller/user"
	"uapply/utils/auth"
	"uapply/utils/cors"
	"uapply/utils/logger"

	"github.com/gin-gonic/gin"
)

const (
	baseUrl = "uapply"
	userUrl = baseUrl + "/user/"
	orgaUrl = baseUrl + "/orga/"
	depaUrl = baseUrl + "/depa/"
	inteUrl = baseUrl + "/inte/"
)

func SetRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), cors.Cors()) // 跨域等等

	// 用户接口
	SetUserRouter(r)

	// 组织接口
	SetOrgaRouter(r)

	// 部门接口
	SetDepaRouter(r)

	// 面试官接口
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
		// 提交简历到 mysql
		user.POST("save-cv/commit", user_ctr.CommitCv)

		// 从 mysql 查询简历
		user.GET("get-cv/complete", user_ctr.GetCv)

		// 获取所有组织和部门

		// 通过组织 id 获取底下所有部门
	}
}

func SetOrgaRouter(r *gin.Engine) {
	// 注册
	r.POST(orgaUrl+"register", orga_ctr.Register)
	// 登录
	r.POST(orgaUrl+"login", orga_ctr.Login)

	depa := r.Group(orgaUrl)
	depa.Use(auth.JWTAuthOrga())
	{
		// 部门注册
		depa.POST("register-depa", orga_ctr.RegisterDepa)

		// 获取全部部门

		// 设置招新起止时间
	}
}

func SetDepaRouter(r *gin.Engine) {
	// 部门由组织注册
	// 登录
	r.POST(depaUrl+"login", depa_ctr.Login)

	depa := r.Group(depaUrl)
	depa.Use(auth.JWTAuthDepa())
	{

	}
}
