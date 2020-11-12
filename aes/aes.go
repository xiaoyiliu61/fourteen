package aes

import (
	"CryptCode/utils"
	"crypto/aes"
	"crypto/cipher"
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
