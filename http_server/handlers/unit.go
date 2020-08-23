package handlers

import (
	"arc/model"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)


func GetAllUnits(c *gin.Context) {
	uints, err := GlobalTableManager.GetAllUnits()
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "操作成功。", "data": uints})
}


func InsertUnits(c *gin.Context) {
	var formOfInsertUnits []*model.FormOfInsertUnit
	if err := c.ShouldBindJSON(&formOfInsertUnits); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	for _,formOfInsertUnit := range formOfInsertUnits{
		if err := GlobalTableManager.InsertIfNotExist(&model.Unit{
			Solution:formOfInsertUnit.Solution,
			Link:formOfInsertUnit.Link,
			MainBody:formOfInsertUnit.MainBody,
			Purpose:formOfInsertUnit.Purpose,
			Code:formOfInsertUnit.Code,
			Topic:formOfInsertUnit.Topic,
			Comment:formOfInsertUnit.Comment,
		}); err != nil {
			logs.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"msg": "操作成功。"})
}

func GetUnitsBySolution(c *gin.Context) {
	solution := c.Param("solution")
	uints, err := GlobalTableManager.GetUnitsBySolution(solution)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "操作成功。", "data": uints})
}

func GetUnitsByMainBody(c *gin.Context) {
	mainBody := c.Param("main_body")
	uints, err := GlobalTableManager.GetUnitsByMainBody(mainBody)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "操作成功。", "data": uints})
}
