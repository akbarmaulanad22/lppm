package entity

// PenelitianRDRP adalah entity untuk Restra dan Roadmap Penelitian
type PenelitianRDRP struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PenelitianRDRP) TableName() string {
	return "lppm_penelitian_rdrp"
}
