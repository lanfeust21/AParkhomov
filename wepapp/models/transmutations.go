package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/go-errors/errors"
)

func init() {
	orm.RegisterModel(new(Transmutation22))
}

type Transmutation22 struct {
	Id       int64  `orm:"auto;pk;column(id)" db:"id"`
	Element1 string `orm:"column(element1)"  db:"element1"`
	A1       int    `orm:"column(A1)"  db:"A1"`
	Z1       int    `orm:"column(Z1)"  db:"Z1"`
	Element2 string `orm:"column(element2)"  db:"element2"`
	A2       int    `orm:"column(A2)"  db:"A2"`
	Z2       int    `orm:"column(Z2)"  db:"Z2"`
	Element3 string `orm:"column(element3)"  db:"element3"`
	A3       int    `orm:"column(A3)"  db:"A3"`
	Z3       int    `orm:"column(Z3)"  db:"Z3"`
	Element4 string `orm:"column(element4)"  db:"element4"`
	A4       int    `orm:"column(A4)"  db:"A4"`
	Z4       int    `orm:"column(Z4)"  db:"Z4"`
	Mev      string `orm:"column(Mev)"  db:"Mev"`
}

func (i *Transmutation22) TableName() string {
	return "stable_isotopes_22"
}

func (i *Transmutation22) ToList() []string {
	return []string{i.Element1, string(i.A1), string(i.Z1), i.Element2, string(i.A2), string(i.Z2), i.Element3, string(i.A3), string(i.Z3), i.Element4, string(i.A4), string(i.Z4), i.Mev}
}
func (i *Transmutation22) ToHeader() []string {
	return []string{"Element1", "A1", "Z1", "Element2", "i.A2", "i.Z2", "Element3", "A3", "Z3", "Element4", "A4", "Z4", "Mev"}
}

func GetAllTransmutations(sortorder[]string, offset int, limit int) ([]*Transmutation22, error) {
	var transmutations []*Transmutation22

	q := orm.NewOrm().QueryTable(&Transmutation22{}).Limit(limit).Offset(offset)
	if len(sortorder) > 0 {
		q = q.OrderBy(sortorder...)
	} else {
		q = q.OrderBy("-Mev")
	}
	if _, err := q.All(&transmutations); err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return transmutations, nil
}
func GetTransmutationsCount(element1 string, A1 string, Z1 string, element2 string, A2 string, Z2 string, element3 string, A3 string, Z3 string, element4 string, A4 string, Z4 string) (int64, error) {
	q := orm.NewOrm().QueryTable(&Transmutation22{})
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

	if len(element3) > 0 {
		q = q.Filter("element3", element3)
	}
	if len(A3) > 0 {
		q = q.Filter("A3", A3)
	}
	if len(Z3) > 0 {
		q = q.Filter("Z3", Z3)
	}

	if len(element4) > 0 {
		q = q.Filter("element4", element4)
	}
	if len(A4) > 0 {
		q = q.Filter("A4", A4)
	}
	if len(Z4) > 0 {
		q = q.Filter("Z4", Z4)
	}
	count, err := q.Count()
	if err != nil {
		return -1, errors.Wrap(err, 0)
	}

	return count, nil
}

func GetTransmutations(element1 string, A1 string, Z1 string, element2 string, A2 string, Z2 string, element3 string, A3 string, Z3 string, element4 string, A4 string, Z4 string,sortorder []string, offset int, limit int) ([]*Transmutation22, error) {
	var transmutations []*Transmutation22

	q := orm.NewOrm().QueryTable(&Transmutation22{}).Limit(limit).Offset(offset)

	if len(sortorder) > 0 {
		q = q.OrderBy(sortorder...)
	} else {
		q = q.OrderBy("-Mev")
	}
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

	if len(element3) > 0 {
		q = q.Filter("element3", element3)
	}
	if len(A3) > 0 {
		q = q.Filter("A3", A3)
	}
	if len(Z3) > 0 {
		q = q.Filter("Z3", Z3)
	}

	if len(element4) > 0 {
		q = q.Filter("element4", element4)
	}
	if len(A4) > 0 {
		q = q.Filter("A4", A4)
	}
	if len(Z4) > 0 {
		q = q.Filter("Z4", Z4)
	}

	if _, err := q.All(&transmutations); err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return transmutations, nil
}

type Element1 struct {
	Element1 string
}

func GetTransmutationElement1s(element string) ([]string, error) {

	var elements []Element1
	_, err := orm.NewOrm().Raw("SELECT DISTINCT  T0.`element1` FROM `stable_isotopes_22` T0 WHERE T0.`element1` LIKE '" + element + "%' ORDER BY T0.`element1` ASC LIMIT 10").QueryRows(&elements)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	elms := []string{}
	for _, elm := range elements {
		elms = append(elms, elm.Element1)
	}

	return elms, nil
}

type Element2 struct {
	Element2 string
}

func GetTransmutationElement2s(element string) ([]string, error) {

	var elements []Element2
	_, err := orm.NewOrm().Raw("SELECT DISTINCT  T0.`element2` FROM `stable_isotopes_22` T0 WHERE T0.`element2` LIKE '" + element + "%' ORDER BY T0.`element2` ASC LIMIT 10").QueryRows(&elements)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	elms := []string{}
	for _, elm := range elements {
		elms = append(elms, elm.Element2)
	}

	return elms, nil
}

type Element3 struct {
	Element3 string
}

func GetTransmutationElement3s(element string) ([]string, error) {

	var elements []Element3
	_, err := orm.NewOrm().Raw("SELECT DISTINCT  T0.`element3` FROM `stable_isotopes_22` T0 WHERE T0.`element3` LIKE '" + element + "%' ORDER BY T0.`element3` ASC LIMIT 10").QueryRows(&elements)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	elms := []string{}
	for _, elm := range elements {
		elms = append(elms, elm.Element3)
	}

	return elms, nil
}

type Element4 struct {
	Element4 string
}

func GetTransmutationElement4s(element string) ([]string, error) {

	var elements []Element4
	_, err := orm.NewOrm().Raw("SELECT DISTINCT  T0.`element4` FROM `stable_isotopes_22` T0 WHERE T0.`element4` LIKE '" + element + "%' ORDER BY T0.`element4` ASC LIMIT 10").QueryRows(&elements)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	elms := []string{}
	for _, elm := range elements {
		elms = append(elms, elm.Element4)
	}

	return elms, nil
}
