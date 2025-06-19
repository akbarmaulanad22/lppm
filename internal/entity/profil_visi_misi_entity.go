package entity

type ProfilVisiMisi struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*ProfilVisiMisi) TableName() string {
	return "lppm_profile_visimisi"
}
