package controllers

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/mangenotwork/ManGe-Notes/dao"

	"github.com/mangenotwork/ManGe-Notes/conn"
	"github.com/mangenotwork/ManGe-Notes/models"
	"github.com/mangenotwork/ManGe-Notes/object"
	"github.com/mangenotwork/ManGe-Notes/util"
)

type InstallController struct {
	Controller
}

//开始安装
func (this *InstallController) InstallStart() {
	//创建安装文件，并写入安装信息
	installInfo, err := json.Marshal(&object.InstallInfo{
		Versions: "v0.1",
		Step:     2,
	})

	log.Println("installInfo = ", string(installInfo), err)
	f, err := os.Create(object.InstallJsonPath)
	if err != nil {
		this.RetuenJson(0, 0, err)
		return
	}
	l, err := f.WriteString(string(installInfo))
	if err != nil {
		log.Println(err)
		f.Close()
		this.RetuenJson(0, 0, err)
		return
	}
	log.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		log.Println(err)
		this.RetuenJson(0, 0, err)
		return
	}

	this.RetuenJson(0, 0, string(installInfo))
}

//测试mysql连接
func (this *InstallController) InstallMysqlTset() {
	host := this.GetString("host")
	port := this.GetString("port")
	user := this.GetString("user")
	password := this.GetString("password")
	dbname := this.GetString("dbname")
	log.Println(host, port, user, password, dbname)

	if host == "" || user == "" || password == "" || dbname == "" {
		this.RetuenJson(1, 0, "mysql连接信息不全")
	}

	if port == "" {
		port = "3306"
	}

	//mysql 测试
	isconn, err := conn.MysqlConnTest(host, port, user, password, dbname)
	if isconn {
		this.RetuenJson(0, 0, "连接成功")
		return
	}
	this.RetuenJson(1, 0, err)
}

//mysql
func (this *InstallController) InstallMysql() {
	host := this.GetString("host")
	port := this.GetString("port")
	user := this.GetString("user")
	password := this.GetString("password")
	dbname := this.GetString("dbname")
	log.Println(host, port, user, password, dbname)

	if host == "" || user == "" || password == "" || dbname == "" {
		this.RetuenJson(1, 0, "mysql连接信息不全")
	}

	if port == "" {
		port = "3306"
	}

	//mysql 连接并创建
	isconn, _ := conn.CreateMysqlDB(host, port, user, password, dbname)
	if !isconn {
		this.RetuenJson(1, 0, "连接不上mysql.")
		return
	}

	//将信息写入 install.json
	infodata := object.OpenInstallFile()
	infodata.Step = 3
	infodata.DBType = "mysql"
	infodata.MysqlHost = host
	infodata.MysqlPort = port
	infodata.MysqlUser = user
	infodata.MysqlPassword = password
	infodata.MysqlDBName = dbname
	object.GlobalDBType = "mysqls"
	object.GlobalMysqlHost = host
	object.GlobalMysqlPort = port
	object.GlobalMysqlUser = user
	object.GlobalMysqlPassword = password
	object.GlobalMysqlDBName = dbname

	installInfo, err := json.Marshal(infodata)
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	err = object.WriteInstallInfo(string(installInfo))
	if err != nil {
		this.RetuenJson(1, 0, err)
	}

	this.RetuenJson(0, 0, "ok")
}

//测试pgsql连接
func (this *InstallController) InstallPgsqlTset() {
	host := this.GetString("host")
	user := this.GetString("user")
	password := this.GetString("password")
	dbname := this.GetString("dbname")
	log.Println(host, user, password, dbname)

	if host == "" || user == "" || password == "" || dbname == "" {
		this.RetuenJson(1, 0, "pgsql连接信息不全")
	}

	//pgsql 测试
	isconn, err := conn.PgsqlConnTest(host, user, password, dbname)
	if isconn {
		this.RetuenJson(0, 0, "连接成功")
		return
	}
	this.RetuenJson(1, 0, err)
}

//pgsql
func (this *InstallController) InstallPgsql() {
	host := this.GetString("host")
	user := this.GetString("user")
	password := this.GetString("password")
	dbname := this.GetString("dbname")
	log.Println(host, user, password, dbname)

	if host == "" || user == "" || password == "" || dbname == "" {
		this.RetuenJson(1, 0, "pgsql连接信息不全")
	}

	//pgsql 测试
	isconn, _ := conn.PgsqlConnTest(host, user, password, dbname)
	if !isconn {
		this.RetuenJson(1, 0, "连接不上pgsql.")
		return
	}

	//将信息写入 install.json
	infodata := object.OpenInstallFile()
	infodata.Step = 3
	infodata.DBType = "pgsql"
	infodata.PgsqlHost = host
	infodata.PgsqlUser = user
	infodata.PgsqlPassword = password
	infodata.PgsqlDBName = dbname
	object.GlobalDBType = "pgsql"
	object.GlobalPgsqlHost = host
	object.GlobalPgsqlUser = user
	object.GlobalPgsqlPassword = password
	object.GlobalPgsqlDBName = dbname

	installInfo, err := json.Marshal(infodata)
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	err = object.WriteInstallInfo(string(installInfo))
	if err != nil {
		this.RetuenJson(1, 0, err)
	}

	this.RetuenJson(0, 0, "ok")
}

//sqlite
func (this *InstallController) InstallSqlite() {
	isconn, err := conn.CreateSqliteDB()
	log.Println("sqlite 连接 = ", isconn, err)
	if isconn {

		//将信息写入 install.json
		infodata := object.OpenInstallFile()
		infodata.Step = 3
		infodata.DBType = "sqlite"
		infodata.SqlitePath = "./db/base.db"
		object.GlobalDBType = "sqlite"
		object.GlobalSqlitePath = "./db/base.db"

		installInfo, err := json.Marshal(infodata)
		if err != nil {
			this.RetuenJson(1, 0, err)
		}
		err = object.WriteInstallInfo(string(installInfo))
		if err != nil {
			this.RetuenJson(1, 0, err)
		}
		this.RetuenJson(0, 0, "创建成功")
		return
	}

	this.RetuenJson(1, 0, err)
}

//本地存储多媒体资源
//TODO  本地多媒体资源地址
func (this *InstallController) InstallLocalMedia() {
	path := this.GetString("path")

	if path == "" {
		this.RetuenJson(1, 0, "请输入")
	}

	//检查路径是否存在，不存在则创建
	err := util.CreateMutiDir(path)
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	infodata := object.OpenInstallFile()
	infodata.MediaType = "local"
	infodata.Step = 4
	infodata.MediaPath = path
	object.GlobalMediaType = "local"
	object.GlobalMediaPath = path

	installInfo, err := json.Marshal(infodata)
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	err = object.WriteInstallInfo(string(installInfo))
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	this.RetuenJson(0, 0, "成功")
}

//阿里云存储多媒体资源
func (this *InstallController) InstallAliyunMedia() {
	endPoint := this.GetString("oss_access_keyid")
	accessKeyID := this.GetString("oss_secret")
	ossSecret := this.GetString("oss_endpoint")
	bucketName := this.GetString("oss_bucketName")

	if endPoint == "" || accessKeyID == "" || ossSecret == "" || bucketName == "" {
		this.RetuenJson(1, 0, "请输入")
	}

	_, err := conn.TestConnAliOSS(endPoint, accessKeyID, ossSecret, bucketName)
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	infodata := object.OpenInstallFile()
	infodata.MediaType = "ali"
	infodata.Step = 4
	infodata.AliOSSAccessKeyid = accessKeyID
	infodata.AliOSSBucketName = bucketName
	infodata.AliOSSEndpoint = endPoint
	infodata.AliOSSSecret = ossSecret
	object.GlobalMediaType = "ali"
	object.GlobalAliOSSAccessKeyid = accessKeyID
	object.GlobalAliOSSBucketName = bucketName
	object.GlobalAliOSSEndpoint = endPoint
	object.GlobalAliOSSSecret = ossSecret

	installInfo, err := json.Marshal(infodata)
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	err = object.WriteInstallInfo(string(installInfo))
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	this.RetuenJson(0, 0, "成功")
}

//腾讯云存储多媒体资源
func (this *InstallController) InstallTencentMedia() {
	cos_url := this.GetString("cos_url")
	cos_secretid := this.GetString("cos_secretid")
	cos_secretkey := this.GetString("cos_secretkey")

	if cos_url == "" || cos_secretid == "" || cos_secretkey == "" {
		this.RetuenJson(1, 0, "请输入")
	}

	_, err := conn.TestConnTencentCOS(cos_url, cos_secretid, cos_secretkey)
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	infodata := object.OpenInstallFile()
	infodata.MediaType = "tencent"
	infodata.Step = 4
	infodata.TencenCosUrl = cos_url
	infodata.TencenSecretid = cos_secretid
	infodata.TencenSecretkey = cos_secretkey
	object.GlobalMediaType = "tencent"
	object.GlobalTencenCosUrl = cos_url
	object.GlobalTencenSecretid = cos_secretid
	object.GlobalTencenSecretkey = cos_secretkey

	installInfo, err := json.Marshal(infodata)
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	err = object.WriteInstallInfo(string(installInfo))
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	this.RetuenJson(0, 0, "成功")
}

//创建管理账号
func (this *InstallController) InstallAdmin() {
	account := this.GetString("account")
	password := this.GetString("password")
	phone := this.GetString("phone")
	mail := this.GetString("mail")

	if account == "" || password == "" {
		this.RetuenJson(1, 0, "账号或密码为空")
	}

	now := time.Now().Unix()
	admin := &models.ACC{
		UserId:     "admin",
		Account:    account,
		Password:   util.Md5Crypt(password, now),
		Phone:      phone,
		Mail:       mail,
		Createtime: now,
	}
	err := new(dao.DaoACC).CreateUser(admin)
	if err != nil {
		this.RetuenJson(1, 0, err)
	}

	infodata := object.OpenInstallFile()
	infodata.Step = -1
	installInfo, err := json.Marshal(infodata)
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	err = object.WriteInstallInfo(string(installInfo))
	if err != nil {
		this.RetuenJson(1, 0, err)
	}
	this.RetuenJson(0, 0, "成功")
}
