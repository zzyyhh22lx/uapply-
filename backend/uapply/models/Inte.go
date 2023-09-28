package models

type InteLoginInfo struct {
	ID       uint   `json:"id"`
	Account  string `json:"account" binding:"required" gorm:"index:index_account,unique"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"`
	OrgaID   uint   `json:"orga_id"`
	DepaID   uint   `json:"depa_id"`
}
