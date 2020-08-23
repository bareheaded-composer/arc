package utils

import (
	"arc/src/model"
	"bytes"
	"fmt"
	"github.com/astaxie/beego/logs"
	"html/template"
	"io/ioutil"
)

func MakeFilerContentToFile(targetFileDir string, templateFilePath string, filer model.Filer) error {
	contentBuffer, err := GetContentBuffer(templateFilePath, filer)
	if err != nil {
		logs.Error(err)
		return err
	}

	if err := ioutil.WriteFile(
		fmt.Sprintf("%s/%s", targetFileDir,filer.GetFileName(), ),
		contentBuffer.Bytes(),
		0777,
		);
		err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func GetContent(templateFilePath string, readerData interface{}) (string, error) {
	contentBuffer, err := GetContentBuffer(templateFilePath, readerData)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return contentBuffer.String(), nil
}

func GetContentBuffer(templateFilePath string, readerData interface{}) (*bytes.Buffer, error) {
	templateText, err := ioutil.ReadFile(templateFilePath)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	unescaped := func (str string) template.HTML{
		return template.HTML(str)
	}
	templateObject := template.New("")
	templateObject.Funcs(map[string]interface{}{
		"unescaped":unescaped,
	})
	if _, err := templateObject.Parse(string(templateText)); err != nil {
		logs.Error(err)
		return nil, err
	}
	var contentBuffer bytes.Buffer
	if err := templateObject.Execute(&contentBuffer, readerData); err != nil {
		logs.Error(err)
		return nil, err
	}
	return &contentBuffer, nil
}
