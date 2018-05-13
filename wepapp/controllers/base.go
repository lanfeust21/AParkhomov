package controllers

import (
	"github.com/go-errors/errors"
	"github.com/astaxie/beego"

)
type BaseController struct {
	beego.Controller
}

func (c *BaseController) Error(status int, err error) {
	if _, ok := err.(*errors.Error); ok {
		beego.Error(err.(*errors.Error).ErrorStack())
	} else {
		beego.Error(err.Error())
	}

	c.CustomAbort(status, err.Error())
}


