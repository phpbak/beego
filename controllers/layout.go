package controllers

import(
	
)

type LayoutController struct {
	BaseController
}

func (self *LayoutController) Index() {
	//布局页面
	self.Layout   = "common/layout.tpl"
	//内容页面
	self.TplName  = "layout/index.tpl"
	//其他的部分
    self.LayoutSections = make(map[string]string)
    //页面使用的css部分
    self.LayoutSections["HtmlHead"] = "common/head.tpl"
    //页面使用的js部分
    self.LayoutSections["Scripts"] = "common/scripts.tpl"
    //页面的补充部分，可做为底部的版权部分时候
    self.LayoutSections["SideBar"] = "common/sidebar.tpl"
}
