package inte

import (
	"log"
	"net/http"
	"uapply/dao"
	"uapply/models"
	"uapply/utils/jwt"
	"uapply/utils/response"

	"github.com/gin-gonic/gin"
)

type InteLoginRsp struct {
	Token   string `json:"token"`
	OrgaID  uint   `json:"orga_id"`
	DepaID  uint   `json:"depa_id"`
	InteID  uint   `json:"inte_id"`
	Name    string `json:"name"`
	Account string `json:"account"`
}

func Login(c *gin.Context) {
	// 绑定前端数据
	var inteLoginInfo models.InteLoginInfo
	err := c.ShouldBindJSON(&inteLoginInfo)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println("bind err : ", err)
		return
	}

	var inteLoginInfoDb models.InteLoginInfo
	err = dao.GetDb().Where("account=?", inteLoginInfo.Account).First(&inteLoginInfoDb).Error
	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "account or password error")
			return
		}
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "try again")
		return
	}

	if inteLoginInfoDb.Account != inteLoginInfo.Account || inteLoginInfoDb.Password != inteLoginInfo.Password {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "account or password error")
		return
	}

	token, err := jwt.GenTokenInte(inteLoginInfoDb.ID, inteLoginInfoDb.DepaID, inteLoginInfoDb.OrgaID)
	if err != nil {
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "try again")
		return
	}
	dr := InteLoginRsp{
		OrgaID: inteLoginInfoDb.OrgaID,
		Name:   inteLoginInfoDb.Name,
		DepaID: inteLoginInfoDb.DepaID,
		InteID: inteLoginInfo.ID,
		Token:  token,
	}
	response.Success(c, dr)
}
