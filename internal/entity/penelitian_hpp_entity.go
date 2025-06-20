package entity

type PenelitianHPP struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PenelitianHPP) TableName() string {
	return "lppm_penelitian_hpp"
} 