package controllers

import (
	"bytes"
	//"fmt"
	"io/ioutil"
	"net/http"
)

type HttpController struct {
	BaseController
}

func (self *HttpController) Index() {
	self.Data["pageTitle"] = "post json数据"
	self.display("http/index")
}

func (self *HttpController) Post() {
	//self.Ctx.WriteString(self.GetString("url"))
	//self.Ctx.WriteString(self.GetString("data"))

	url := self.GetString("url")
	post := self.GetString("data")
	var jsonStr = []byte(post)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	self.Data["json"] = string(body)
	self.ServeJSON()
	self.StopRun()

}
