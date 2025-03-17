package cryptoT

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
)

// GenerateKeyPair 生成 RSA 密钥对
func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// EncryptWithPublicKey 使用公钥加密数据
func EncryptWithPublicKey(plaintext []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// DecryptWithPrivateKey 使用私钥解密数据
func DecryptWithPrivateKey(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// PrivateKeyToPEM 将私钥转换为 PEM 格式
func PrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	return pem.EncodeToMemory(privateKeyPEM)
}

// PublicKeyToPEM 将公钥转换为 PEM 格式
func PublicKeyToPEM(publicKey *rsa.PublicKey) ([]byte, error) {
	publicKeyDER, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyDER,
	}
	return pem.EncodeToMemory(publicKeyPEM), nil
}

func TestRSA(t *testing.T) {
	// 生成 RSA 密钥对
	privateKey, publicKey, err := GenerateKeyPair(2048)
	if err != nil {
		fmt.Println("密钥生成失败:", err)
		return
	}

	// 将密钥转换为 PEM 格式
	privateKeyPEM := PrivateKeyToPEM(privateKey)
	publicKeyPEM, err := PublicKeyToPEM(publicKey)
	if err != nil {
		fmt.Println("公钥转换为 PEM 格式失败:", err)
		return
	}

	fmt.Println("私钥 (PEM 格式):")
	fmt.Println(string(privateKeyPEM))
	fmt.Println("公钥 (PEM 格式):")
	fmt.Println(string(publicKeyPEM))

	// 待加密的明文
	plaintext := []byte("Hello, RSA encryption!")

	// 使用公钥加密
	ciphertext, err := EncryptWithPublicKey(plaintext, publicKey)
	if err != nil {
		fmt.Println("加密失败:", err)
		return
	}
	fmt.Println("加密后的密文:", ciphertext)

	// 使用私钥解密
	decryptedText, err := DecryptWithPrivateKey(ciphertext, privateKey)
	if err != nil {
		fmt.Println("解密失败:", err)
		return
	}
	fmt.Println("解密后的明文:", string(decryptedText))
}
