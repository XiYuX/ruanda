package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
)

type RegisterCon struct {
	beego.Controller
}

func (r *RegisterCon) Post() {
	var user models.User
	//解析数据
	err := r.ParseForm(&user)
	if err!=nil{
		r.TplName ="analysisErrorPage.html"
		return
	}
	//将数据存入数据库
	_,err = user.SaveUser()
	if err!=nil{
		r.TplName ="analysisErrorPage.html"
		return
	}
	//
	r.TplName = "login.html"



}