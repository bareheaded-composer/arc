package handlers

import (
	"arc/model"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)




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

func GetUnits(c *gin.Context) {
	parameter := &model.GetUnitsParameter{
		MainBody:c.Query("main_body"),
		Solution:c.Query("solution"),
	}
	uints, err := GlobalTableManager.GetUnits(parameter)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "操作成功。", "data": uints})
}

