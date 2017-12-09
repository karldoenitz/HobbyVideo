package controllers

import (
	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

func (c *VideoController) Get() {
	c.TplName = "videoPage.html"
}

type ListController struct {
	beego.Controller
}

func (c *ListController) Get()  {

	c.TplName = "frontPage.html"
}
