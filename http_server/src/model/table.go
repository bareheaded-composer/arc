package model

import (
	"github.com/jinzhu/gorm"
)



type Table interface {
	TableName() string
	GetIdentity()map[string]interface{}
}

type SolvedProblem struct{
	gorm.Model
	Topic    string `json:"topic"`
	Link     string `json:"link"`
	MainBody string `json:"main_body"`
	Solution string `json:"solution"`
	Purpose  string `json:"purpose"`
	Code     string `gorm:"type:text"json:"code"`
	Comment string `json:"comment"`
	Complexity string `json:"complexity"`
}

func (*SolvedProblem) TableName() string{
	return "solved_problem"
}

func (s *SolvedProblem)  GetIdentity() map[string]interface{}{
	return map[string]interface{}{
		"link":s.Link,
		"solution":s.Solution,
		"complexity":s.Complexity,
	}
}
