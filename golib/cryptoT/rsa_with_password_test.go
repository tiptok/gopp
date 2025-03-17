package cryptoT

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
)

// GenerateKeyPairWithPass 生成 RSA 密钥对，并对私钥进行加密
func GenerateKeyPairWithPass(bits int, password []byte) ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	// 对私钥进行加密
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	encryptedBlock, err := x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, password, x509.PEMCipherAES256)
	if err != nil {
		return nil, nil, err
	}
	privateKeyPEM := pem.EncodeToMemory(encryptedBlock)

	// 处理公钥
	publicKey := &privateKey.PublicKey
	publicKeyDER, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyDER,
	}
	publicKeyPEMBytes := pem.EncodeToMemory(publicKeyPEM)

	return privateKeyPEM, publicKeyPEMBytes, nil
}

// EncryptWithPublicKeyPEM 使用公钥加密数据
func EncryptWithPublicKeyPEM(plaintext []byte, publicKeyPEM []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil {
		return nil, fmt.Errorf("无法解码公钥 PEM 块")
	}
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, plaintext, nil)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// DecryptWithPrivateKeyWithPass 使用带密码的私钥解密数据
func DecryptWithPrivateKeyWithPass(ciphertext []byte, privateKeyPEM []byte, password []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		return nil, fmt.Errorf("无法解码私钥 PEM 块")
	}
	decryptedBlock, err := x509.DecryptPEMBlock(block, password)
	if err != nil {
		return nil, err
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(decryptedBlock)
	if err != nil {
		return nil, err
	}

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func TestRsaPass(t *testing.T) {
	// 生成带密码的 RSA 密钥对
	privateKeyPEM, publicKeyPEM, err := GenerateKeyPairWithPass(2048, []byte("yourpassword"))
	if err != nil {
		fmt.Println("密钥生成失败:", err)
		return
	}

	fmt.Println("私钥 (PEM 格式):")
	fmt.Println(string(privateKeyPEM))
	fmt.Println("公钥 (PEM 格式):")
	fmt.Println(string(publicKeyPEM))

	// 待加密的明文
	plaintext := []byte("Hello, RSA encryption with password!")

	// 使用公钥加密
	ciphertext, err := EncryptWithPublicKeyPEM(plaintext, publicKeyPEM)
	if err != nil {
		fmt.Println("加密失败:", err)
		return
	}
	fmt.Println("加密后的密文:", ciphertext)

	// 使用带密码的私钥解密
	decryptedText, err := DecryptWithPrivateKeyWithPass(ciphertext, privateKeyPEM, []byte("yourpassword"))
	if err != nil {
		fmt.Println("解密失败:", err)
		return
	}
	fmt.Println("解密后的明文:", string(decryptedText))
}

func TestRsaPassPubEncrypt(t *testing.T) {
	result, err := EncryptWithPublicKeyPEM([]byte("18860183050"), []byte("-----BEGIN public key-----\nMDwwDQYJKoZIhvcNAQEBBQADKwAwKAIhALWmbq2aCTM9RsZIk25bXo5BM70YnRZc\n7vtuaKiY0C6jAgMBAAE=\n-----END public key-----"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(result))
}
