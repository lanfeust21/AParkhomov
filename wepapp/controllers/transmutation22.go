package controllers

import (
	"github.com/lanfeust21/AParkhomov/wepapp/models"
	"net/http"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
	"encoding/csv"
	"fmt"
	"strings"
)

type Transmutation22Controller struct {
	BaseController
}

func (c *Transmutation22Controller) Get() {

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
	element3 := c.GetString("element3")
	if len(element3) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "element3", element3)
	}
	A3 := c.GetString("A3")
	if len(A3) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "A3", A3)
	}
	Z3 := c.GetString("Z3")
	if len(Z3) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "Z3", Z3)
	}
	element4 := c.GetString("element4")
	if len(element4) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "element4", element4)
	}
	A4 := c.GetString("A4")
	if len(A4) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "A4", A4)
	}
	Z4 := c.GetString("Z4")
	if len(Z4) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "Z4", Z4)
	}
	sortorders := []string{"-Mev"}
	sortorder := c.GetString("sortorder")
	if len(sortorder) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "sortorder", sortorder)
		sortorders = strings.Split(sortorder, ",")
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

	count, err := models.GetTransmutationsCount(element1, A1, Z1, element2, A2, Z2, element3, A3, Z3, element4, A4, Z4)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	paginator := pagination.SetPaginator(c.Ctx, dataPerPage, count)

	transmutations, err := models.GetTransmutations(element1, A1, Z1, element2, A2, Z2, element3, A3, Z3, element4, A4, Z4,sortorders, paginator.Offset(), dataPerPage)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	c.Data["Transmutations"] = transmutations
	c.Data["paginator"] = paginator
	c.Data["Count"] = count
	c.Data["QueryString"] = queryString
	c.Data["Element1"] = element1
	c.Data["A1"] = A1
	c.Data["Z1"] = Z1
	c.Data["Element2"] = element2
	c.Data["A2"] = A2
	c.Data["Z2"] = Z2
	c.Data["Element3"] = element3
	c.Data["A3"] = A3
	c.Data["Z3"] = Z3
	c.Data["Element4"] = element4
	c.Data["A4"] = A4
	c.Data["Z4"] = Z4
	c.Data["Limit"] = limit
	c.Data["Sortorder"] = sortorder
	c.TplName = "transmutation22.tpl"
}

func (c *Transmutation22Controller) Post() {

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
	element3 := c.GetString("element3")
	if len(element3) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "element3", element3)
	}
	A3 := c.GetString("A3")
	if len(A3) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "A3", A3)
	}
	Z3 := c.GetString("Z3")
	if len(Z3) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "Z3", Z3)
	}
	element4 := c.GetString("element4")
	if len(element4) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "element4", element4)
	}
	A4 := c.GetString("A4")
	if len(A4) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "A4", A4)
	}
	Z4 := c.GetString("Z4")
	if len(Z4) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "Z4", Z4)
	}
	sortorders := []string{"-Mev"}
	sortorder := c.GetString("sortorder")
	if len(sortorder) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "sortorder", sortorder)
		sortorders = strings.Split(sortorder, ",")
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

	count, err := models.GetTransmutationsCount(element1, A1, Z1, element2, A2, Z2, element3, A3, Z3, element4, A4, Z4)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	paginator := pagination.SetPaginator(c.Ctx, dataPerPage, count)

	transmutations, err := models.GetTransmutations(element1, A1, Z1, element2, A2, Z2, element3, A3, Z3, element4, A4, Z4,sortorders, paginator.Offset(), dataPerPage)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	c.Data["Transmutations"] = transmutations
	c.Data["paginator"] = paginator
	c.Data["Count"] = count
	c.Data["QueryString"] = queryString
	c.Data["Element1"] = element1
	c.Data["A1"] = A1
	c.Data["Z1"] = Z1
	c.Data["Element2"] = element2
	c.Data["A2"] = A2
	c.Data["Z2"] = Z2
	c.Data["Element3"] = element3
	c.Data["A3"] = A3
	c.Data["Z3"] = Z3
	c.Data["Element4"] = element4
	c.Data["A4"] = A4
	c.Data["Z4"] = Z4
	c.Data["Limit"] = limit
	c.Data["Sortorder"] = sortorder
	c.TplName = "transmutation22.tpl"
}

func (c *Transmutation22Controller) Csv() {
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
	sortorders := []string{"-Mev"}
	sortorder := c.GetString("sortorder")
	if len(sortorder) > 0 {
		sortorders = strings.Split(sortorder, ",")
	}

	count, err := models.GetTransmutationsCount(element1, A1, Z1, element2, A2, Z2, element3, A3, Z3, element4, A4, Z4)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	transmutations, err := models.GetTransmutations(element1, A1, Z1, element2, A2, Z2, element3, A3, Z3, element4, A4, Z4,sortorders, 0, int(count))
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	c.Ctx.Output.Header("Content-Type", "text/csv")
	c.Ctx.Output.Header("Content-Disposition", "attachment; filename=transmutations.csv")

	writer := csv.NewWriter(c.Controller.Ctx.ResponseWriter)
	err = writer.Write(transmutations[0].ToList())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, transmutation := range transmutations {
		err := writer.Write(transmutation.ToList())
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()
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
