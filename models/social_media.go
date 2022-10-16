package models

type SocialMedia struct {
	ID             uint   `gorm:"primary_key" json:"id"`
	Name           string `gorm:"not null" json:"name" form:"name"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url"`
	GormModel
	UserID uint `gorm:"not null" json:"user_id" form:"user_id"`
	User   User `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
}
