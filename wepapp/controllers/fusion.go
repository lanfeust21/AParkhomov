package controllers

import (
	"github.com/lanfeust21/AParkhomov/wepapp/models"
	"net/http"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

type FusionController struct {
	BaseController
}

func (c *FusionController) Get() {

	element1 := c.GetString("element1")
	A1 := c.GetString("A1")
	Z1 := c.GetString("Z1")
	element2 := c.GetString("element2")
	A2 := c.GetString("A2")
	Z2 := c.GetString("Z2")
	element := c.GetString("element")
	A := c.GetString("A")
	Z := c.GetString("Z")

	count, err := models.GetFusionsCount(element1, A1, Z1, element2, A2, Z2, element, A, Z)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	limit := "50"
	if len(c.GetString("limit")) > 0 {
		limit = c.GetString("limit")
	}
	dataPerPage, err := strconv.Atoi(limit)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	paginator := pagination.SetPaginator(c.Ctx, dataPerPage, count)

	Fusions, err := models.GetAllFusions(paginator.Offset(), dataPerPage)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	c.Data["Fusions"] = Fusions
	c.Data["paginator"] = paginator
	c.Data["Limit"] = limit
	c.TplName = "fusions.tpl"
}

func (c *FusionController) Post() {

	element1 := c.GetString("element1")
	A1 := c.GetString("A1")
	Z1 := c.GetString("Z1")
	element2 := c.GetString("element2")
	A2 := c.GetString("A2")
	Z2 := c.GetString("Z2")
	element := c.GetString("element")
	A := c.GetString("A")
	Z := c.GetString("Z")

	count, err := models.GetFusionsCount(element1, A1, Z1, element2, A2, Z2, element, A, Z)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	limit := "50"
	if len(c.GetString("limit")) > 0 {
		limit = c.GetString("limit")
	}
	dataPerPage, err := strconv.Atoi(limit)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	paginator := pagination.SetPaginator(c.Ctx, dataPerPage, count)

	Fusions, err := models.GetFusions(element1, A1, Z1, element2, A2, Z2, element, A, Z, paginator.Offset(), dataPerPage)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	c.Data["Fusions"] = Fusions
	c.Data["paginator"] = paginator
	c.Data["Count"] = count
	c.Data["Element1"] = element1
	c.Data["A1"] = A1
	c.Data["Z1"] = Z1
	c.Data["Element2"] = element2
	c.Data["A2"] = A2
	c.Data["Z2"] = Z2
	c.Data["Element"] = element
	c.Data["A"] = A
	c.Data["Z"] = Z
	c.Data["Limit"] = limit
	c.TplName = "fusions.tpl"
}

func (c *FusionController) Autocomplete1() {
	element := c.GetString("query")
	elements, err := models.GetFusionElement1s(element)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	json := map[string][]map[string]string{}
	data := []map[string]string{}
	for _, element := range elements {
		elem := map[string]string{}
		elem["value"] = element
		elem["data"] = element
		data = append(data, elem)
	}
	json["suggestions"] = data

	c.Data["json"] = json
	c.Controller.ServeJSON()
}

func (c *FusionController) Autocomplete2() {
	element := c.GetString("query")
	elements, err := models.GetFusionElement2s(element)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	json := map[string][]map[string]string{}
	data := []map[string]string{}
	for _, element := range elements {
		elem := map[string]string{}
		elem["value"] = element
		elem["data"] = element
		data = append(data, elem)
	}
	json["suggestions"] = data

	c.Data["json"] = json
	c.Controller.ServeJSON()
}

func (c *FusionController) Autocomplete() {
	element := c.GetString("query")
	elements, err := models.GetFusionElements(element)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	json := map[string][]map[string]string{}
	data := []map[string]string{}
	for _, element := range elements {
		elem := map[string]string{}
		elem["value"] = element
		elem["data"] = element
		data = append(data, elem)
	}
	json["suggestions"] = data

	c.Data["json"] = json
	c.Controller.ServeJSON()
}
