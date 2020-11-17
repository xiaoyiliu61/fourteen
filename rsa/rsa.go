package rsa

import (
	"CryptCode/utils"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"os"
)

func CreatePairKeys() (*rsa.PrivateKey,error) {
	//1.先生成私钥
	var bits int
	flag.IntVar(&bits,"b",1024,"密钥长度")

	privateKey,err:=rsa.GenerateKey(rand.Reader,bits)
	if err!=nil{
		return nil,err
	}
	//2.根据私钥生成公钥
	
	//3.将私钥和公钥进行返回
	return privateKey,nil
}

/*
根据给定的私钥数据生成对应的pem认证文件
*/
func GeneratePemFileByPrivateKey(pri *rsa.PrivateKey) (error) {
	//根据PKCS1规则，序列化后的私钥
	priStream:=x509.MarshalPKCS1PrivateKey(pri)
	privatFile,err:=os.Create("rsa_pri.pem")//存私钥生成文件
	if err != nil {
		return err
	}
	//pem
	block:=&pem.Block{
		Type:    "RSA PRIVATE KEY",
		Bytes:   priStream,
	}

	err=pem.Encode(privatFile,block)
	if err != nil {
		return err
	}
	return nil
}

func GeneratePubFileByPupKey(pub rsa.PublicKey) (error) {
	strem:=x509.MarshalPKCS1PublicKey(&pub)
	block:=pem.Block{
		Type:    "RSA PRIVATE KEY",
		Bytes:   strem,
	}
	pubFile,err:=os.Create("rsa_pub.pem")
	if err != nil {
		return err
	}
	return pem.Encode(pubFile,&block)
}

func GenerateKeys(file_name string) error {
	pri,err:=CreatePairKeys()
	if err != nil {
		return err
	}

}


/*
使用RSA算法对数据进行加密，返回加加密后的密文
*/
func RSAEncrypt(key rsa.PublicKey,data []byte) ([]byte,error)  {
   return rsa.EncryptPKCS1v15(rand.Reader,&key,data)

}

func RSADecrypt(private *rsa.PrivateKey, cipher []byte)([]byte,error)  {
	return rsa.DecryptPKCS1v15(rand.Reader,private,cipher)
}

//私钥签名，公钥验签
func RSASign(private *rsa.PrivateKey,data []byte,hash crypto.Hash) ([]byte,error)  {
	if hash == crypto.MD5 {
		hashed:=utils.Md5Hash(data)
		return rsa.SignPKCS1v15(rand.Reader,private,crypto.MD5,hashed)
	} else if hash == crypto.SHA256 {
		hashed:=utils.Sha256Hash(data)
		return rsa.SignPKCS1v15(rand.Reader,private,crypto.SHA256,hashed)
	}
	 
}

func RSAVerify(pub rsa.PrivateKey,data []byte, signText []byte) (bool,error) {
	hashed:=utils.Md5Hash(data)
	err:=rsa.VerifyPKCS1v15(&pub,crypto.MD5,hashed,signText)
	return err==nil,err
}

