package controllers

import (
	//"github.com/astaxie/beego"
	"log"
    "fmt"
    "os"
    "io/ioutil"
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

func (self *FileController) List(){
    self.Data["pageTitle"] = "文件操作"
    self.display()
}
//创建文件
func (self *FileController) Create(){

      //在当前目录下，创建一个以test为前缀的临时文件夹，并返回文件夹路径
    name, _ := ioutil.TempDir("./", "test");
    fmt.Println(name);
    
    //创建文件
    file, error := ioutil.TempFile("static/txtfile", "tmp")
    //文件关闭
    defer file.Close()
    if error != nil {
       fmt.Println("创建文件失败")
    }
    self.Data["json"]=map[string]interface{}{"data":file.Name()};
    self.ServeJSON()
}
//写入文件
func (self *FileController) FileWrite(){
      confPath := self.GetString("path")
    info:=self.GetString("info")
    content,err := parseWriteConfig(confPath,info)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(content)
    self.Data["json"]=map[string]interface{}{"data":string(content)};
    self.ServeJSON();
}

//写入text文件内容
func parseWriteConfig(confPath,info string) ([]byte,error) {
    fl, err := os.OpenFile(confPath, os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("打开文件失败")
    }
    defer fl.Close()
    byteinfo:=[]byte (info)
    n, err := fl.Write(byteinfo)
    if err == nil && n < len(byteinfo) {
        fmt.Println("写入失败")
        fmt.Println(err)
    }
    return byteinfo, err
}
func (self *FileController) FileRead() {
    confPath := self.GetString("path")
    fmt.Println("文件的地址:")
    fmt.Println(confPath)
    content,err := ReadFile(confPath)
    if err != nil {
        self.Data["data"]="";
        fmt.Println(err)
    } else{
        self.Data["data"]=content;
    }
    fmt.Println(content)
    self.Data["json"]=map[string]interface{}{"data":content};
    self.ServeJSON()
    self.StopRun()
}

//解析text文件内容
func ReadFile(path string) (str string, err error) {
    //打开文件的路径
    fi, err := os.Open(path)
    if err!=nil{
        fmt.Println("打开文件失败")
        fmt.Println(err)
    }
    defer fi.Close()
    //读取文件的内容
    fd, err := ioutil.ReadAll(fi)
    if err!=nil{
        fmt.Println("读取文件失败")
        fmt.Println(err)
    }
    str = string(fd)
    return str,err
}

func (self *FileController) FileDelete(){
    isdel :=false;
    file := self.GetString("path");           //源文件路径
    err := os.Remove(file)               //删除文件
    if err != nil {
        //删除失败,输出错误详细信息
        fmt.Println(err)
    }else {
        //如果删除成功则输出
        isdel=true
    }
    self.Data["json"]=map[string]interface{}{"data":isdel};
    self.ServeJSON()
    self.StopRun()
}
