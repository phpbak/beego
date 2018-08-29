package controllers

import (
    "hello/models"
    "strings"
    "time"
    "github.com/astaxie/beego"
)
type DbController struct {
	//beego.Controller
	BaseController
}
//测试
func (self *DbController) Get() {
	//orm.RegisterModel(new(models.User))
	admin := new(models.Admin)
	//user.Id = 0
	admin.LoginName = "get323664890"
	if _, err := models.AdminAdd(admin); err != nil {
		self.Ctx.WriteString("fail")
	}
	self.Ctx.WriteString("success")
}
//列表
func (self *DbController) Index() {
	self.Data["pageTitle"] = "Db列表"

	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	realName := strings.TrimSpace(self.GetString("realName"))

	StatusText := make(map[int]string)
	StatusText[0] = "<font color='red'>禁用</font>"
	StatusText[1] = "正常"

	pageSize := limit
	//查询条件
	filters := make([]interface{}, 0)
	//
	if realName != "" {
		filters = append(filters, "real_name__icontains", realName)
	}
	result, count := models.AdminGetList(page, pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["login_name"] = v.LoginName
		row["real_name"] = v.RealName
		row["phone"] = v.Phone
		row["email"] = v.Email
		row["role_ids"] = v.RoleIds
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["status"] = v.Status
		row["status_text"] = StatusText[v.Status]
		list[k] = row
	}
    self.Data["list"] = list
    self.Data["count"] = count
	self.display("db/index")
}
//新增
func (self *DbController) Add() {
	self.Data["pageTitle"] = "Db新增"
	self.display()
}
//新增--写入
func (self *DbController) AddSave() {
	//获取数据
	Admin := new(models.Admin)
	Admin.LoginName = strings.TrimSpace(self.GetString("login_name"))
	Admin.RealName = strings.TrimSpace(self.GetString("real_name"))
	Admin.Phone = strings.TrimSpace(self.GetString("phone"))
	Admin.Email = strings.TrimSpace(self.GetString("email"))
	Admin.RoleIds = strings.TrimSpace(self.GetString("roleids"))
	Admin.UpdateTime = time.Now().Unix()
	//Admin.UpdateId = self.userId
	Admin.Status = 1
	// 检查登录名是否已经存在
	_, err := models.AdminGetByName(Admin.LoginName)

	if err == nil {
		self.Ctx.WriteString("登录名已经存在")
		self.StopRun() //停
	}
	//新增
	if _, err := models.AdminAdd(Admin); err != nil {
			self.Ctx.WriteString(err.Error())
			self.StopRun()//停
		}
	self.Ctx.Redirect(302,self.URLFor("DbController.Index"))
}

func (self *DbController) Edit() {
	self.Data["pageTitle"] = "Db修改"

	id, _ := self.GetInt("id", 0)
	Admin, err := models.AdminGetById(id)
	if err!=nil {
		self.Ctx.WriteString(err.Error())
		self.StopRun()//停
	}
	row := make(map[string]interface{})
	row["id"] = Admin.Id
	row["login_name"] = Admin.LoginName
	row["real_name"] = Admin.RealName
	row["phone"] = Admin.Phone
	row["email"] = Admin.Email
	row["role_ids"] = Admin.RoleIds

	self.Data["admin"] = row

	self.display()
}
//修改－－写入
func (self *DbController) EditSave() {

	Admin_id, _ := self.GetInt("id")
	Admin, _ := models.AdminGetById(Admin_id)

	Admin.Id = Admin_id
	Admin.UpdateTime = time.Now().Unix()
	//Admin.UpdateId = self.userId
	Admin.LoginName = strings.TrimSpace(self.GetString("login_name"))
	Admin.RealName = strings.TrimSpace(self.GetString("real_name"))
	Admin.Phone = strings.TrimSpace(self.GetString("phone"))
	Admin.Email = strings.TrimSpace(self.GetString("email"))
	Admin.RoleIds = strings.TrimSpace(self.GetString("roleids"))
	Admin.UpdateTime = time.Now().Unix()
	//Admin.UpdateId = self.userId
	Admin.Status = 1

	if err := Admin.Update(); err != nil {
		self.Ctx.WriteString(err.Error())
		self.StopRun()//停
	}

	self.Ctx.Redirect(302,self.URLFor("DbController.Index"))
}
//删除
func  (self *DbController) Del() {
	Admin_id, _ := self.GetInt("id",0)
	if err := models.AdminDel(Admin_id); err != nil {
		self.Ctx.WriteString(err.Error())
		self.StopRun()//停
	}
	self.Ctx.Redirect(302,self.URLFor("DbController.Index"))
}
//分页
func (self *DbController) Pages() {

	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 5
	}

	realName := strings.TrimSpace(self.GetString("realName"))

	StatusText := make(map[int]string)
	StatusText[0] = "<font color='red'>禁用</font>"
	StatusText[1] = "正常"

	pageSize := limit
	//查询条件
	filters := make([]interface{}, 0)
	//
	if realName != "" {
		filters = append(filters, "real_name__icontains", realName)
	}
	result, count := models.AdminGetList(page, pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["login_name"] = v.LoginName
		row["real_name"] = v.RealName
		row["phone"] = v.Phone
		row["email"] = v.Email
		row["role_ids"] = v.RoleIds
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["status"] = v.Status
		row["status_text"] = StatusText[v.Status]
		list[k] = row
	}

	pages := models.Paginator(page,pageSize,count)
	self.Data["datas"] = list  //用户的数据
    self.Data["paginator"] = pages //分页的数据
    self.Data["totals"] =count //分页的数据
	self.display()
}