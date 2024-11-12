package models

type Product struct {
	Id           int     `gorm:"primaryKey"`
	Name         string  `gorm:"type:varchar(300)" json:"name"`
	Price_per_kg float64 `gorm:"type:float" json:"price_per_kg"`
	Desc         string  `gorm:"type:varchar(300)" json:"description"`
}

