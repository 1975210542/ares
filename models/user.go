package models

import (
	"fmt"
	"time"
)

type User struct {
	Id       int    `gorm:"primary_key"`
	UserName string `form:"username" json:"username" binding:"required"`
	PassWord string `form:"passwd" json:"passwd" bdinding:"required"`
	BaseModel
}

func init() {
	SetMigrate(&User{})
	//		db.AutoMigrate(&User{})
}
func NewUser() *User {

	user := new(User)
	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	user.CreateAt = time
	user.UpdateAt = time
	return user
}

func AddUser(u *User) {
	err := db.Create(u).Error
	fmt.Println("Add err:", err)
}

func FindUserByUserName(username string) *User {
	user := User{}
	err := db.Where("user_name = ?", username).Find(&user).Error

	fmt.Println("err:", err)
	return &user

}

func UpdataUser(u *User) {

	err := db.Save(&u).Error
	fmt.Println("err:", err)
}
