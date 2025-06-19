package entity

type ProfilSODT struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*ProfilSODT) TableName() string {
	return "lppm_profil_sodt"
}
