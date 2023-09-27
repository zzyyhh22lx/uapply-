package models

type OrgaLoginInfo struct {
	ID        uint   `json:"id"`
	Account   string `json:"account" binding:"required" gorm:"index:index_account,unique"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"name"`
	Introduce string `json:"introduce"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

type DepaLoginInfo struct {
	ID        uint   `json:"id"`
	Account   string `json:"account" binding:"required" gorm:"index:index_account,unique"`
	Password  string `json:"password" binding:"required"`
	OrgaID    uint   `json:"orga_id" gorm:"index:index_orga"`
	Name      string `json:"name"`
	Introduce string `json:"introduce"`
}
