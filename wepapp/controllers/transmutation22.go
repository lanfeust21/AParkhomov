package controllers

import (
	"github.com/lanfeust21/AParkhomov/wepapp/models"
	"net/http"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

type Transmutation22Controller struct {
	BaseController
}

func (c *Transmutation22Controller) Get() {

	count, err := models.GetTransmutationsCount()
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	dataPerPage := 20
	// sets this.Data["paginator"] with the current offset (from the url query param)
	if len(c.GetString("limit")) > 0 {
		dataPerPage, err = strconv.Atoi(c.GetString("limit"))
		if err != nil {
			c.Error(http.StatusInternalServerError, err)
			return
		}
	}

	paginator := pagination.SetPaginator(c.Ctx, dataPerPage, count)

	transmutations, err := models.GetAllTransmutations(paginator.Offset(), dataPerPage)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	c.Data["Transmutations"] = transmutations
	c.Data["paginator"] = paginator
	c.TplName = "transmutation22.tpl"
}

func (c *Transmutation22Controller) Post() {
	count, err := models.GetTransmutationsCount()
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	dataPerPage := 20
	// sets this.Data["paginator"] with the current offset (from the url query param)
	if len(c.GetString("limit")) > 0 {
		dataPerPage, err = strconv.Atoi(c.GetString("limit"))
		if err != nil {
			c.Error(http.StatusInternalServerError, err)
			return
		}
	}

	paginator := pagination.SetPaginator(c.Ctx, dataPerPage, count)

	element1 := c.GetString("element1")
	A1 := c.GetString("A1")
	Z1 := c.GetString("Z1")
	element2 := c.GetString("element2")
	A2 := c.GetString("A2")
	Z2 := c.GetString("Z2")
	element3 := c.GetString("element3")
	A3 := c.GetString("A3")
	Z3 := c.GetString("Z3")
	element4 := c.GetString("element4")
	A4 := c.GetString("A4")
	Z4 := c.GetString("Z4")

	transmutations, err := models.GetTransmutations(element1, A1, Z1, element2, A2, Z2, element3, A3, Z3, element4, A4, Z4, paginator.Offset(), dataPerPage)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	c.Data["Transmutations"] = transmutations
	c.Data["paginator"] = paginator
	c.TplName = "transmutation22.tpl"
}

func (c *Transmutation22Controller) Autocomplete1() {
	element := c.GetString("query")
	elements, err := models.GetTransmutationElement1s(element)
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

func (c *Transmutation22Controller) Autocomplete2() {
	element := c.GetString("query")
	elements, err := models.GetTransmutationElement2s(element)
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

func (c *Transmutation22Controller) Autocomplete3() {
	element := c.GetString("query")
	elements, err := models.GetTransmutationElement3s(element)
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

func (c *Transmutation22Controller) Autocomplete4() {
	element := c.GetString("query")
	elements, err := models.GetTransmutationElement4s(element)
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
