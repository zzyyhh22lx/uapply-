package orga

import (
	"log"
	"net/http"
	"uapply/dao"
	"uapply/models"
	"uapply/utils/auth"
	"uapply/utils/response"

	"github.com/gin-gonic/gin"
)

type OrgDepa struct {
	DepaID        uint   `json:"depa_id"`
	DepaName      string `json:"depa_name"`
	DepaIntroduce string `json:"depa_intro"`
}

type OrgaTime struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type OrgaMaxApply struct {
	MaxApply int `json:"max_apply"`
}

func GetAllDepa(c *gin.Context) {
	idAny, exist := c.Get(auth.OrgaIDKey)
	orgaId := idAny.(uint)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}
	// var depasInfos []OrgDepa
	var depaInfos []models.DepaLoginInfo
	err := dao.GetDb().Where("orga_id=?", orgaId).Find(&depaInfos).Error
	if err != nil {
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "get depa by orga error")
		return
	}
	// for _, depa := range depaInfos {
	// 	depasInfos = append(depasInfos, OrgDepa{
	// 		DepaID:        depa.ID,
	// 		DepaName:      depa.Name,
	// 		DepaIntroduce: depa.Introduce,
	// 	})
	// }
	response.Success(c, depaInfos)
}

func SetTime(c *gin.Context) {
	idAny, exist := c.Get(auth.OrgaIDKey)
	orgaId := idAny.(uint)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}
	// 绑定前端数据
	var time OrgaTime
	err := c.ShouldBindJSON(&time)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println("bind err : ", err)
		return
	}
	if time.Start > time.End {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "start need small than end")
		return
	}

	var orgaInfo models.OrgaLoginInfo
	orgaInfo.StartTime = time.Start
	orgaInfo.EndTime = time.End
	err = dao.GetDb().Where("id=?", orgaId).Updates(&orgaInfo).Error
	if err != nil {
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "get depa by orga error")
		return
	}
	response.Success(c, "set success")
}

func SetMaxApply(c *gin.Context) {
	idAny, exist := c.Get(auth.OrgaIDKey)
	orgaId := idAny.(uint)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}
	// 绑定前端数据
	var apply OrgaMaxApply
	err := c.ShouldBindJSON(&apply)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println("bind err : ", err)
		return
	}

	var orgaInfo models.OrgaLoginInfo
	orgaInfo.MaxApply = apply.MaxApply
	err = dao.GetDb().Where("id=?", orgaId).Updates(&orgaInfo).Error
	if err != nil {
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "get depa by orga error")
		return
	}
	response.Success(c, "set success")
}
