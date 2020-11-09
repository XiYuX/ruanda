package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//注册页面
	beego.Router("/", &controllers.MainController{})
	//用户注册的接口请求
	beego.Router("/user_register", &controllers.RegisterController{})
	//直接登录的页面请求接口
	beego.Router("/login.html", &controllers.LoginController{})
	//用户登录请求接口
	beego.Router("/user_login", &controllers.LoginController{})
	//文件上传接口
	beego.Router("/upload", &controllers.UploadFileController{})
	//查看认证数据的证书(cert_detail.html)
	beego.Router("/cert_detail.html",&controllers.CertDetsilController{})
	//
	beego.Router("/user_kyc.html",&controllers.UserKycController{})
	//用户实名认证接口
	beego.Router("/user_kyc",&controllers.UserKycController{})
	//短信验证码登录
	beego.Router("/login_sms.html",&controllers.SmsLoginController{})
	//
	beego.Router("/sms_")
}

