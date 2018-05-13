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
	c.Data["Count"] = len(isotopes)
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
	c.Data["Count"] = len(isotopes)
	c.Data["Isotopes"] = isotopes
	c.Data["Element"] = element
	c.Data["A"] = A
	c.Data["Z"] = Z
	c.TplName = "isotopes.tpl"
}


/*
{
"suggestions": [
	{ "value": "United Arab Emirates", "data": "AE" },
	{ "value": "United Kingdom",       "data": "UK" },
	{ "value": "United States",        "data": "US" }
]
}
*/

func (c *IsotopesController) Autocomplete() {
	element := c.GetString("query")
	elements, err := models.GetIsotopeElements(element)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	json := map[string][]map[string]string{}
	data := []map[string]string{}
	for _,element := range elements {
		elem := map[string]string{}
		elem["value"]= element
		elem["data"]= element
		data = append(data,elem)
	}
	json["suggestions"]= data

	c.Data["json"] =json
	c.Controller.ServeJSON()
}
