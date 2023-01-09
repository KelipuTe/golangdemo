package encode

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func f8SHA256Encrypt() {
	strSecret := "secret"
	s5Secret := []byte(strSecret)
	strBeforeHmacSHA256 := "string before hmac sha256"
	s5BeforeHmacSHA256 := []byte(strBeforeHmacSHA256)

	i9HmacSHA256 := hmac.New(sha256.New, s5Secret)
	i9HmacSHA256.Write(s5BeforeHmacSHA256)
	s5AfterHmacSHA256 := i9HmacSHA256.Sum(nil)

	fmt.Printf("strSecret: %s\n", strSecret)
	fmt.Printf("strBeforeHmacSHA256: %s\n", strBeforeHmacSHA256)
	fmt.Printf("s5AfterHmacSHA256: %v\n", s5AfterHmacSHA256)
}
