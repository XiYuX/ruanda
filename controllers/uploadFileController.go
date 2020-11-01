package controllers

import (
	"DataCertProject/blockchain"
	"DataCertProject/models"
	"DataCertProject/util"
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"time"
)

type UploadFileController struct {
	beego.Controller
}

//使用post方法上传

func (u *UploadFileController)Post(){
	//1、获取客户端上传的文件、以及其他from表单信息
	//获取标题
	fileTitle :=u.Ctx.Request.PostFormValue("upload_title")
	phone := u.Ctx.Request.PostFormValue("phone")
	//获取文件
	file,header,err :=u.GetFile("upload_file")
	if err != nil{
		u.TplName = "fileErrorPage.hyml"
		return
	}
	fmt.Println("自定义的文件名称：",fileTitle)
	fmt.Println("文件名称",header.Filename)
	fmt.Println("文件大小",header.Size)//字节大小

	fmt.Println(file)

	//2、将文件保存在本地的一个目录当中
	//文件全路径：路径 +文件名 +"."+扩展名
	//要的文件的路径
	uploadDir := "./static/img" + header.Filename
	//创建一个writer：用于向硬盘上写一个文件
	saveFile, err := os.OpenFile(uploadDir, os.O_RDWR|os.O_CREATE, 777)

	//创建一个writer: 用于向硬盘上写一个文件
	writer := bufio.NewWriter(saveFile)
	_, err = io.Copy(writer, file)
	if err != nil { //invalid argument
		fmt.Println(err.Error())
		u.Ctx.WriteString("抱歉，保存电子数据失败，请重试")
		return
	}

	defer saveFile.Close()

	//2、计算文件的hash
	hashFile, err := os.Open(uploadDir)
	defer hashFile.Close()
	hash, err := util.MD5HashReader(hashFile)

	//3、将上传的记录保存到数据库中
	record := models.UploadRecord{}
	record.FileName = header.Filename
	record.FileSize = header.Size
	record.FileTitle = fileTitle
	record.CertTime = time.Now().Unix() //毫秒数
	record.FileCert = hash
	record.Phone = phone //手机
	_, err = record.SaveRecord()
	if err != nil {
		u.TplName = "fileErrorPage.html"
		return
	}
	//将要认证的文件的hash值及个人实名信息保存到区块链上
	_ ,err =blockchain.CHAIN.SaveData([]byte(hash))
	if err !=nil{
		u.TplName = "blockErrorPage.html"
		return
	}

	//4、从数据库中读取phone用户对应的所有认证数据记录
	records, err := models.QueryRecordByPhone(phone)

	//5、根据文件保存结果，返回相应的提示信息或者页面跳转
	if err != nil {
		u.TplName = "fileErrorPage.html"
		return
	}
	fmt.Println(records)
	u.Data["Records"] = records
	u.Data["Phone"] = phone
	u.TplName = "list_record.html"
}
