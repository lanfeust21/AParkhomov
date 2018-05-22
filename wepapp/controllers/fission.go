package controllers

import (
	"github.com/lanfeust21/AParkhomov/wepapp/models"
	"net/http"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
	"encoding/csv"
	"fmt"
)

type FissionController struct {
	BaseController
}

func (c *FissionController) Get() {

	queryString := ""
	element1 := c.GetString("element1")
	if len(element1) > 0 {
		queryString += fmt.Sprintf("%s=%s", "element1", element1)
	}
	A1 := c.GetString("A1")
	if len(A1) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "A1", A1)
	}
	Z1 := c.GetString("Z1")
	if len(Z1) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "Z1", Z1)
	}
	element2 := c.GetString("element2")
	if len(element2) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "element2", element2)
	}
	A2 := c.GetString("A2")
	if len(A2) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "A2", A2)
	}
	Z2 := c.GetString("Z2")
	if len(Z2) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "Z2", Z2)
	}
	element := c.GetString("element")
	if len(element) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "element", element)
	}
	A := c.GetString("A")
	if len(A) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "A", A)
	}
	Z := c.GetString("Z")
	if len(Z) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "Z", Z)
	}

	count, err := models.GetFissionsCount(element1, A1, Z1, element2, A2, Z2, element, A, Z)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	limit := "50"
	if len(c.GetString("limit")) > 0 {
		limit = c.GetString("limit")
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "limit", limit)
	}
	dataPerPage, err := strconv.Atoi(limit)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	paginator := pagination.SetPaginator(c.Ctx, dataPerPage, count)

	fissions, err := models.GetFissions(element1, A1, Z1, element2, A2, Z2, element, A, Z, paginator.Offset(), dataPerPage)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	c.Data["Fissions"] = fissions
	c.Data["paginator"] = paginator
	c.Data["Count"] = count
	c.Data["QueryString"] = queryString
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

func (c *FissionController) Post() {

	queryString := ""
	element1 := c.GetString("element1")
	if len(element1) > 0 {
		queryString += fmt.Sprintf("%s=%s", "element1", element1)
	}
	A1 := c.GetString("A1")
	if len(A1) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "A1", A1)
	}
	Z1 := c.GetString("Z1")
	if len(Z1) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "Z1", Z1)
	}
	element2 := c.GetString("element2")
	if len(element2) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "element2", element2)
	}
	A2 := c.GetString("A2")
	if len(A2) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "A2", A2)
	}
	Z2 := c.GetString("Z2")
	if len(Z2) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "Z2", Z2)
	}
	element := c.GetString("element")
	if len(element) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "element", element)
	}
	A := c.GetString("A")
	if len(A) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "A", A)
	}
	Z := c.GetString("Z")
	if len(Z) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "Z", Z)
	}

	count, err := models.GetFissionsCount(element1, A1, Z1, element2, A2, Z2, element, A, Z)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	limit := "50"
	if len(c.GetString("limit")) > 0 {
		limit = c.GetString("limit")
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "limit", limit)
	}
	dataPerPage, err := strconv.Atoi(limit)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	paginator := pagination.SetPaginator(c.Ctx, dataPerPage, count)

	fissions, err := models.GetFissions(element1, A1, Z1, element2, A2, Z2, element, A, Z, paginator.Offset(), dataPerPage)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	c.Data["Fissions"] = fissions
	c.Data["paginator"] = paginator
	c.Data["Count"] = count
	c.Data["QueryString"] = queryString
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

func (c *FissionController) Csv() {

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

	fissions, err := models.GetFissions(element1, A1, Z1, element2, A2, Z2, element, A, Z, 0, int(count))
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	//filename:=fmt.Sprintf("fissions_%s",element)

	c.Ctx.Output.Header("Content-Type", "text/csv")
	c.Ctx.Output.Header("Content-Disposition", "attachment; filename=fissions.csv")

	writer := csv.NewWriter(c.Controller.Ctx.ResponseWriter)
	err = writer.Write(models.FissionToHeader())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, fission := range fissions {
		err := writer.Write(fission.ToList())
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()
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
