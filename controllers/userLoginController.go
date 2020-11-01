package controllers

import (
	"github.com/astaxie/beego"
	"DataCertProject/models"
	"fmt"
)

type LoginController struct {
	beego.Controller
}


 //访问login.html页面的请求
func (l *LoginController) Get(){
	l.TplName = "login.html"
}

/**
 * 用户登录接口
 */
func (l  *LoginController) Post(){
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
	//登录成功,跳转项目核心功能页面
	l.Data["Phone"] = u.Phone
	l.TplName = "home.html"
}
