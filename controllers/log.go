package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
	"path/filepath"
    "os"
    "strings"
)

type LogController struct {
	//beego.Controller
    BaseController
}

func (self *LogController) Get() {

    config := make(map[string]interface{})
    var str = getCurrentDirectory()
    str += "\\static\\logs\\collect.log"
    config["filename"] = str
    config["level"] = logs.LevelDebug

    configStr, err := json.Marshal(config)
    if err != nil {
        fmt.Println("marshal failed, err:", err)
        return
    }
    logs.SetLogger(logs.AdapterFile, string(configStr))
    logs.Debug("this is a test, my name is %s", "stu01")
    logs.Trace("this is a trace, my name is %s", "stu02")
    logs.Warn("this is a warn, my name is %s", "stu03")
    self.Data["str"] = str
    self.TplName = "log/index.tpl"
}

/*
获取程序运行路径
*/
func getCurrentDirectory() string {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        beego.Debug(err)
    }
    return strings.Replace(dir, "\\\\", "/", -1)
}