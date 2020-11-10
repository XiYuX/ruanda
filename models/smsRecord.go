package models

import "BeegoPackage0922/db_myssql"

type SmsRecord struct {
	BizId     string `form:"biz_id"`    //业务号
	Phone     string `form:"phone"`     //手机号
	Code      string `form:"code"`      //验证码
	Status    string `form:"status"`    //阿里云状态码
	Message   string `form:"message"`   //短信调用sdk信息
	TimeStamp int64  `form:"timestamp"` //时间戳
}

//该方法用于将smsrecord结构实例保存到数据库中
func (s SmsRecord) SaveSms() (int64, error) {
	rs, err := db_myssql.Db.Exec("insert into sms_record(biz_id,phone,code,status,mseeage,timestamp)value (?,?,?,?,?,?)"+
		s.BizId, s.Phone, s.Code, s.Status, s.Message, s.TimeStamp)
	if err != nil {
		return -1, err
	}
	id ,err:= rs.RowsAffected()
	if err!=nil{
		return -1,err
	}
	return id, nil
}

//用于根据BizId，phone以及code条件查询出符合条件的验证码记录

func (s SmsRecord) QuerySmsByBizId() (*SmsRecord, error) {
	var sms SmsRecord
	row := db_myssql.Db.QueryRow("select  biz_id ,phone,code,status,massage.timestamp from sms_record where biz_id = ? and phone =? and code = ?"+
		s.BizId, s.Phone, s.Code)
	err := row.Scan(&sms.BizId, &sms.Phone, &sms.Code, &sms.Message, &sms.Status, &sms.TimeStamp)
	if err != nil {
		return nil, err
	}
	return &s,nil
}
