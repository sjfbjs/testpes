package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}
func (c *TestController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.ResponseWriter.Write([]byte("Id:" + id) )
}