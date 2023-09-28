package user

import (
	"log"
	"net/http"
	"strings"
	"uapply/dao"
	"uapply/models"
	"uapply/utils/auth"
	"uapply/utils/jwt"
	"uapply/utils/response"

	"github.com/gin-gonic/gin"
)

type UserRsp struct {
	Token   string `json:"token"`
	Account string `json:"account"`
}

func Register(c *gin.Context) {
	// 绑定前端数据
	var loginInfo models.UserLoginInfo
	err := c.ShouldBindJSON(&loginInfo)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println("bind err : ", err)
		return
	}

	err = dao.GetDb().Create(&loginInfo).Error
	if err != nil {
		log.Println(err)
		split := strings.Split(err.Error(), " ")
		if split[1] == "1062" {
			response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "ACCOUNT EXIST")
			return
		}
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "database insert error")
		return
	}
	err = dao.GetDb().Create(models.UserCV{UserID: loginInfo.ID, IsInit: 1}).Error
	if err != nil {
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "Create CV error")
		return
	}
	response.Success(c, loginInfo.Account)
}

func Login(c *gin.Context) {
	// 绑定前端数据
	var UserLoginInfo models.UserLoginInfo
	err := c.ShouldBindJSON(&UserLoginInfo)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println("bind err : ", err)
		return
	}

	var UserLoginInfoDb models.UserLoginInfo
	err = dao.GetDb().Where("account=?", UserLoginInfo.Account).First(&UserLoginInfoDb).Error
	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "account or password error")
			return
		}
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "try again")
		return
	}

	if UserLoginInfoDb.Account != UserLoginInfo.Account || UserLoginInfoDb.Password != UserLoginInfo.Password {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "account or password error")
		return
	}

	token, err := jwt.GenToken(UserLoginInfoDb.ID, jwt.User)
	if err != nil {
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "try again")
		return
	}
	rsp := UserRsp{
		Token:   token,
		Account: UserLoginInfoDb.Account,
	}
	response.Success(c, rsp)
}

func GetAccount(c *gin.Context) {
	idAny, exist := c.Get(auth.UserIDKey)
	id := idAny.(uint)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}
	var login models.UserLoginInfo

	err := dao.GetDb().Where("id=?", id).First(&login).Error
	if err != nil {
		log.Println(err)
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "database get cv error")
		return
	}
	response.Success(c, login)
}
