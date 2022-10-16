package models

type Photo struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Title    string `gorm:"not null" json:"title" form:"title"`
	Caption  string `gorm:"null" json:"caption" form:"caption"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url"`
	UserID   uint   `gorm:"not null" json:"user_id" form:"user_id"`
	GormModel
	User User `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
}
