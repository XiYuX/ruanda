package controllers

import "github.com/astaxie/beego"

type SmsLoginController struct {
	beego.Controller
}

//该方法用于处理浏览器中的跳转手机验证码
func (s SmsLoginController) Get(){
	s.TplName = "login_sms.html"
}
