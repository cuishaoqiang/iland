package encoder

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

var publicKey []byte
var privateKey []byte

func init(){
	var err error
	publicKey, err = getRsaKey("./config/rsa_key/rsa_public_key.pem")
	if err != nil {
		fmt.Println("## rsa public key read failed!")
	}

	privateKey, err = getRsaKey("./config/rsa_key/rsa_private_key.pem")
	if err != nil {
		fmt.Println("## rsa private key read failed!")
	}
}

func getRsaKey(path string) ([]byte, error) {
	config, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	return config, nil
}

func GetPublicKey() string {
	return string(publicKey)
}

// 加密
func RSAEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, fmt.Errorf("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

//RSA解密
func RSADecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, fmt.Errorf("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
