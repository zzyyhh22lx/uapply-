package inte

import (
	"log"
	"net/http"
	"strconv"
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

type InteUserReq struct {
	UserID  uint `json:"user_id" binding:"required"`
	SendSms bool `json:"send_sms"`
}

type InteScoreReq struct {
	UserID uint    `json:"user_id" binding:"required"`
	Score  float32 `json:"score" binding:"required"`
	Review string  `json:"review" binding:"required"`
	Step   uint    `json:"step" binding:"required"`
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
	var userInfo InteUserReq
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

func Pass(c *gin.Context) {
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
	var userInfo InteUserReq
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
	userRegi.Status = 1
	err = dao.GetDb().Where("orga_id=? and depa_id=? and user_id = ?", orgaId, depaId, userInfo.UserID).Updates(&userRegi).Error
	if err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeSystemBusy, response.CodeSystemBusy.Msg())
		return
	}

	// 发送短信
	if userInfo.SendSms {
		response.Success(c, "短信服务尚未开通")
		return
	}
	response.Success(c, userRegi)
}

func Out(c *gin.Context) {
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
	var userInfo InteUserReq
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
	userRegi.Status = 2
	err = dao.GetDb().Where("orga_id=? and depa_id=? and user_id = ?", orgaId, depaId, userInfo.UserID).Updates(&userRegi).Error
	if err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeSystemBusy, response.CodeSystemBusy.Msg())
		return
	}

	// 发送短信
	if userInfo.SendSms {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeSystemBusy, "短信服务尚未开通")
		return
	}
	response.Success(c, userRegi)
}

func Score(c *gin.Context) {
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
	var userInfo InteScoreReq
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
	//
	if userInfo.Step <= 0 || userInfo.Step > userRegi.Step {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsFalse, "面试阶段不合法")
		return
	}
	scoreAdd := strconv.FormatUint(uint64(userInfo.Step), 10) + ":" + strconv.FormatFloat(float64(userInfo.Score), 'f', 1, 64) + "\n"
	reviewAdd := strconv.FormatUint(uint64(userInfo.Step), 10) + ":" + userInfo.Review + "\n"
	userRegi.Score += scoreAdd
	userRegi.Review += reviewAdd
	err = dao.GetDb().Where("orga_id=? and depa_id=? and user_id = ?", orgaId, depaId, userInfo.UserID).Updates(&userRegi).Error
	if err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeSystemBusy, response.CodeSystemBusy.Msg())
		return
	}
	response.Success(c, userRegi)
}
