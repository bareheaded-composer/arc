package model


type FormOfInsertSolvedProblem struct {
	Topic    string `json:"topic"`
	Link     string `json:"link"`
	MainBody string `json:"main_body"`
	Solution string `json:"solution"`
	Purpose  string `json:"purpose"`
	Code     string `json:"code"`
	Comment string `json:"comment"`
	Complexity string `json:"complexity"`
}
