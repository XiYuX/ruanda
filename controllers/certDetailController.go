package controllers

import (
	"DataCertProject/blockchain"
	"DataCertProject/models"
	"DataCertProject/util"
	"github.com/astaxie/beego"
	"strings"
)

type CertDetsilController struct {
	beego.Controller
}

func (c *CertDetsilController) Get() {
	//获取前端页面get请求时携带的cert_id数据
	certId := c.GetString("cert_id")

	//准备数据:根据cert_id到区块链上查询具体的信息,获取区块信息
	blocks, err := blockchain.CHAIN.QueryBlockByCertId([]byte(certId))
	if err != nil {
		c.Ctx.WriteString("链上数据查询失败")
		return
	}
	//查询未遇到错误:未查到
	if blocks == nil {
		c.Ctx.WriteString("抱歉,未查到链上数据,请重试!!!")
		return
	}
	//序列化
	certRecord,err:=models.DeSerializeRecord(blocks.Data)
	certRecord.CerHashStr =string(certRecord.CerHash)
	certRecord.CertIdStr = strings.ToUpper(string(certRecord.CertId))
	certRecord.CertTimeFormat = util.TimeFormat(certRecord.CertTime,0,util.TIME_FORMAT_THREE)
	//查到了
	//certId = hex.EncodeToString()
	c.Data["CertRecord"] = certRecord

	//跳转页面
	c.TplName = "cert_detail.html"
}
