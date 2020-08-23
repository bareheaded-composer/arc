package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)


type Filer interface {
	GetFileName() string
}

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
	Code     string `gorm:"type:text"json:"code"`
	Comment string `json:"comment"`
}

func (u *Unit)GetFileName() string{
	return fmt.Sprintf("%s.md",u.Topic)
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
