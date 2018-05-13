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
	beego.Router("/transmutations22/csv", &controllers.Transmutation22Controller{}, "get:Csv")
	beego.Router("/transmutations22/element1", &controllers.Transmutation22Controller{}, "get:Autocomplete1")
	beego.Router("/transmutations22/element2", &controllers.Transmutation22Controller{}, "get:Autocomplete2")
	beego.Router("/transmutations22/element3", &controllers.Transmutation22Controller{}, "get:Autocomplete3")
	beego.Router("/transmutations22/element4", &controllers.Transmutation22Controller{}, "get:Autocomplete4")

	beego.Router("/fusions", &controllers.FusionController{})
	beego.Router("/fusions/csv", &controllers.FusionController{}, "get:Csv")
	beego.Router("/fusions/element1", &controllers.FusionController{}, "get:Autocomplete1")
	beego.Router("/fusions/element2", &controllers.FusionController{}, "get:Autocomplete2")
	beego.Router("/fusions/element", &controllers.FusionController{}, "get:Autocomplete")

	beego.Router("/fissions", &controllers.FissionController{})
	beego.Router("/fissions/csv", &controllers.FissionController{}, "get:Csv")
	beego.Router("/fissions/element1", &controllers.FissionController{}, "get:Autocomplete1")
	beego.Router("/fissions/element2", &controllers.FissionController{}, "get:Autocomplete2")
	beego.Router("/fissions/element", &controllers.FissionController{}, "get:Autocomplete")
}
