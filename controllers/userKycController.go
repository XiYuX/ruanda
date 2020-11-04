package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
)

type UserKycController struct {
	beego.Controller
}

func (u *UserKycController) Get() {

}
func (u *UserKycController) Post(){
	var user models.User
	err:= u.ParseForm(&user)
	if err !=nil{
		u.Ctx.WriteString("用户注册失败，请重试！！！")
		return
	}
	//
	record,err:=models.QueryRecordByPhone(user.Phone)
	if err!=nil{
		u.Ctx.WriteString("抱歉，获取认证失败")
		return
	}
	u.Data["Record"] = record
	u.Data["phone"] = user.Phone
	u.TplName = "list_record.html"
}