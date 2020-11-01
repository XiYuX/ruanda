package controllers

import (
	"DataCertProject/blockchain"
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
	//查到了
	//certId = hex.EncodeToString()
	c.Data["CertId"] = strings.ToUpper(string(blocks.Data))

	//跳转页面
	c.TplName = "cert_detail.html"
}
