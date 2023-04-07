module github.com/mangenotwork/ManGe-Notes

go 1.13

replace github.com/mangenotwork/ManGe-Notes => ./

require (
	github.com/aliyun/aliyun-oss-go-sdk v2.1.5+incompatible
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/garyburd/redigo v1.6.2
	github.com/imroc/req v0.3.0
	github.com/jinzhu/gorm v1.9.16
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/rs/xid v1.2.1
	github.com/tencentyun/cos-go-sdk-v5 v0.7.17
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
)
