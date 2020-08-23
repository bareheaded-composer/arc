package env

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

var Conf = &conf{}

type conf struct {
	Path        Path        `json:"path"`
	Mysql       Mysql       `json:"mysql"`
	Http        Http        `json:"http"`
}

func (c *conf) Load(pathOfConfFile string) error {
	confFileData, err := ioutil.ReadFile(pathOfConfFile)
	if err != nil {
		logs.Error(err)
		return err
	}
	if err := json.Unmarshal(confFileData, c); err != nil {
		logs.Error(err)
		return err
	}
	logs.Debug("%+v", c)
	return nil
}
