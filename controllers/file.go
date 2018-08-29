package controllers

import (
	//"github.com/astaxie/beego"
	"log"
)

type FileController struct {
	//beego.Controller
    BaseController
}

func (self *FileController) Index() {
	//self.TplName = "file/index.tpl"
    self.display()
}

func (self *FileController) Post() {
	f, h, err := self.GetFile("uploadname")
    if err != nil {
        log.Fatal("getfile err ", err)
    }
    defer f.Close()
    self.SaveToFile("uploadname", "static/upload/" + h.Filename) // 保存位置在 static/upload, 
    self.Ctx.WriteString("files upload success")
}