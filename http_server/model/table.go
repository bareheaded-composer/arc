package model

import "github.com/jinzhu/gorm"


type Table interface {
	TableName() string
	GetIdentity()map[string]interface{}
}
type Unit struct{
	gorm.Model
	Topic    string `json:"topic"`
	Link     string `json:"link"`
	MainBody string `json:"main_body"`
	Solution string `json:"solution"`
	Purpose  string `json:"purpose"`
	Code     string `json:"code"`
	Comment string `json:"comment"`
}

func (*Unit) TableName() string{
	return "unit"
}

func (u *Unit)  GetIdentity() map[string]interface{}{
	return map[string]interface{}{
		"link":u.Link,
		"solution":u.Solution,
	}
}
