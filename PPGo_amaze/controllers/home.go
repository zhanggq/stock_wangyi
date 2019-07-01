package controllers

type HomeController struct {
	BaseController
}

func (self *HomeController) Index() {
	self.Data["pageTitle"] = "系统首页"
	//self.display()
	self.TplName = "public/main.html"
}

func (self *HomeController) Start() {
	self.Data["pageTitle"] = "控制面板"
	self.display()
}
