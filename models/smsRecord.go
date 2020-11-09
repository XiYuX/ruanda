package models

import "BeegoPackage0922/db_myssql"

type SmsRecord struct {
	BizId   string `form:"biz_id"`  //业务号
	Phone   string `form:"phone"`   //手机号
	Code    string `form:"code"`    //验证码
	Status  string `form:"status"`  //阿里云状态码
	Message string `form:"message"` //短信调用sdk信息
}
//该方法用于将smsrecord结构实例保存到数据库中
func (s SmsRecord)SaveSms(){
	rs,err:=db_myssql.Db.Exec("insert into sms_record(biz_id,phone,code,status,mseeage)value (?,?,?,?,?)"+
		s.BizId,s.Phone,s.Code,s.Status,s.Message)
	if err!=nil{

	}


}


