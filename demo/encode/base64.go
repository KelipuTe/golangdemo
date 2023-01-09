package encode

import (
	"encoding/base64"
	"fmt"
)

func f8Base64Encrypt() {
	strBeforeBase64 := "string before base64"
	s5BeforeBase64 := []byte(strBeforeBase64)

	strAfterBase64 := base64.StdEncoding.EncodeToString(s5BeforeBase64)

	fmt.Printf("strBeforeBase64: %s\n", strBeforeBase64)
	fmt.Printf("strAfterBase64: %s\n", strAfterBase64)
}
