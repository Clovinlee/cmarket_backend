package models

type RarityModel struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Level string `json:"level"`
	Color string `json:"color"`
}

func (RarityModel) TableName() string {
	return "rarities"
}
