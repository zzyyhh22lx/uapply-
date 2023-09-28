package user

import (
	"log"
	"net/http"
	"strconv"
	"uapply/dao"
	"uapply/models"
	"uapply/utils/response"

	"github.com/gin-gonic/gin"
)

type UserOrgaDepaRsp struct {
	UserOrgaDepaInfo []UserOrgaDepa `json:"all_orga_depa_infos"`
}

type UserOrgaDepa struct {
	OrgaID        uint       `json:"orga_id"`
	OrgaName      string     `json:"orga_name"`
	OrgaIntrodece string     `json:"orga_intro"`
	StartTime     int64      `json:"start_time"`
	EndTime       int64      `json:"end_time"`
	DepaInfo      []UserDepa `json:"depa_info"`
}

type UserDepa struct {
	DepaID        uint   `json:"depa_id"`
	DepaName      string `json:"depa_name"`
	DepaIntroduce string `json:"depa_intro"`
}

func GetAllOrgaDepa(c *gin.Context) {
	var orgaInfos []models.OrgaLoginInfo
	err := dao.GetDb().Find(&orgaInfos).Error
	if err != nil {
		log.Println("Get all orga_depa DataBase err : ", err)
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "database insert error")
		return
	}
	if len(orgaInfos) == 0 {
		response.Success(c, "暂无组织和部门数据")
		return
	}
	var odInfos []UserOrgaDepa
	for _, orga := range orgaInfos {
		var depaInfos []models.DepaLoginInfo
		err := dao.GetDb().Where("orga_id=?", orga.ID).Find(&depaInfos).Error
		if err != nil {
			log.Println("Get all depa DataBase err : ", err)
			continue
		}
		uod := UserOrgaDepa{
			OrgaID:        orga.ID,
			OrgaName:      orga.Name,
			OrgaIntrodece: orga.Introduce,
			StartTime:     orga.StartTime,
			EndTime:       orga.EndTime,
		}
		var dInfos []UserDepa
		for _, depa := range depaInfos {
			dInfos = append(dInfos, UserDepa{
				DepaID:        depa.ID,
				DepaName:      depa.Name,
				DepaIntroduce: depa.Introduce,
			})
		}
		uod.DepaInfo = dInfos
		odInfos = append(odInfos, uod)
	}
	odr := UserOrgaDepaRsp{
		UserOrgaDepaInfo: odInfos,
	}
	response.Success(c, odr)
}

func GetAllDepaByOrga(c *gin.Context) {
	var depasInfos []UserDepa
	orgaIdStr := c.Query("orga_id")
	orgaId, err := strconv.Atoi(orgaIdStr)
	if err != nil {
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "atoi error")
		return
	}
	var depaInfos []models.DepaLoginInfo
	err = dao.GetDb().Where("orga_id=?", orgaId).Find(&depaInfos).Error
	if err != nil {
		response.FailWithMsg(c, http.StatusServiceUnavailable, response.CodeSystemBusy, "get depa by orga error")
		return
	}
	for _, depa := range depaInfos {
		depasInfos = append(depasInfos, UserDepa{
			DepaID:        depa.ID,
			DepaName:      depa.Name,
			DepaIntroduce: depa.Introduce,
		})
	}
	response.Success(c, depasInfos)
}
