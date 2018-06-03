package controllers

import (
	"github.com/lanfeust21/AParkhomov/wepapp/models"
	"net/http"
	"encoding/csv"
	"fmt"
	"strings"
)

type IsotopesController struct {
	BaseController
}

func (c *IsotopesController) Get() {
	queryString := ""
	sortorders := []string{"-Mev"}
	sortorder := c.GetString("sortorder")
	if len(sortorder) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "sortorder", sortorder)
		sortorders = strings.Split(sortorder, ",")
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

	isotopes, err := models.GetAllIsotopes(sortorders)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	c.Data["Count"] = len(isotopes)
	c.Data["Isotopes"] = isotopes
	c.Data["Sortorder"] = sortorder
	c.Data["QueryString"] = queryString
	c.TplName = "isotopes.tpl"
}


func (c *IsotopesController) Post() {

	queryString := ""
	sortorders := []string{"-Mev"}
	sortorder := c.GetString("sortorder")
	if len(sortorder) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "sortorder", sortorder)
		sortorders = strings.Split(sortorder, ",")
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

	count, err := models.GetIsotopesCount( element, A, Z)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	isotopes, err := models.GetIsotopes( element, A, Z,sortorders, 0, int(count))
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	c.Data["Count"] = len(isotopes)
	c.Data["Isotopes"] = isotopes
	c.Data["Element"] = element
	c.Data["A"] = A
	c.Data["Z"] = Z
	c.Data["Sortorder"] = sortorder
	c.Data["QueryString"] = queryString
	c.TplName = "isotopes.tpl"
}



func (c *IsotopesController) Csv() {

	queryString := ""
	sortorders := []string{"-Mev"}
	sortorder := c.GetString("sortorder")
	if len(sortorder) > 0 {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%s=%s", "sortorder", sortorder)
		sortorders = strings.Split(sortorder, ",")
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

	count, err := models.GetIsotopesCount( element, A, Z)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	isotopes, err := models.GetIsotopes( element, A, Z,sortorders, 0, int(count))
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	c.Ctx.Output.Header("Content-Type", "text/csv")
	c.Ctx.Output.Header("Content-Disposition", "attachment; filename=fusions.csv")

	writer := csv.NewWriter(c.Controller.Ctx.ResponseWriter)
	err = writer.Write(models.IsotopeToHeader())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, isotope := range isotopes {
		err := writer.Write(isotope.ToList())
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()
}

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
