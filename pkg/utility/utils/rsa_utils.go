package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"

	"github.com/gogf/gf/v2/errors/gerror"
)

// GenerateKeyPair 生成RSA密钥对
func GenerateKeyPair(bits int, privateKeyFile string, publicKeyFile string) (err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		err = gerror.New("生成RSA密钥对时发生错误")
		return
	}

	err = SavePrivateKeyToFile(privateKey, privateKeyFile)
	if err != nil {
		err = gerror.New("保存私钥文件时发生错误")
		return
	}

	publicKey := &privateKey.PublicKey
	err = SavePublicKeyToFile(publicKey, publicKeyFile)
	if err != nil {
		err = gerror.New("保存公钥文件时发生错误")
		return
	}

	return
}

// EncryptPKCS1v15 使用公钥加密数据
func EncryptPKCS1v15(publicKeyFile string, content string) (ciphertext string, err error) {
	publicKeyBytes, err := os.ReadFile(publicKeyFile)
	if err != nil {
		err = gerror.New("读取公钥文件时发生错误:" + err.Error())
		return
	}

	publicKey, err := ParsePublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		err = gerror.New("解析公钥时发生错误:" + err.Error())
		return
	}

	ciphertextBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(content))
	if err != nil {
		err = gerror.New("加密数据时发生错误:" + err.Error())
		return
	}
	ciphertext = base64.StdEncoding.EncodeToString(ciphertextBytes)
	return
}

// DecryptPKCS1v15 使用私钥解密数据
func DecryptPKCS1v15(privateKeyFile, ciphertext string) (plaintext string, err error) {
	privateKeyBytes, err := os.ReadFile(privateKeyFile)
	if err != nil {
		err = gerror.New("读取私钥文件时发生错误:" + err.Error())
		return
	}

	privateKey, err := ParsePrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		err = gerror.New("解析私钥时发生错误:" + err.Error())
		return
	}
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return
	}
	plaintextBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertextBytes)
	if err != nil {
		err = gerror.New("解密数据时发生错误:" + err.Error())
		return
	}
	plaintext = string(plaintextBytes)
	return
}

// EncryptOAEP 使用公钥加密数据
func EncryptOAEP(publicKeyFile string, content string) (ciphertext string, err error) {
	publicKeyBytes, err := os.ReadFile(publicKeyFile)
	if err != nil {
		err = gerror.New("读取公钥文件时发生错误:" + err.Error())
		return
	}

	publicKey, err := ParsePublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		err = gerror.New("解析公钥时发生错误:" + err.Error())
		return
	}

	ciphertextBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(content), nil)
	if err != nil {
		err = gerror.New("加密数据时发生错误:" + err.Error())
		return
	}
	ciphertext = base64.StdEncoding.EncodeToString(ciphertextBytes)
	return
}

// DecryptOAEP 使用私钥解密数据
func DecryptOAEP(privateKeyFile, ciphertext string) (plaintext string, err error) {
	privateKeyBytes, err := os.ReadFile(privateKeyFile)
	if err != nil {
		err = gerror.New("读取私钥文件时发生错误:" + err.Error())
		return
	}

	privateKey, err := ParsePrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		err = gerror.New("解析私钥时发生错误:" + err.Error())
		return
	}
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return
	}
	plaintextBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertextBytes, nil)
	if err != nil {
		err = gerror.New("解密数据时发生错误:" + err.Error())
		return
	}
	plaintext = string(plaintextBytes)
	return
}

// SavePrivateKeyToFile 将 RSA 私钥保存到文件
func SavePrivateKeyToFile(privateKey *rsa.PrivateKey, filename string) (err error) {
	keyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return
	}
	privateBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: keyBytes,
	}
	privatePem := pem.EncodeToMemory(privateBlock)
	err = os.WriteFile(filename, privatePem, 0600)
	if err != nil {
		err = gerror.New("保存私钥文件时发生错误:" + err.Error())
		return
	}
	return
}

// SavePublicKeyToFile  将 RSA 公钥保存到文件
func SavePublicKeyToFile(publicKey *rsa.PublicKey, filename string) (err error) {
	keyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: keyBytes,
	}
	publicPem := pem.EncodeToMemory(publicBlock)
	err = os.WriteFile(filename, publicPem, 0644)
	if err != nil {
		err = gerror.New("保存公钥文件时发生错误:" + err.Error())
		return
	}
	return
}

// ParsePrivateKeyFromPEM 从 PEM 格式的字节切片中解析 RSA 私钥
func ParsePrivateKeyFromPEM(pemBytes []byte) (privateKey *rsa.PrivateKey, err error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "PRIVATE KEY" {
		err = gerror.New("无效的私钥")
		return
	}
	PKCS8PrivateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		err = gerror.New("解析私钥时发生错误:" + err.Error())
		return
	}
	privateKey, ok := PKCS8PrivateKey.(*rsa.PrivateKey)
	if !ok {
		err = gerror.New("解析私钥时发生错误:invalid public key type")
		return
	}
	return
}

// ParsePublicKeyFromPEM 从 PEM 格式的字节切片中解析 RSA 公钥
func ParsePublicKeyFromPEM(pemBytes []byte) (publicKey *rsa.PublicKey, err error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
		err = gerror.New("无效的公钥")
		return
	}

	pKIXPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		err = gerror.New("解析公钥时发生错误:" + err.Error())
		return
	}

	publicKey, ok := pKIXPublicKey.(*rsa.PublicKey)
	if !ok {
		err = gerror.New("解析公钥时发生错误:invalid public key type")
		return
	}

	return
}

// Decrypt 使用私钥解密数据
func Decrypt(privateKeyFile, ciphertext string, types string) (plaintext string, err error) {
	switch types {
	case "OAEP":
		plaintext, err = DecryptOAEP(privateKeyFile, ciphertext)
		break
	case "PKCS1v15":
		plaintext, err = DecryptPKCS1v15(privateKeyFile, ciphertext)
		break
	default:
		plaintext, err = DecryptPKCS1v15(privateKeyFile, ciphertext)
		break
	}
	return
}

// Encrypt 使用公钥加密数据
func Encrypt(privateKeyFile, ciphertext string, types string) (plaintext string, err error) {
	switch types {
	case "OAEP":
		plaintext, err = EncryptOAEP(privateKeyFile, ciphertext)
		break
	case "PKCS1v15":
		plaintext, err = EncryptPKCS1v15(privateKeyFile, ciphertext)
		break
	default:
		plaintext, err = EncryptPKCS1v15(privateKeyFile, ciphertext)
		break
	}
	return
}
