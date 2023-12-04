package inte

import (
	"log"
	"net/http"
	depa "uapply/controller/department"
	"uapply/dao"
	"uapply/models"
	"uapply/utils/auth"
	"uapply/utils/response"

	"github.com/gin-gonic/gin"
)

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
	rsp := depa.DepaRsp{
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
