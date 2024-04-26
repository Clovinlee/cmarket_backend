package models

type UserModel struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	EmailConfirm bool   `json:"email_confirm"`
}

func (UserModel) TableName() string {
	return "users"
}
