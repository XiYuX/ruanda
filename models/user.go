package models

import (
	"DataCertProject/db_mysql"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
	Name     string `form:name`
	Card     string `form:card`
	Sex      string `form:sex`
}

//定义一个方法

func (u User) SaveUser() (int64, error) {
	//1.密码脱敏处理
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes := md5Hash.Sum(nil)
	u.Password = hex.EncodeToString(passwordBytes)
	//2.执行数据库操作
	row, err := db_mysql.Db.Exec("insert into user(phone,password)"+"values(?,?)", u.Phone, u.Password)
	if err != nil {
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

//查询用户信息

func (u User) QueryUser() (*User, error) {
	if u.Password == "" {
		err := fmt.Errorf("hello error")
		return nil, err
	} else {
		md5Hash := md5.New()
		md5Hash.Write([]byte(u.Password))
		passwordBytes := md5Hash.Sum(nil)
		u.Password = hex.EncodeToString(passwordBytes)
		row := db_mysql.Db.QueryRow("select phone ,name, card, sex,from user where phone = ? and password = ? ",
			u.Phone, u.Password, )
		err := row.Scan(&u.Phone, &u.Name, &u.Card, &u.Sex)
		if err != nil {
			return nil, err
		}
		return &u, nil
	}
}

//根据用户的phone信息查询对应的用户信息
func QuerUserByPhone(phone string) (*User, error) {
	row := db_mysql.Db.QueryRow("select phone,name,card,sex from user where phone = ? ", phone)
	var user User
	err := row.Scan(&user.Phone, &user.Password, &user.Name, &user.Card, &user.Sex)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u User) Update() (int64, error) {
	rs, err := db_mysql.Db.Exec("update user set ,name = ?,card =?, sex = ?from user where phone = ?",u.Phone,u.Name,u.Card,u.Sex)
	if err != nil {
		return -1, err
	}
	return rs, nil
}
