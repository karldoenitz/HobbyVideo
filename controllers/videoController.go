package controllers

import (
	"github.com/astaxie/beego"
	"HobbyVideo/models"
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
	videos := models.GetVideosInFront()
	c.TplName = "frontPage.html"
}

type SearchController struct {
	beego.Controller
}

func (c *SearchController) Get()  {

}
