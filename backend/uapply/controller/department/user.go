package depa

import (
	"net/http"
	"uapply/dao"
	"uapply/models"
	"uapply/utils/auth"
	"uapply/utils/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DepaUserRegRsp struct {
	Data []DepaUserRegInfo `json:"data"`
}

type DepaUserRegInfo struct {
	UserCv           models.UserCV           `json:"user_cv"`
	UserRegistration models.UserRegistration `json:"user_registration"`
}

func GetRegUsers(c *gin.Context) {
	idAny, exist := c.Get(auth.DepaIDKey)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}
	depaId := idAny.(uint)
	var rsp DepaUserRegRsp
	var rInfos []models.UserRegistration
	err := dao.GetDb().Where("depa_id=?", depaId).Find(&rInfos).Error
	if err == gorm.ErrRecordNotFound {
		response.Success(c, rsp)
		return
	} else if err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeSystemBusy, response.CodeSystemBusy.Msg())
		return
	}

	var userIDs []uint
	userIDsMap := make(map[uint]models.UserRegistration)
	for _, rInfo := range rInfos {
		userIDs = append(userIDs, rInfo.UserID)
		userIDsMap[rInfo.UserID] = rInfo
	}

	var cvInfos []models.UserCV
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
		rsp.Data = append(rsp.Data, DepaUserRegInfo{
			UserCv:           cvInfo,
			UserRegistration: rInfo,
		})
	}
	response.Success(c, rsp)
}
