package models

type Comment struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	UserID  uint   `gorm:"not null" json:"user_id" form:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id" form:"photo_id"`
	Message string `gorm:"not null" json:"message" form:"message"`
	GormModel
	User  User  `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
	Photo Photo `gorm:"constraint:OnDelete:CASCADE;" json:"photo"`
}
