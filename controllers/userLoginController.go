package controllers

import (
	"DataCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type LoginController struct {
	beego.Controller
}

//访问login.html页面的请求
func (l *LoginController) Get() {
	l.TplName = "login.html"
}

/**
 * 用户登录接口
 */
func (l *LoginController) Post() {
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		l.TplName = "anaiysisErrorPage.html"
		return
	}
	//查询数据库的用户信息
	u, err := user.QueryUser()
	if err != nil {
		fmt.Println(err.Error())
		l.TplName = "userErrorPage.html"
		return
	}
	//trim：修剪 (将字符串中两端的空格去处)
	name:=strings.TrimSpace(u.Name)
	card :=strings.TrimSpace(u.Card)
	sex :=strings.TrimSpace(u.Sex)
	if name == ""||card =="" ||sex == ""{
		//直接跳转到实名页面
		l.Data[""] = u.Phone
		l.TplName = "user_kyc.html"
		return
	}
	id ,err:= user.Update()
	if err !=nil{
		l.Ctx.WriteString("")
		return
	}

	//登录成功,跳转项目核心功能页面
	l.Data["Phone"] = u.Phone
	l.TplName = "home.html"
}
