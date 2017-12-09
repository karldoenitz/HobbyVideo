package routers

import (
	"HobbyVideo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/video", &controllers.VideoController{})
	beego.Router("/front", &controllers.ListController{})
}
