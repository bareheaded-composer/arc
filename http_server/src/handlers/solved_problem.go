package handlers

import (
	"arc/src/env"
	"arc/src/model"
	"arc/src/utils"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertSolvedProblems(c *gin.Context) {
	var formOfInsertSolvedProblems []*model.FormOfInsertSolvedProblem
	if err := c.ShouldBindJSON(&formOfInsertSolvedProblems); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	for _, formOfInsertSolvedProblem := range formOfInsertSolvedProblems {
		if err := GlobalTableManager.InsertIfNotExist(&model.SolvedProblem{
			Solution:   formOfInsertSolvedProblem.Solution,
			Link:       formOfInsertSolvedProblem.Link,
			MainBody:   formOfInsertSolvedProblem.MainBody,
			Purpose:    formOfInsertSolvedProblem.Purpose,
			Code:       formOfInsertSolvedProblem.Code,
			Topic:      formOfInsertSolvedProblem.Topic,
			Comment:    formOfInsertSolvedProblem.Comment,
			Complexity: formOfInsertSolvedProblem.Complexity,
		}); err != nil {
			logs.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"msg": "操作成功。"})
}

func GetSolvedProblems(c *gin.Context) {
	parameter := &model.GetSolvedProblemsParameter{
		MainBody: c.Query("main_body"),
		Solution: c.Query("solution"),
	}
	solvedProblems, err := GlobalTableManager.GetSolvedProblems(parameter)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	logs.Debug(utils.MakeSolvedProblemToFile(
		env.Conf.Path.PathOfStoringSolvedProblemTextDirPath,
		env.Conf.Path.PathOfSolvedProblemText,
		solvedProblems...,
	))
	c.JSON(http.StatusOK, gin.H{"msg": "操作成功。", "data": solvedProblems})
}
