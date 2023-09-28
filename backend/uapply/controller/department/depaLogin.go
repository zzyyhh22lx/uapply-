package depa

import (
	"log"
	"net/http"
	"uapply/dao"
	"uapply/models"
	"uapply/utils/auth"
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

	token, err := jwt.GenTokneDepa(DepaLoginInfoDb.ID, DepaLoginInfoDb.OrgaID)
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

type DepaRsp struct {
	OrgaID        uint   `json:"orga_id"`
	OrgaName      string `json:"orga_name"`
	OrgaIntroduce string `json:"orga_introduce"`
	OrgaStartTime int64  `json:"orga_start_time"`
	OrgaEndTime   int64  `json:"orga_end_time"`
	OrgaMaxApply  int    `json:"orga_max_apply"`
	DepaID        uint   `json:"depa_id"`
	DepaName      string `json:"depa_name"`
	DepaIntrodece string `json:"depa_introduce"`
}

func GetOrgaDepa(c *gin.Context) {
	idAny, exist := c.Get(auth.DepaIDKey)
	depaId := idAny.(uint)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}
	idAny, exist = c.Get(auth.OrgaIDKey)
	orgaId := idAny.(uint)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}
	var depaInfo models.DepaLoginInfo
	var orgaInfo models.OrgaLoginInfo
	err := dao.GetDb().Where("id=?", orgaId).Find(&orgaInfo).Error
	if err != nil {
		log.Println(err)
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "find orga error")
		return
	}
	err = dao.GetDb().Where("id=?", depaId).Find(&depaInfo).Error
	if err != nil {
		log.Println(err)
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "find depa error")
		return
	}
	rsp := DepaRsp{
		OrgaID:        orgaInfo.ID,
		OrgaName:      orgaInfo.Name,
		OrgaIntroduce: orgaInfo.Introduce,
		OrgaStartTime: orgaInfo.StartTime,
		OrgaEndTime:   orgaInfo.EndTime,
		OrgaMaxApply:  orgaInfo.MaxApply,
		DepaID:        depaInfo.ID,
		DepaName:      depaInfo.Name,
		DepaIntrodece: depaInfo.Introduce,
	}
	response.Success(c, rsp)
}
