/**
 * Created by GOLAND.
 * User: pengyu
 * Time: 2023/10/18 11:11
 */

package model

const StatusOk uint8 = 0
const StatusDel uint8 = 1

type UserAddress struct {
	Id         int64  `json:"id" gorm:"column:id"`
	UserId     int64  `json:"userId" gorm:"column:user_id"`
	PostCode   string `json:"postCode" gorm:"column:post_code"`
	Phone      string `json:"phone" gorm:"column:phone"`
	Name       string `json:"name" gorm:"column:name"`
	Province   string `json:"province" gorm:"column:province"`
	City       string `json:"city" gorm:"column:city"`
	District   string `json:"district" gorm:"column:district"`
	Address    string `json:"address" gorm:"column:address"`
	Default    bool   `json:"default" gorm:"column:default"`
	CreateTime int64  `json:"createTime" gorm:"autoCreateTime;column:create_time"`
	Status     uint8  `json:"status" gorm:"column:status"`
}

func (ua *UserAddress) TableName() string {
	return "user_address"
}

func AddAddress(address UserAddress) error {
	var err error
	if address.Default {
		err = db.Model(&UserAddress{}).Where("user_id = ?", address.UserId).Update("default", 0).Error
	}

	if err != nil {
		return err
	}

	err = db.Create(&address).Error
	return err
}

func AddressList(userId int64) ([]UserAddress, error) {
	var addressList []UserAddress
	err := db.Where("user_id = ? and status = ?", userId, StatusOk).Find(&addressList).Error

	return addressList, err
}

func DelAddress(userId, id int64) error {
	err := db.Model(&UserAddress{}).Where("user_id = ? and id = ?", userId, id).Update("status", StatusDel).Error
	return err
}
