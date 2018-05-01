package routers

import (
	"github.com/lanfeust21/AParkhomov/wepapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
