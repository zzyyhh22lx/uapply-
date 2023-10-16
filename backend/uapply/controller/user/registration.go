package user

import (
	"log"
	"net/http"
	"uapply/dao"
	"uapply/models"
	"uapply/utils/auth"
	"uapply/utils/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Registration(c *gin.Context) {
	idAny, exist := c.Get(auth.UserIDKey)
	id := idAny.(uint)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}
	var re models.UserRegistration
	err := c.ShouldBindJSON(&re)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println("bind err : ", err)
		return
	}
	re.UserID = id
	// 检查是否提交简历了
	var cv models.UserCV
	err = dao.GetDb().Where("user_id=?", re.UserID).First(&cv).Error
	if err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "database get user CV error")
		return
	}
	if cv.IsInit == 1 {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "请先提交简历再进行报名")
		return
	}
	// 检查社团是否存在
	var depa models.DepaLoginInfo
	err = dao.GetDb().Where("id=?", re.DepaID).First(&depa).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "社团不存在")
			return
		}
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "database get depa error")
		return
	}
	if depa.OrgaID != re.OrgaID || depa.ID != re.DepaID {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "社团不存在")
		return
	}

	var orga models.OrgaLoginInfo

	err = dao.GetDb().Where("id=?", re.OrgaID).First(&orga).Error
	if err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "database get orga error")
		return
	}
	var rega []models.UserRegistration
	err = dao.GetDb().Where("user_id=?", re.UserID).Find(&rega).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "database get rega error")
		return
	}
	// 检查重复报名
	for _, r := range rega {
		if r.OrgaID == re.OrgaID && r.DepaID == re.DepaID {
			response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "您已报名过该社团")
			return
		}
	}
	// 查看是否超过组织限制
	if orga.MaxApply <= len(rega) {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "超过组织最大报名社团数")
		return
	}
	// 添加报名信息
	re.Step = 0
	re.Score = ""
	re.Review = ""
	re.Status = 0
	err = dao.GetDb().Create(&re).Error
	if err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "database get rega error")
		return
	}
	response.Success(c, re)
}

func GetRegistration(c *gin.Context) {
	idAny, exist := c.Get(auth.UserIDKey)
	id := idAny.(uint)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}

	var rega []models.UserRegistration
	err := dao.GetDb().Where("user_id=?", id).Find(&rega).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "database get rega error")
		return
	}
	response.Success(c, rega)
}
