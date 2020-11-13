package aes

import (
	"CryptCode/utils"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

/*
使用AES算法对明文进行加密
*/
func AESEnCrypt(origin []byte,key []byte) ([]byte,error) {
	//3元素：key ，data ，mode
	block,err:=aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cryptData:=utils.PKCS5EndPadding(origin,block.BlockSize())

	blockMode:=cipher.NewCBCEncrypter(block,key)
	//加密
	cipherData:=make([]byte,len(cryptData))
	blockMode.CryptBlocks(cipherData,cryptData)
	return cipherData, nil
}

func AESDeCrypt(data []byte, key []byte) ([]byte, error) {
	block,err:=aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	/*blockMode:=cipher.NewCBCDecrypter(block,key)
	originData:=make([]byte,len(origin))
	blockMode.CryptBlocks(originData,origin)
	utils.ClearPKCS5Padding(originData,block.BlockSize())
	return originData,nil
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}*/
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	originData := make([]byte, len(data))
	blockMode.CryptBlocks(originData, data)
	originData = utils.ClearPKCS5Padding(originData, block.BlockSize())
	return originData, nil
}
