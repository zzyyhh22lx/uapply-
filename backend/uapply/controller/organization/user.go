package orga

import (
	"net/http"
	"uapply/dao"
	"uapply/models"
	"uapply/utils/auth"
	"uapply/utils/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrgaRsp struct {
	Data []OrgaUserRegDepa `json:"data"`
}

type OrgaUserRegDepa struct {
	DepaID uint              `json:"depa_id"`
	Users  []OrgaUserRegInfo `json:"data"`
}

type OrgaUserRegInfo struct {
	UserCv           models.UserCV           `json:"user_cv"`
	UserRegistration models.UserRegistration `json:"user_registration"`
}

func GetRegUsers(c *gin.Context) {
	idAny, exist := c.Get(auth.OrgaIDKey)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}
	orgaID := idAny.(uint)
	var rsp OrgaRsp
	var depaInfos []models.DepaLoginInfo
	err := dao.GetDb().Where("orga_id=?", orgaID).Find(&depaInfos).Error
	if err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeSystemBusy, response.CodeSystemBusy.Msg())
		return
	}

	for _, depaInfo := range depaInfos {
		var tmp OrgaUserRegDepa
		var rInfos []models.UserRegistration
		err = dao.GetDb().Where("depa_id=?", depaInfo.ID).Find(&rInfos).Error
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
			tmp.Users = append(tmp.Users, OrgaUserRegInfo{
				UserCv:           cvInfo,
				UserRegistration: rInfo,
			})
		}
		tmp.DepaID = depaInfo.ID
		rsp.Data = append(rsp.Data, tmp)
	}
	response.Success(c, rsp)
}
