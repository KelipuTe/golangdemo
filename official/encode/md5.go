package encode

import (
	"crypto/md5"
	"fmt"
)

func f8MD5Encrypt() {
	strBeforeMD5 := "string before md5"
	s5BeforeMD5 := []byte(strBeforeMD5)

	i9md5 := md5.New()
	i9md5.Write(s5BeforeMD5)
	s5AfterMD5 := i9md5.Sum(nil)

	fmt.Printf("strBeforeMD5: %s\n", strBeforeMD5)
	fmt.Printf("s5AfterMD5: %v\n", s5AfterMD5)
}
