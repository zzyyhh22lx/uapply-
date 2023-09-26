package user

import (
	"log"
	"net/http"
	"regexp"
	"uapply/dao"
	"uapply/models"
	"uapply/utils/auth"
	"uapply/utils/response"

	"github.com/gin-gonic/gin"
)

func validatePhoneNumber(phoneNumber string) bool {
	// 正则表达式验证手机号码
	regex := `^1[3456789]\d{9}$`
	match, _ := regexp.MatchString(regex, phoneNumber)
	return match
}

func validateEmailAddress(emailAddress string) bool {
	// 正则表达式验证邮箱号码
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, emailAddress)
	return match
}

func CommitCv(c *gin.Context) {
	idAny, exist := c.Get(auth.UserIDKey)
	id := idAny.(uint)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}
	// 绑定前端数据
	var cv models.UserCV
	err := c.ShouldBindJSON(&cv)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println("bind err : ", err)
		return
	}
	if !validatePhoneNumber(cv.Phone) {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "phone number error")
		return
	}
	if !validateEmailAddress(cv.Email) {
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "email address error")
		return
	}
	cv.IsInit = 2
	err = dao.GetDb().Where("user_id=?", id).Updates(&cv).Error
	if err != nil {
		log.Println(err)
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "database insert error")
		return
	}
	response.Success(c, nil)
}

func GetCv(c *gin.Context) {
	idAny, exist := c.Get(auth.UserIDKey)
	id := idAny.(uint)
	if !exist {
		response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}
	var cv models.UserCV

	err := dao.GetDb().Where("user_id=?", id).First(&cv).Error
	if err != nil {
		log.Println(err)
		response.FailWithMsg(c, http.StatusBadRequest, response.CodeParamsInvalid, "database get cv error")
		return
	}
	response.Success(c, cv)
}
