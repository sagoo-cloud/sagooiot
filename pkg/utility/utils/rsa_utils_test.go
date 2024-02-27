package utils

import (
	"fmt"
	"sagooiot/internal/consts"
	"testing"
)

func TestRsaReq(t *testing.T) {

	/*err := GenerateKeyPair(2048, "../../resource/rsa/private.pem", "../../resource/rsa/public.pem")
	if err != nil {
		fmt.Println("生成RSA密钥对时发生错误:", err)
		return
	}*/

	message := "Zd-Gf3h8_r4"
	ciphertext, err := Encrypt("../../resource/rsa/public.pem", message, consts.RsaOAEP)
	if err != nil {
		fmt.Println("加密数据时发生错误:", err)
		return
	}

	fmt.Println("加密后的数据:", ciphertext)

	message = ciphertext
	plaintext, err := Decrypt("../../resource/rsa/private.pem", message, consts.RsaOAEP)
	if err != nil {
		fmt.Println("解密数据时发生错误:", err)
		return
	}

	fmt.Println("解密后的数据:", plaintext)
}
