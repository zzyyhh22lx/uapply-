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
	"gorm.io/gorm"
)

func GetAllUser(c *gin.Context) {
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
	var regiInfo []models.UserRegistration
	err := dao.GetDb().Where("orga_id=? and depa_id=?", orgaId, depaId).Find(&regiInfo).Error
	if err != nil {
		log.Println(err)
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "find orga error")
		return
	}
	var userIDs []uint
	userIDsMap := make(map[uint]models.UserRegistration)
	for _, rInfo := range regiInfo {
		userIDs = append(userIDs, rInfo.UserID)
		userIDsMap[rInfo.UserID] = rInfo
	}

	var cvInfos []models.UserCV
	var rsp depa.DepaUserRegRsp
	err = dao.GetDb().Where("user_id in (?)", userIDs).Find(&cvInfos).Error
	if err == gorm.ErrRecordNotFound {
		response.Success(c, rsp)
		return
	} else if err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeSystemBusy, response.CodeSystemBusy.Msg())
		return
	}

	for _, cvInfo := range cvInfos {
		rInfo, exist := userIDsMap[cvInfo.UserID]
		if !exist {
			continue
		}
		rsp.Data = append(rsp.Data, depa.DepaUserRegInfo{
			UserCv:           cvInfo,
			UserRegistration: rInfo,
		})
	}
	response.Success(c, rsp)
}

type UserStepReq struct {
	UserID uint `json:"user_id" binding:"required"`
}

func UserStep(c *gin.Context) {
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
	// 绑定前端数据
	var userInfo UserStepReq
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println("bind err : ", err)
		return
	}
	var userRegi models.UserRegistration
	err = dao.GetDb().Where("orga_id=? and depa_id=? and user_id = ?", orgaId, depaId, userInfo.UserID).First(&userRegi).Error
	if err == gorm.ErrRecordNotFound {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "该用户无报名该部门")
		return
	} else if err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeSystemBusy, response.CodeSystemBusy.Msg())
		return
	}
	userRegi.Step++
	err = dao.GetDb().Where("orga_id=? and depa_id=? and user_id = ?", orgaId, depaId, userInfo.UserID).Updates(&userRegi).Error
	if err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeSystemBusy, response.CodeSystemBusy.Msg())
		return
	}
	response.Success(c, userRegi)
}
