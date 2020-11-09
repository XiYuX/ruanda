package util

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/astaxie/beego"
	"math/rand"
	"strings"
	"time"
)

const SMS_TPL_REGISTER = "SMS_205393604" //用户注册模板号
const SMS_TPL_LOGIN = "SMS_205398654"    //用户登录的模板号
const SMS_TPL_KYC = ""                   //实名认证的模板号

type SmsCode struct {
	Code string `json:"code"`
}

type SmsResult struct {
	//business 商业,业务
	BizId     string//业务编码 唯一
	Code      string//短信发送调用状态码:OK,Failed
	Message   string//对应状态码的详细的说明信息
	RequestId string//网络请求的id
}

/**
 * 该函数用于向指定的手机号码发送一条短信
 * phone：接收短信的手机号
 * code：验证码
 * templateType: 模板类型
 */
func SendSms(phone string, code string, templateType string) (*SmsResult, error)  {
	config := beego.AppConfig
	accessKey := config.String("sms_access_key")
	accessKeySecret := config.String("sms_access_secret")
	signName := config.String("sms_sign_name")
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKey, accessKeySecret)

	//CreateSendSmsRequest: 创建一个发送短信息的请求
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https" //协议使用https，更安全

	//电话号码
	request.PhoneNumbers = phone
	//短信签名
	request.SignName = signName
	//短信模板号
	request.TemplateCode = templateType

	smsCode := SmsCode{
		Code: code,
	}
	codeBytes, _ := json.Marshal(smsCode)
	//指定短信模板中的动态验证码 的数据
	request.TemplateParam = string(codeBytes)

	//调用阿里云发送短信
	response, err := client.SendSms(request)
	if err != nil {
		return nil, err
	}
	result := &SmsResult{
		BizId:     response.BizId,
		Code:      response.Code,
		Message:   response.Message,
		RequestId: response.RequestId,
	}
	return result, nil
}

/**
 * 用于产生一个固定长度的随机验证码的函数
 */
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)

	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}