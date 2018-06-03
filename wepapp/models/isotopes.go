package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/go-errors/errors"
	"strconv"
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
func (i *Isotope) ToList() []string {
	return []string{ i.Element,strconv.Itoa(i.A),strconv.Itoa(i.Z),i.AEM,i.Mev,i.Percent}
}
func IsotopeToHeader() []string {
	return []string{"Element","A","Z","AEM","Mev","%"}
}

func (i *Isotope) TableName() string {
	return "stable_isotopes"
}

func GetAllIsotopes(sortorder []string) ([]*Isotope, error) {
	var isotopes []*Isotope

	q := orm.NewOrm().QueryTable(&Isotope{}).OrderBy("-id")
	if len(sortorder) > 0 {
		q = q.OrderBy(sortorder...)
	} else {
		q = q.OrderBy("-Mev")
	}

	if _, err := q.All(&isotopes); err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return isotopes, nil
}

func GetIsotopes( element string, A string, Z string,sortorder []string, offset int, limit int) ([]*Isotope, error) {
	var isotopes []*Isotope

	q := orm.NewOrm().QueryTable(&Isotope{}).OrderBy([]string{"-Mev","-%"}...).Limit(limit).Offset(offset)
	if len(sortorder) > 0 {
		q = q.OrderBy(sortorder...)
	} else {
		q = q.OrderBy("-Mev")
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