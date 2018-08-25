package controllers

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"fmt"
)

type CacheController struct {
	//beego.Controller
	BaseController
}

func (self *CacheController) Edit1() {
	bm, err := cache.NewCache("memory", `{"interval":60}`)
	if err != nil {
		fmt.Println(err)
	}
	bm.Put("astaxie", 12222, 100000)
	//bm.Get("astaxie")
	self.Data["errStr"] = bm.Get("astaxie")
	self.TplName = "cache/index.tpl"
}

func (self *CacheController) Get() {


	redis, err := cache.NewCache("redis", `{"conn":"127.0.0.1:6379", "key":"beecacheRedis""dbNum":"0"}`)
	if err != nil {
	    fmt.Println(err)
	}
	redis.Put("beecacheRedis","sdsdsdsdsds",1000000)
	//self.Ctx.WriteString(redis.Get("beecacheRedis"))
	
	self.Data["errStr"] = redis.Get("beecacheRedis")
	self.TplName = "cache/index.tpl"
}