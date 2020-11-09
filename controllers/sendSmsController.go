package controllers

import (
	"DataCertProject/models"
	"DataCertProject/util"
	"github.com/astaxie/beego"
)

type SendSmsController struct {
	beego.Controller
}

//该方法用于短信验证
func (s *SendSmsController) Post() {
	//1、解析用户提交的手机号
	var sms models.SmsRecord
	if err := s.ParseForm(&sms); err != nil {
		s.Ctx.WriteString("抱歉，解析手机号失败，请重试！！！")
		return
	}
	//2、调用工具函数生成一个验证码
	code := util.GenValidateCode(6)
	//3、将生成的验证码调用阿里云sdk，进行发送
	result, err := util.SendSms(sms.Phone, code, util.SMS_TPL_LOGIN)
	//4、接收阿里云sdk的调用结果，进行判断并处理
	//	①发送失败，将错误信息返回给前端页面进行提示
	//调用失败：比如网络断开，链接超时
	if err != nil {
		s.Ctx.WriteString("发送验证码失败，请重试！！！")
		return
	}
	//调用请求成功了，但是短信没有发送成功
	if result.Code != "OK" {
		s.Ctx.WriteString(result.Message)
		return
	}
	//	②发送成功
	//a、将验证码存储mysql数据库中
	smsRecord := models.SmsRecord{
		BizId:   result.BizId,
		Phone:   sms.Phone,
		Code:    code,
		Status:  result.Code,
		Message: result.Message,
	}
	_,err = smsRecord.SaveSms()
	if err!=nil {
		s.Ctx.WriteString("获取验证码失败，请重试！！！")
		return
	}
	//b、跳转到登录提交页面

	s.TplName = "login_sms.html"


}
