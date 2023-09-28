package main

import (
	"log"
	"uapply/dao"
	"uapply/models"
	"uapply/routers"
	"uapply/utils/logger"
)

//	ww的测试host	spm.shrewdboy.top

// @title			NovelNock
// @version		1.0
// @description	SPM-B2
// @host			localhost:8888
// @BasePath		/bms
func main() {
	dao.InitMysql()

	dao.GetDb().AutoMigrate(&models.UserLoginInfo{})
	dao.GetDb().AutoMigrate(&models.UserCV{})
	dao.GetDb().AutoMigrate(&models.OrgaLoginInfo{})
	dao.GetDb().AutoMigrate(&models.DepaLoginInfo{})
	dao.GetDb().AutoMigrate(&models.InteLoginInfo{})

	// 日志配置
	err := logger.Init("dev")
	if err != nil {
		log.Printf("%+v \n", err)
	}

	r := routers.SetRouter()
	rErr := r.Run(":" + "8888")
	if rErr != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
