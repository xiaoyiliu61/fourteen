package _des

import (
	"CryptCode/utils"
	"crypto/cipher"
	"crypto/des"
)

/*
该函数用于实现3des算法的加密
*/
func TripleDesEncrypt(origintext ,key []byte) ([]byte,error)  {
 //1.实例化一个cipher
  block,err:=des.NewTripleDESCipher(key)
	if err != nil {
		return nil,err
	}
	//对明文进行尾部填充
	cryptData:=utils.PKCS5EndPadding(origintext,block.BlockSize())
	//3.实例化加密模式mode
	mode:=cipher.NewCBCEncrypter(block,key)
	//对填充后的明文进行分组加密
    cipherText:=make([]byte,len(cryptData))
	mode.CryptBlocks(cipherText,cryptData)
    return cipherText,nil
}

/*
该函数用于密文进行解密，并返回明文
*/
func TripleDESDecrypt(ciphertext []byte, key []byte) ([]byte,error) {
   block,err:=des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	//2.不需要对密文进行尾部填充，
	blockMoed:=cipher.NewCBCDecrypter(block,key)
    originText:=make([]byte,len(ciphertext))
    blockMoed.CryptBlocks(ciphertext,originText)
    return originText,nil
}

/*
该函数将对明文进行尾部填充，采用PKCS5方式

func PKCS5EndPadding(text []byte, size int) []byte {
	paddingSize:=size-len(text)/size
	paddingText:=bytes.Repeat([]byte{(byte(paddingSize))},paddingSize)
	return append(text,paddingText...)
}
func ZeroEndPadding(text []byte, size int) []byte {
	paddingSize:=size-len(text)/size
	paddingText:=bytes.Repeat([]byte{(byte(0))},paddingSize)
	return append(text,paddingText...)
}*/
