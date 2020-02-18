package util

/*
	常用web 信息处理的方法，如获取链接的title ico 处理链接等等
*/

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/imroc/req"
)

type WebInfo struct {}

//获取链接的ico 链接
func (this *WebInfo) GetIcoLink(url string) string {
	url1 := strings.Split(url, "//")
	url2 := strings.Split(url1[1], "/")
	return fmt.Sprintf("%s//%s/favicon.ico",url1[0],url2[0])
}

//获取链接的 Title
func (this *WebInfo) GetWebTitle(url string) (string,error) {
	authHeader := req.Header{}
	r, err := req.Get(url, authHeader, req.Header{"User-Agent": "V1.1"})
	if err != nil{
		return "",err
	}
	rtext := r.String()
	reg := regexp.MustCompile(`<title>(.*?)</title>`)
	matchs := reg.FindStringSubmatch(rtext)
	var data string;
	for _, s := range matchs {
        //fmt.Println(s)
        if s != ""{
        	data = s
        }
    }
    fmt.Println(data)
    return data,nil
}