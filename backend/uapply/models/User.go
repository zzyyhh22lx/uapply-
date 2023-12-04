package models

type UserLoginInfo struct {
	ID       uint   `json:"id"`
	Account  string `json:"account" binding:"required" gorm:"index:index_account,unique"`
	Password string `json:"password" binding:"required"`
}

type UserCV struct {
	UserID   uint   `json:"user_id" gorm:"primary_key"`
	Name     string `json:"name" binding:"required"`
	Age      uint16 `json:"age" binding:"required"`
	Sex      uint8  `json:"sex" binding:"required"`
	Major    string `json:"major" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Interest string `json:"interest" binding:"required"`
	Wechat   string `json:"wc"`
	QQ       string `json:"qq"`
	IsInit   int    `json:"init"`
	Extra    string `json:"extra"`
}

type UserRegistration struct {
	ID     uint `json:"id"`
	UserID uint `json:"user_id" gorm:"index:index_user_id"`
	OrgaID uint `json:"orga_id" binding:"required" gorm:"index:index_orga_id"`
	DepaID uint `json:"depa_id" binding:"required" gorm:"index:index_depa_id"`
	// step 0 表示面试未进行， 1 2 3 分别表示 1 轮 2 轮 3 轮
	Step uint `json:"step"`
	// 拼接起来的分数,比如 1:3,2:4 表示第一轮 3 分，第二轮 4 分
	Score string `json:"score"`
	// 评价，拼接起来的评价,比如 1:不错,2:不行 表示第一轮评价是 不错，第二轮评价是不行
	Review string `json:"review"`
	// status 0 表示进行中，1 表示录取
	Status int `json:"status"`
}
