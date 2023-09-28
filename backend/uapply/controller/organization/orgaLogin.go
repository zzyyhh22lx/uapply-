package orga

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

type OrgaResponse struct {
	OrgaID    uint   `json:"orga_id"`
	Name      string `json:"name"`
	Introduce string `json:"introduce"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	MaxApply  int    `json:"max_apply"`
	Token     string `json:"token"`
}

func Register(c *gin.Context) {
	// 绑定前端数据
	var loginInfo models.OrgaLoginInfo
	err := c.ShouldBindJSON(&loginInfo)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println("bind err : ", err)
		return
	}

	if loginInfo.Name == "" {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "组织名字不能为空")
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
	response.Success(c, nil)
}

func Login(c *gin.Context) {
	// 绑定前端数据
	var OrgaLoginInfo models.OrgaLoginInfo
	err := c.ShouldBindJSON(&OrgaLoginInfo)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println("bind err : ", err)
		return
	}

	var OrgaLoginInfoDb models.OrgaLoginInfo
	err = dao.GetDb().Where("account=?", OrgaLoginInfo.Account).First(&OrgaLoginInfoDb).Error
	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "account or password error")
			return
		}
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "try again")
		return
	}

	if OrgaLoginInfoDb.Account != OrgaLoginInfo.Account || OrgaLoginInfoDb.Password != OrgaLoginInfo.Password {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "account or password error")
		return
	}

	token, err := jwt.GenToken(OrgaLoginInfoDb.ID, jwt.Orga)
	if err != nil {
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "try again")
		return
	}
	or := OrgaResponse{
		OrgaID:    OrgaLoginInfoDb.ID,
		Name:      OrgaLoginInfoDb.Name,
		Introduce: OrgaLoginInfoDb.Introduce,
		StartTime: OrgaLoginInfoDb.StartTime,
		EndTime:   OrgaLoginInfoDb.EndTime,
		MaxApply:  OrgaLoginInfoDb.MaxApply,
		Token:     token,
	}
	response.Success(c, or)
}

func RegisterDepa(c *gin.Context) {
	idAny, exist := c.Get(auth.OrgaIDKey)
	id := idAny.(uint)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}
	// 绑定前端数据
	var loginInfo models.DepaLoginInfo
	err := c.ShouldBindJSON(&loginInfo)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println("bind err : ", err)
		return
	}
	if loginInfo.Name == "" {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "部门名字不能为空")
		return
	}
	loginInfo.OrgaID = id
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
	response.Success(c, nil)
}
