package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100);unique"`
	Password string `gorm:"type:varchar(50)"`
}
