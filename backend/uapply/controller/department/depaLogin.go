package depa

import (
	"log"
	"net/http"
	"uapply/dao"
	"uapply/models"
	"uapply/utils/jwt"
	"uapply/utils/response"

	"github.com/gin-gonic/gin"
)

type DepaResponse struct {
	OrgaID    uint   `json:"orga_id"`
	Name      string `json:"name"`
	Introduce string `json:"introduce"`
	DepaID    uint   `json:"depa_id"`
	Token     string `json:"token"`
}

func Login(c *gin.Context) {
	// 绑定前端数据
	var DepaLoginInfo models.DepaLoginInfo
	err := c.ShouldBindJSON(&DepaLoginInfo)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println("bind err : ", err)
		return
	}

	var DepaLoginInfoDb models.DepaLoginInfo
	err = dao.GetDb().Where("account=?", DepaLoginInfo.Account).First(&DepaLoginInfoDb).Error
	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "account or password error")
			return
		}
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "try again")
		return
	}

	if DepaLoginInfoDb.Account != DepaLoginInfo.Account || DepaLoginInfoDb.Password != DepaLoginInfo.Password {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "account or password error")
		return
	}

	token, err := jwt.GenToken(DepaLoginInfoDb.ID, jwt.Depa)
	if err != nil {
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "try again")
		return
	}
	dr := DepaResponse{
		OrgaID:    DepaLoginInfoDb.OrgaID,
		Name:      DepaLoginInfoDb.Name,
		Introduce: DepaLoginInfoDb.Introduce,
		DepaID:    DepaLoginInfoDb.ID,
		Token:     token,
	}
	response.Success(c, dr)
}
