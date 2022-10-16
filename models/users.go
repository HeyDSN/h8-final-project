package models

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email"`
	Password string `gorm:"not null" json:"password" form:"password"`
	Age      int    `gorm:"not null" json:"age" form:"age"`
	GormModel
	SocialMedia []SocialMedia `gorm:"constraint:OnDelete:CASCADE;" json:"social_medias"`
	Comment     []Comment     `gorm:"constraint:OnDelete:CASCADE;" json:"comments"`
	Photo       []Photo       `gorm:"constraint:OnDelete:CASCADE;" json:"photos"`
}
