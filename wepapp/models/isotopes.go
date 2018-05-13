package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/go-errors/errors"
)

func init() {
	orm.RegisterModel(new(Isotope))
}



type Isotope struct {
	Id      int64  `orm:"auto;pk;column(id)" db:"id"`
	Element string `orm:"column(element)"  db:"element"`
	A       int `orm:"column(A)"  db:"A"`
	Z       int `orm:"column(Z)"  db:"Z"`
	AEM     string `orm:"column(a.e.m)"  db:"a.e.m"`
	Mev     string `orm:"column(Mev)"  db:"Mev"`
	Percent string `orm:"column(%)"  db:"%"`
}

func (i *Isotope) TableName() string {
	return "stable_isotopes"
}

func GetIsotopes() ([]*Isotope, error) {
	var isotopes []*Isotope

	q := orm.NewOrm().QueryTable(&Isotope{}).OrderBy("-id")

	if _, err := q.All(&isotopes); err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return isotopes, nil
}

func GetIsotopesFrom(element string, A string, Z string) ([]*Isotope, error) {
	var isotopes []*Isotope

	q := orm.NewOrm().QueryTable(&Isotope{}).OrderBy("-id").Limit(50)

	if len(element) > 0 {
		q = q.Filter("element", element)
	}
	if len(A) > 0 {
		q = q.Filter("A", A)
	}
	if len(Z) > 0 {
		q = q.Filter("Z", Z)
	}

	if _, err := q.All(&isotopes); err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return isotopes, nil
}

type Element struct {
	Element string
}

func GetIsotopeElements(element string) ([]string, error) {

	var elements []Element
	_,err := orm.NewOrm().Raw("SELECT DISTINCT  T0.`element` FROM `stable_isotopes` T0 WHERE T0.`element` LIKE '" + element + "%' ORDER BY T0.`element` ASC LIMIT 10").QueryRows(&elements)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	elms := []string{}
	for _, elm := range elements {
		elms = append(elms, elm.Element)
	}

	return elms, nil
}

func GetIsotopesCount(element string, A string, Z string) (int64, error) {
	q:= orm.NewOrm().QueryTable(&Isotope{})
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