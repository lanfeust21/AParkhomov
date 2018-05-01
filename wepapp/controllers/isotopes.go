package controllers

import (
	"github.com/lanfeust21/AParkhomov/wepapp/models"
	"net/http"
)

type IsotopesController struct {
	BaseController
}

func (c *IsotopesController) Get() {

	isotopes, err := models.GetIsotopes()
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	c.Data["Isotopes"] = isotopes
	c.TplName = "isotopes.tpl"
}


func (c *IsotopesController) Post() {
	element := c.GetString("element")
	A := c.GetString("A")
	Z := c.GetString("Z")
	isotopes, err := models.GetIsotopesFrom(element,A,Z)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	c.Data["Isotopes"] = isotopes
	c.TplName = "isotopes.tpl"
}
