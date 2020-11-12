package des

import (
	"CryptCode/utils"
	"crypto/cipher"
	"crypto/des"
)

/*
使用DES算法对明文进行加密，使用密钥key
*/
func DESEnCrypt(data ,key []byte) ([]byte,error) {
	block,err:=des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	originData:=utils.PKCS5EndPadding(data,block.BlockSize())

	blockMode:=cipher.NewCBCEncrypter(block,key)
	cipherData:=make([]byte,len(originData))
	blockMode.CryptBlocks(cipherData,originData)
	return cipherData ,nil
}

/*
使用DES算法对密文进行解密，使用可以作为密钥
*/
func DESDeCrypt(data, key []byte) ([]byte, error) {
	block,err:=des.NewCipher(key)
	if err != nil {
		return nil,err
	}
    blockMode:=cipher.NewCBCDecrypter(block,key)
    originData:=make([]byte,len(data))
	blockMode.CryptBlocks(originData,data)
    utils.ClearPKCS5Padding(originData,block.BlockSize())
	return originData,nil
}
