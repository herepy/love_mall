/**
 * Created by GOLAND.
 * User: pengyu
 * Time: 2023/10/12 16:41
 */

package model

import "love_mall/user/internal/types"

type User struct {
	Id         int64  `json:"id" gorm:"column:id"`
	Nickname   string `json:"nickname" gorm:"column:nickname"`
	Avatar     string `json:"avatar" gorm:"column:avatar"`
	Gender     int    `json:"gender" gorm:"column:gender"`
	Phone      string `json:"phone" gorm:"column:phone"`
	Openid     string `json:"openid" gorm:"column:openid"`
	Province   string `json:"province" gorm:"column:province"`
	City       string `json:"city" gorm:"column:city"`
	CreateTime int64  `json:"createTime" gorm:"autoCreateTime;column:create_time"`
}

func (u *User) TableName() string {
	return "user"
}

func GetUserById(id int64) (*User, error) {
	var user User
	err := db.First(&user, id).Error
	return &user, err
}

func GetUserByOpenid(openid string) (*User, error) {
	var user User
	err := db.Where("openid = ?", openid).First(&user).Error
	return &user, err
}

func GetUserByPhone(phone string) (*User, error) {
	var user User
	err := db.Where("phone = ?", phone).First(&user).Error
	return &user, err
}

func UpdateUserInfo(userId int64, info types.EditUserInfoRequest) (*User, error) {
	var user User
	err := db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}

	user.Nickname = info.Nickname
	user.Avatar = info.Avatar
	user.Gender = info.Gender
	user.Province = info.Province
	user.City = info.City
	err = db.Save(&user).Error

	return &user, err
}
