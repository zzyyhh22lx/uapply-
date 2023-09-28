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
	SetInteRouter(r)

	return r
}

func SetUserRouter(r *gin.Engine) {
	// 注册
	r.POST(userUrl+"register", user_ctr.Register)
	// 登录
	r.POST(userUrl+"login", user_ctr.Login)

	user := r.Group(userUrl)
	user.Use(auth.JWTAuthUser())
	user_no_token := r.Group(userUrl)
	{
		// 根据 token 获取用户 account
		user.GET("get-account", user_ctr.GetAccount)

		// 提交简历到 mysql
		user.POST("save-cv/commit", user_ctr.CommitCv)

		// 从 mysql 查询简历
		user.GET("get-cv/complete", user_ctr.GetCv)

		// 获取所有组织和部门
		user_no_token.GET("get-all", user_ctr.GetAllOrgaDepa)

		// 通过组织 id 获取底下所有部门，用于报名的时候
		user_no_token.GET("get-all-dep", user_ctr.GetAllDepaByOrga)

		// 报名
		user.POST("registration", user_ctr.Registration)

		// 获取流程状态
		user.GET("registration", user_ctr.GetRegistration)
	}
}

func SetOrgaRouter(r *gin.Engine) {
	// 注册
	r.POST(orgaUrl+"register", orga_ctr.Register)
	// 登录
	r.POST(orgaUrl+"login", orga_ctr.Login)

	orga := r.Group(orgaUrl)
	orga.Use(auth.JWTAuthOrga())
	{
		// 部门注册
		orga.POST("register-depa", orga_ctr.RegisterDepa)

		// 获取全部部门
		orga.GET("get-all-depa", orga_ctr.GetAllDepa)

		// 设置招新起止时间
		orga.POST("set-time", orga_ctr.SetTime)

		// 设置组织下最多能同时报名 n 个社团
		orga.POST("set-max-apply", orga_ctr.SetMaxApply)

		// 获取组织全部部门的全部候选人信息和状态
		orga.GET("get-reg-users", orga_ctr.GetRegUsers)
	}
}

func SetDepaRouter(r *gin.Engine) {
	// 部门由组织注册
	// 登录
	r.POST(depaUrl+"login", depa_ctr.Login)

	depa := r.Group(depaUrl)
	depa.Use(auth.JWTAuthDepa())
	{
		// 添加面试官
		depa.POST("add-inte", depa_ctr.AddInterviewer)

		// 获取全部面试官
		depa.GET("get-all-inte", depa_ctr.GetAllInterviewer)

		// 获取组织和部门全部信息
		depa.GET("get-org-dep", depa_ctr.GetOrgaDepa)

		// 获取部门的全部候选人信息和状态
		depa.GET("get-reg-users", depa_ctr.GetRegUsers)
	}
}

func SetInteRouter(r *gin.Engine) {
	// 面试官由部门注册
	// 登录
	r.POST(inteUrl + "login")

	inte := r.Group(inteUrl)
	inte.Use(auth.JWTAuthInte())
	{
		// 获取部门和组织信息

		// 获取面试官部门的全部候选人信息和状态

		// 推进状态，进入下一轮面试

		// 通过面试（是否发送短信）

		// 淘汰（是否发送短信）

		// 填写评价
	}
}
