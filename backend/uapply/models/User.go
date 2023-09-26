package models

type UserLoginInfo struct {
	ID       uint   `json:"id"`
	Account  string `json:"account" binding:"required" gorm:"index:index_account,unique"`
	Password string `json:"password"`
}

type UserCV struct {
	UserID   uint   `json:"user_id"`
	Name     string `json:"name" binding:"required"`
	Age      uint16 `json:"age" binding:"required"`
	Sex      uint8  `json:"sex" binding:"required"`
	Major    string `json:"major" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Interest string `json:"interest" binding:"required"`
	IsInit   int    `json:"init"`
	Extra    string `json:"extra"`
}
