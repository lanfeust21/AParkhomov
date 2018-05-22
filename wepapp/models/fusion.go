package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/go-errors/errors"
	"strconv"
)

func init() {
	orm.RegisterModel(new(Fusion))
}

type Fusion struct {
	Id       int64  `orm:"auto;pk;column(id)" db:"id"`
	Element1 string `orm:"column(element1)"  db:"element1"`
	A1       int    `orm:"column(A1)"  db:"A1"`
	Z1       int    `orm:"column(Z1)"  db:"Z1"`
	Element2 string `orm:"column(element2)"  db:"element2"`
	A2       int    `orm:"column(A2)"  db:"A2"`
	Z2       int    `orm:"column(Z2)"  db:"Z2"`
	Element  string `orm:"column(element)"  db:"element"`
	A        int    `orm:"column(A)"  db:"A"`
	Z        int    `orm:"column(Z)"  db:"Z"`
	Mev      string `orm:"column(Mev)"  db:"Mev"`
}

func (i *Fusion) TableName() string {
	return "fusion"
}
func (i *Fusion) ToList() []string {
	return []string{i.Element1,strconv.Itoa(i.A1),strconv.Itoa(i.Z1),i.Element2,strconv.Itoa(i.A2),strconv.Itoa(i.Z2),i.Element,strconv.Itoa(i.A),strconv.Itoa(i.Z),i.Mev}
}
func FusionToHeader() []string {
	return []string{"Element1","A1","Z1","Element2","A2","Z2","Element","A","Z","Mev"}
}

func GetAllFusions(offset int, limit int) ([]*Fusion, error) {
	var Fusions []*Fusion

	q := orm.NewOrm().QueryTable(&Fusion{}).Limit(limit).Offset(offset).OrderBy("-Mev")

	if _, err := q.All(&Fusions); err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return Fusions, nil
}
func GetFusionsCount(element1 string, A1 string, Z1 string, element2 string, A2 string, Z2 string, element string, A string, Z string) (int64, error) {
	q := orm.NewOrm().QueryTable(&Fusion{})
	if len(element1) > 0 {
		q = q.Filter("element1", element1)
	}
	if len(A1) > 0 {
		q = q.Filter("A1", A1)
	}
	if len(Z1) > 0 {
		q = q.Filter("Z1", Z1)
	}

	if len(element2) > 0 {
		q = q.Filter("element2", element2)
	}
	if len(A2) > 0 {
		q = q.Filter("A2", A2)
	}
	if len(Z2) > 0 {
		q = q.Filter("Z2", Z2)
	}

	if len(element) > 0 {
		q = q.Filter("element", element)
	}
	if len(A) > 0 {
		q = q.Filter("A", A)
	}
	if len(Z) > 0 {
		q = q.Filter("Z", Z)
	}

	count, err := q.Count()
	if err != nil {
		return -1, errors.Wrap(err, 0)
	}

	return count, nil
}

func GetFusions(element1 string, A1 string, Z1 string, element2 string, A2 string, Z2 string, element string, A string, Z string, offset int, limit int) ([]*Fusion, error) {
	var Fusions []*Fusion

	q := orm.NewOrm().QueryTable(&Fusion{}).Limit(limit).Offset(offset).OrderBy("-Mev")

	if len(element1) > 0 {
		q = q.Filter("element1", element1)
	}
	if len(A1) > 0 {
		q = q.Filter("A1", A1)
	}
	if len(Z1) > 0 {
		q = q.Filter("Z1", Z1)
	}

	if len(element2) > 0 {
		q = q.Filter("element2", element2)
	}
	if len(A2) > 0 {
		q = q.Filter("A2", A2)
	}
	if len(Z2) > 0 {
		q = q.Filter("Z2", Z2)
	}

	if len(element) > 0 {
		q = q.Filter("element", element)
	}
	if len(A) > 0 {
		q = q.Filter("A", A)
	}
	if len(Z) > 0 {
		q = q.Filter("Z", Z)
	}

	if _, err := q.All(&Fusions); err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return Fusions, nil
}

func GetFusionElement1s(element string) ([]string, error) {

	var elements []Element1
	_, err := orm.NewOrm().Raw("SELECT DISTINCT  T0.`element1` FROM `fusion` T0 WHERE T0.`element1` LIKE '" + element + "%' ORDER BY T0.`element1` ASC LIMIT 10").QueryRows(&elements)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	elms := []string{}
	for _, elm := range elements {
		elms = append(elms, elm.Element1)
	}

	return elms, nil
}

func GetFusionElement2s(element string) ([]string, error) {

	var elements []Element2
	_, err := orm.NewOrm().Raw("SELECT DISTINCT  T0.`element2` FROM `fusion` T0 WHERE T0.`element2` LIKE '" + element + "%' ORDER BY T0.`element2` ASC LIMIT 10").QueryRows(&elements)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	elms := []string{}
	for _, elm := range elements {
		elms = append(elms, elm.Element2)
	}

	return elms, nil
}

func GetFusionElements(element string) ([]string, error) {

	var elements []Element
	_, err := orm.NewOrm().Raw("SELECT DISTINCT  T0.`element` FROM `fusion` T0 WHERE T0.`element` LIKE '" + element + "%' ORDER BY T0.`element` ASC LIMIT 10").QueryRows(&elements)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	elms := []string{}
	for _, elm := range elements {
		elms = append(elms, elm.Element)
	}

	return elms, nil
}
