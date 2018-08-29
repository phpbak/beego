package controllers

import (
	"hello/models"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/astaxie/beego/cache/memcache"
	"fmt"
	//"github.com/garyburd/redigo/redis"
)

type CacheController struct {
	//beego.Controller
	BaseController
}

//memory,需要加载"github.com/astaxie/beego/cache"

func (self *CacheController) Memory() {

	bm, err := cache.NewCache("memory", `{"interval":60}`)
	if err != nil {
		fmt.Println(err)
	}
	bm.Put("astaxie", "this is value", 100000)
	//bm.Get("astaxie")
	self.Data["errStr"] = bm.Get("astaxie")
	self.TplName = "cache/index.tpl"
}

//redis,需要加载_ "github.com/astaxie/beego/cache/redis"

func (self *CacheController) Redis() {

	redis1, err := cache.NewCache("redis", `{"conn":"127.0.0.1:6379", "key":"beecacheRedis""dbNum":"0"}`)
	if err != nil {
	    fmt.Println(err)
	}
	redis1.Put("Redis1","sdsdsdsdsds",1000000)
	//self.Ctx.WriteString(redis.Get("beecacheRedis"))
	getkey := redis1.Get("Redis1")
	fmt.Println(getkey)
	self.Data["errStr"] = string(getkey.([]byte))
	self.TplName = "cache/index.tpl"
}
/*
//redisgo,需要加载"github.com/garyburd/redigo/redis"

func (self *CacheController) Redisgo(){

	c,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
	 	fmt.Println("Connect to redis error",err)
	}
	defer c.Close()
	//写入数据
	_,err = c.Do("SET","mykey","this is value")

	if err != nil {
		fmt.Println("SET to redis error",err)
	}
	//读取数据
	getKey,err :=redis.String(c.Do("GET","mykey"))
	if err != nil {
		fmt.Println("redis get failed:",err)
	}else{
		fmt.Printf("redis get failed: %v \n",getKey)
	}
	self.Data["errStr"] = getKey
	//self.TplName = "cache/index.tpl"
	self.display("cache/index")
}
*/
//memcahe,需要加载  
//go get "github.com/bradfitz/gomemcache/memcache"
//"github.com/astaxie/beego/cache"
//_ "github.com/astaxie/beego/cache/memcache"

func (self *CacheController) Memcache(){

	cache1, err := cache.NewCache("memcache", `{"conn":"127.0.0.1:11211"}`)
	if err != nil {
	    fmt.Println(err)
	}
	cache1.Put("keys",89,1000000)
	getKey := cache1.Get("keys")
    fmt.Println(string(getKey.([]byte)))
    fmt.Printf("%T\n", getKey)
	//self.Data["errStr"] = string(getKey.([]byte))
	//self.TplName = "cache/index.tpl"
	self.display("cache/index")
	//self.Ctx.WriteString(string(getKey.([]byte)))
}

//file,需要加载"github.com/astaxie/beego/cache"

func (self *CacheController) File() {

	cache, err := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`)
	if err != nil {
	    fmt.Println(err)
	}

	cache.Put("keys","sdsdsdsdsds",1000)
	getKey := cache.Get("keys")
   fmt.Println(getKey)
	self.Data["errStr"] = getKey
	//self.TplName = "cache/index.tpl"
	self.display("cache/index")
}

//ModelsRedis 需要加载"hello/models"

func (self *CacheController) ModelsRedis() {
	//连接redis
	models.ConnectRedis("127.0.0.1:6379")
	//字符串 url
	sUrl := "https://movie.douban.com/subject/26985127/"
	//现将它加入redis的队列
	 models.PutinQueue(sUrl)
     //获取队列长度
	length := models.GetQueueLength()
	if length > 0 {
			
     }
     //判断电影的url是否被访问
		if models.IsVisit(sUrl) {
			
}
	//消费者从队列获取Url
     sUrl = models.PopfromQueue()

     //sUrl应当记录到访问set中,表明已经访问过了
	models.AddToSet(sUrl)

	//set
	models.SetRedis("rediss","123456")
	//get
	sUrl = models.GetRedis("rediss")

    self.Data["errStr"] = sUrl
	self.display("cache/index")
}
