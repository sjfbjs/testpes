package controllers

import (
	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}
func (c *UploadController) Get() {
	path := c.Ctx.Input.Param(":path")
	ext := c.Ctx.Input.Param(":ext")
	beego.Info(path)
	beego.Include(ext)
	c.Ctx.ResponseWriter.Write([]byte("Path:" + path + ", Ext:" + ext) )
}