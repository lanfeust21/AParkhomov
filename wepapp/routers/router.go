package routers

import (
	"github.com/lanfeust21/AParkhomov/wepapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/isotopes", &controllers.IsotopesController{})
	beego.Router("/isotopes/element", &controllers.IsotopesController{}, "get:Autocomplete")
	beego.Router("/transmutations22", &controllers.Transmutation22Controller{})
	beego.Router("/transmutations22/element1", &controllers.Transmutation22Controller{}, "get:Autocomplete1")
	beego.Router("/transmutations22/element2", &controllers.Transmutation22Controller{}, "get:Autocomplete2")
	beego.Router("/transmutations22/element3", &controllers.Transmutation22Controller{}, "get:Autocomplete3")
	beego.Router("/transmutations22/element4", &controllers.Transmutation22Controller{}, "get:Autocomplete4")

	//beego.Router("/fusionsfissions", &controllers.IsotopesController{})
	//beego.Router("/fusions", &controllers.IsotopesController{})
}
