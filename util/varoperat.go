package util

/*
	包含了变量的数据处理 
*/

import (
	"fmt"
	"regexp"
	"strings"
	"strconv" 
)


type Str struct {}

func (this *Str) Strip(s_ string, chars_ string) string {
	s , chars := []rune(s_) , []rune(chars_)
	length := len(s)
	max := len(s) - 1
	l, r := true, true //标记当左端或者右端找到正常字符后就停止继续寻找
	start, end := 0, max
	tmpEnd := 0
	charset := make(map[rune]bool) //创建字符集，也就是唯一的字符，方便后面判断是否存在
	for i := 0; i < len(chars); i++ {
		charset[chars[i]] = true
	}
	for i := 0; i < length; i++ {
		if _, exist := charset[s[i]]; l && !exist {
			start = i
			l = false
		}
		tmpEnd = max - i
		if _, exist := charset[s[tmpEnd]]; r && !exist {
			end = tmpEnd
			r = false
		}
		if !l && !r{
			break
		}
	}
	if l && r {  // 如果左端和右端都没找到正常字符，那么表示该字符串没有正常字符
		return ""
	}
	return string(s[start : end+1])
}

//匹配到MD文档里的第一个图片的链接，匹配第一个图片链接 返回最后匹配到的(.*)或(.*?)
func (this *Str) GetMDImgLink(txt string) string {
	reg := regexp.MustCompile(`\!\[.*\]\((.*?)\"|\!\[.*\]\((.*?)\)`)
	matchs := reg.FindStringSubmatch(txt)
	var data string;
	for _, s := range matchs {
        //fmt.Println(s)
        if s != ""{
        	data = s
        }
    }
    fmt.Println(data)
    return data
}

//返回笔记信息列表描述的内容，替换MD的特殊字符
func (this *Str) RepMDDesc(txt string,strlen int) string {
	txt = strings.Replace(txt, "#", "", -1 )
	txt = strings.Replace(txt, "*", "", -1 )
	txt = strings.Replace(txt, "~", "", -1 )
	txt = strings.Replace(txt, "~", "", -1 )
	txt = strings.Replace(txt, "`", "", -1 )
	txt = strings.Replace(txt, "-", "", -1 )
	txt = strings.Replace(txt, "\n", "   ", -1 )
	if len(txt) > strlen {
		//return txt[0:strlen]
		return ShowSubstr(txt,strlen)
	}else{
		//return txt[0:len(txt)]
		return ShowSubstr(txt,len(txt))
	}
	
}

//分割字符串，按指定长度
func ShowSubstr(s string, l int) string {
    if len(s) <= l {
        return s
    }
    ss, sl, rl, rs := "", 0, 0, []rune(s)
    for _, r := range rs {
        rint := int(r)
        if rint < 128 {
            rl = 1
        } else {
            rl = 2
        }
        if sl + rl > l {
            break
        }
        sl += rl
        ss += string(r)
    }
    return ss
}


//字符串空格转换html空格
func (this *Str) ToNbsp(txt string) string {
	return strings.Replace(txt, " ", "&nbsp;", -1 )
}

//数字字符串转int
func (this *Str) NumberToInt(number string) (int,error) {
	return strconv.Atoi(number)

}

//数组是否存在某元素  字符串
func (this *Str) IsItemStr(array []string, val string) bool{
	for _,k := range(array){
		if k==val{return true}
	}
	return false
}