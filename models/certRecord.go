package models

import (
	"bytes"
	"encoding/gob"
)

type CertRecord struct {
	CerHash        []byte //认证文件hash值
	CerHashStr     string
	CertId         []byte //认证ID
	CertIdStr      string
	CertAuthor     string //认证人
	Phone          string //联系方式
	AuthorCard     string //身份证
	FileName       string //认证文件名称
	FileSize       int64  //文件大小
	CertTime       int64  //认证时间
	CertTimeFormat string
}

//认证数据记录的序列化
func (c CertRecord) SerializeRecord() ([]byte, error) {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(c)
	return buff.Bytes(), err
}

//反序列化
func DeSerializeRecord(data []byte) (*CertRecord, error) {
	var certRecord *CertRecord
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&certRecord)
	return certRecord, err
}
