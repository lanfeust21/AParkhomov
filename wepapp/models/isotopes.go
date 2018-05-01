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
	A       string `orm:"column(A)"  db:"A"`
	Z       string `orm:"column(Z)"  db:"Z"`
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

func GetIsotopesFrom(element string,A string,Z string) ([]*Isotope, error) {
	var isotopes []*Isotope

	q := orm.NewOrm().QueryTable(&Isotope{}).OrderBy("-id").Limit(50)

	if len(element) > 0 {
		q = q.Filter("element", element)
	}
	if len(A) > 0 {
		q = q.Filter("A", element)
	}
	if len(Z) > 0 {
		q = q.Filter("element", Z)
	}

	if _, err := q.All(&isotopes); err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return isotopes, nil
}