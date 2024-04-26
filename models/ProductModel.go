package models

type ProductModel struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	Name        string      `json:"name"`
	Description *string     `json:"description"`
	Price       float64     `json:"price"`
	Image       *string     `json:"image"`
	IDRarity    uint        `gorm:"column:id_rarity" json:"id_rarity"`
	Rarity      RarityModel `gorm:"foreignKey:IDRarity" json:"rarity"`
}

func (ProductModel) TableName() string {
	return "products"
}
