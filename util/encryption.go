package util

/*
	集成了加密
*/

import (
	"fmt"
	"crypto/md5"
	"strings"
	"time"
)

func Md5SaltCrypt(str string) (CryptStr string) {
	salt := fmt.Sprintf("%d",time.Now().Unix()) 
	slice := make([]string, len(salt)+1)
	str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt)
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
