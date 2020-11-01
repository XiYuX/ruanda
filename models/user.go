package models

import (
	"DataCertProject/db_mysql"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type User struct {
	Id int     `form:"id"`
	Phone string `form:"phone"`
	Password string `form:"password"`

}

//定义一个方法

func (u User) SaveUser()(int64,error){
	//1.密码脱敏处理
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes := md5Hash.Sum(nil)
	u.Password =hex.EncodeToString(passwordBytes)
	//2.执行数据库操作
	row,err:=db_mysql.Db.Exec("insert into user(phone,password)"+"values(?,?)",u.Phone,u.Password)
	if err!=nil{
		return -1,err
	}
	id,err :=row.RowsAffected()
	if err!=nil{
		return -1,err
	}
	return id,nil
}


//查询用户信息

func (u User) QueryUser()(*User,error) {
	if u.Password=="" {
		err := fmt.Errorf("hello error")
		return nil,err
	}else {
		md5Hash := md5.New()
		md5Hash.Write([]byte(u.Password))
		passwordBytes := md5Hash.Sum(nil)
		u.Password =hex.EncodeToString(passwordBytes)
		row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ? ",
			u.Phone,u.Password)
		err := row.Scan(&u.Phone)
		if err != nil {
			return nil,err
		}
		return &u,nil
	}
}