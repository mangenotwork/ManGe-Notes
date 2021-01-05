package object

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

const InstallJsonPath = "./install.json"

var InstallJsonFileLock sync.Mutex

type InstallInfo struct {
	Versions          string `json:"versions"` //安装版本
	Step              int    `json:"step"`     //安装进度
	DBType            string `json:"db_type"`  //数据存放类型， mysql,pgsql,sqlite
	MysqlHost         string `json:"mysql_host"`
	MysqlPort         string `json:"mysql_port"`
	MysqlUser         string `json:"mysql_user"`
	MysqlPassword     string `json:"mysql_password"`
	MysqlDBName       string `json:"mysql_dbname"`
	PgsqlHost         string `json:"pgsql_host"`
	PgsqlUser         string `json:"pgsql_user"`
	PgsqlPassword     string `json:"pgsql_password"`
	PgsqlDBName       string `json:"pgsql_dbname"`
	SqlitePath        string `json:"sqlite_path"`
	MediaType         string `json:"media_type"` //多媒体资源存放位置  ali  tencent
	MediaPath         string `json:"media_path"`
	AliOSSAccessKeyid string `json:"oss_access_keyid"` //阿里云对象存储
	AliOSSSecret      string `json:"oss_secret"`       //阿里云对象存储
	AliOSSEndpoint    string `json:"oss_endpoint"`     //阿里云对象存储
	AliOSSBucketName  string `json:"oss_bucketName"`   //阿里云对象存储
	TencenCosUrl      string `json:"cos_url"`          //腾讯云对象存储
	TencenSecretid    string `json:"cos_secretid"`     //腾讯云对象存储
	TencenSecretkey   string `json:"cos_secretkey"`    //腾讯云对象存储
}

//写入数据到install
func WriteInstallInfo(installInfo string) error {
	InstallJsonFileLock.Lock()
	defer InstallJsonFileLock.Unlock()
	f, err := os.Create(InstallJsonPath)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.WriteString(installInfo)
	if err != nil {
		return err
	}
	return nil
}

//读取install info
func OpenInstallFile() InstallInfo {
	InstallJsonFileLock.Lock()
	defer InstallJsonFileLock.Unlock()
	file, _ := os.Open(InstallJsonPath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	installInfo := InstallInfo{}
	err := decoder.Decode(&installInfo)
	if err != nil {
		log.Println("Error:", err)
	}
	return installInfo
}
