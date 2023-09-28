package depa

import (
	"log"
	"net/http"
	"uapply/dao"
	"uapply/models"
	"uapply/utils/auth"
	"uapply/utils/response"

	"github.com/gin-gonic/gin"
)

// type InteInfoRsp struct {
// 	OrgaID uint
// 	DepaID uint
// 	Name   string
// }

func AddInterviewer(c *gin.Context) {
	idAny, exist := c.Get(auth.DepaIDKey)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}
	depaId := idAny.(uint)
	idAny, exist = c.Get(auth.OrgaIDKey)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}
	orgaId := idAny.(uint)
	// 绑定前端数据
	var InteLoginInfo models.InteLoginInfo
	err := c.ShouldBindJSON(&InteLoginInfo)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println("bind err : ", err)
		return
	}
	if InteLoginInfo.Name == "" {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "名称不能为空")
		return
	}
	InteLoginInfo.DepaID = depaId
	InteLoginInfo.OrgaID = orgaId
	err = dao.GetDb().Create(&InteLoginInfo).Error
	if err != nil {
		log.Println(err)
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "database insert error")
		return
	}
	response.Success(c, "set success")
}

func GetAllInterviewer(c *gin.Context) {
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
	var InteLoginInfos []models.InteLoginInfo
	err := dao.GetDb().Where("orga_id=? and depa_id=?", orgaId, depaId).Find(&InteLoginInfos).Error
	if err != nil {
		log.Println(err)
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "find inter error")
		return
	}
	// var InteRsps []InteInfoRsp
	// for _, info := range InteLoginInfos {
	// 	InteRsps = append(InteRsps, InteInfoRsp{
	// 		Name:   info.Name,
	// 		OrgaID: info.OrgaID,
	// 		DepaID: info.DepaID,
	// 	})
	// }
	response.Success(c, InteLoginInfos)
}
