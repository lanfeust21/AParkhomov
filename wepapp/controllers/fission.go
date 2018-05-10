package controllers

import (
	"github.com/lanfeust21/AParkhomov/wepapp/models"
	"net/http"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

type FissionController struct {
	BaseController
}

func (c *FissionController) Get() {

	element1 := c.GetString("element1")
	A1 := c.GetString("A1")
	Z1 := c.GetString("Z1")
	element2 := c.GetString("element2")
	A2 := c.GetString("A2")
	Z2 := c.GetString("Z2")
	element := c.GetString("element")
	A := c.GetString("A")
	Z := c.GetString("Z")

	count, err := models.GetFissionsCount(element1, A1, Z1, element2, A2, Z2, element, A, Z)
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

	Fissions, err := models.GetAllFissions(paginator.Offset(), dataPerPage)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	c.Data["Fissions"] = Fissions
	c.Data["paginator"] = paginator
	c.Data["Limit"] = limit
	c.TplName = "fissions.tpl"
}

func (c *FissionController) Post() {

	element1 := c.GetString("element1")
	A1 := c.GetString("A1")
	Z1 := c.GetString("Z1")
	element2 := c.GetString("element2")
	A2 := c.GetString("A2")
	Z2 := c.GetString("Z2")
	element := c.GetString("element")
	A := c.GetString("A")
	Z := c.GetString("Z")

	count, err := models.GetFissionsCount(element1, A1, Z1, element2, A2, Z2, element, A, Z)
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

	Fissions, err := models.GetFissions(element1, A1, Z1, element2, A2, Z2, element, A, Z, paginator.Offset(), dataPerPage)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	c.Data["Fissions"] = Fissions
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
	c.TplName = "fissions.tpl"
}

func (c *FissionController) Autocomplete1() {
	element := c.GetString("query")
	elements, err := models.GetFissionElement1s(element)
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

func (c *FissionController) Autocomplete2() {
	element := c.GetString("query")
	elements, err := models.GetFissionElement2s(element)
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

func (c *FissionController) Autocomplete() {
	element := c.GetString("query")
	elements, err := models.GetFissionElements(element)
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
