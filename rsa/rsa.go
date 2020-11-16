package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"flag"
)

func CreatePairKeys() (*rsa.PrivateKey,crypto.PrivateKey,error) {
	//1.先生成私钥
	var bits int
	flag.IntVar(&bits,"b",2048,"密钥长度")

	privateKey,err:=rsa.GenerateKey(rand.Reader,bits)
	if err!=nil{
		return nil,nil,err
	}
	//2.根据私钥生成公钥
    publicKey:=privateKey.Public()
	//3.将私钥和公钥进行返回
	return privateKey,publicKey,nil
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

func RSASign(private *rsa.PrivateKey,data []byte) ([]byte,error)  {
	return rsa.SignPKCS1v15(rand.Reader,private,data)
}
