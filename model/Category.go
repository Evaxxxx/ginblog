package model

// Category 分类结构体
type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}
