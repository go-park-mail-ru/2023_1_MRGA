package dataStruct

type User struct {
	Id          uint   `sql:"unique;type:uuid;primary_key;default:" json:"userId" gorm:"primaryKey;unique"`
	Username    string `json:"username" gorm:"unique"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"password"`
	Age         int    `json:"age"`
	Sex         string `json:"sex"`
	City        uint   `json:"city" gorm:"foreignKey"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
}

type City struct {
	CityId uint   `json:"cityId"`
	Name   string `json:"name"`
}
