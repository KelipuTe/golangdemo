package main

import (
	"errors"
	"fmt"
)

func main() {
	for _, strIn := range []string{"hello", "", "my error"} {
		fmt.Printf("strIn: %s\n", strIn)
		strEcho, err := echo(strIn)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("strEcho: %s\n", strEcho)
	}
}

func echo(strIn string) (string, error) {
	var err error
	if strIn == "" {
		err = errors.New("empty str")
		return "", err
	} else if strIn == "my error" {
		return "", &MyError{10000, errors.New("my error")}
	}
	return strIn, nil
}

// 自定义错误
type MyError struct {
	No  int
	Err error
}

// 实现error接口
func (e *MyError) Error() string {
	return fmt.Sprintf("No:%d,Err:%s", e.No, e.Err.Error())
}
