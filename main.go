package main

import (
	"CryptCode/3des"
	"CryptCode/aes"
	"CryptCode/des"
	"CryptCode/ecc"
	"bytes"
	"fmt"
)

func main() {
	/*key:=[]byte("C6190604")

	data:="遇贵人先立业，遇良人先成家."*/

	//加密crypto
	/*block,err:=des.NewCipher(key)
	if err != nil {
		panic("初始化密钥错误，请重试!")
	}
	//dst,src
	dst:=make([]byte,len([]byte(data)))
	//加密过程
	block.Encrypt(dst,[]byte(data))

	fmt.Println("密文：",dst)

	//解密
	deBlock,err:=des.NewCipher(key)
	if err != nil {
		panic("初始化密钥错误，请重试!")
	}
	deData:=make([]byte,len(dst))
	deBlock.Decrypt(deData,dst)

	fmt.Println(string(deData))*/
	//1.得到cipher
	/*block,err:=des.NewCipher(key)
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	//2.对数据明文进行结尾块填充
	paddingData:=EndPadding([]byte(data),block.BlockSize())

	//3.选择模式
	mode:=cipher.NewCBCEncrypter(block,key)

	//4.加密
	dstData:=make([]byte,len(paddingData))
	mode.CryptBlocks( dstData,paddingData)

	fmt.Println("加密后的内容：",dstData)

	/*
	二、对数据进行解密
	DES三元素：key，data，mode
	*/
   /* key1:=[]byte("C6190604")
    data1:=dstData
    block1,err:=des.NewCipher(key1)
	if err!=nil {
		panic(err.Error())
	}
    deMode:=cipher.NewCBCDecrypter(block1,key1)
    originalData:=make([]byte,len(data1))
    //分组解密
    deMode.CryptBlocks(originalData,data1)
    originalData=utils.PKCS5EndPadding(data1,block.BlockSize())
    fmt.Println("解密后的内容：",string(originalData))*/

	key:=[]byte("20201112")
	data:="窗含西岭千秋雪，门泊东吴万里船"
	cipherText,err:=des.DESEnCrypt([]byte(data),key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	originText,err:=des.DESDeCrypt(cipherText,key)
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("DES解密后:",string(originText))

	//二、3DES加解密
	key1 := []byte("202011122020111220201112") //3des的秘钥长度是24字节
	data1 := "穷在闹市无人问，富在深山有远亲"
	cipherText1, err := _des.TripleDESEncrypt([]byte(data1), key1)
	if err != nil {
		fmt.Println("3DES加密失败:", err.Error())
		return
	}
	originalText1, err := _des.TripleDESDecrypt(cipherText1, key1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("3DES解密后的内容：", string(originalText1))

	//三：AES算法
	fmt.Println("AES算法：")
	key2 := []byte("2020111220201112") //8
	data2 := "人生在世不称意，明朝散发弄扁舟"

	cipherText2, err := aes.AESEnCrypt([]byte(data2), key2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(cipherText2)
   /* originText2,err:=aes.AESDeCrypt(cipherText2,key2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("aes解密后的内容：",string(originText2))*/

   //四：RSA算法
/*   data4:="把悲伤留给自己"
   priv,err:=rsa.CreatePairKeys()
   if err!=nil{
   	fmt.Println("生成秘钥对出错",err.Error())
	   return
   }
   rsa.RSAEncrypt(priv.PublicKey,[]byte(data))
   //将生成的私钥保存到硬盘上一个pem文件中，
   /* err =rsa.GeneratePemFileByPrivateKey(priv)
	if err != nil {
		fmt.Println("生成私钥的pem文件遇到错误：",err.Error())
	}
	err=rsa.GeneratePubFileByPupKey(priv.PublicKey)
	if err != nil {
		fmt.Println("生成公钥的pem文件遇到错误：",err.Error())
		return
	}
*/
 /*  err=rsa.GenerateKeys("huang")
	if err != nil {
		fmt.Println("生成私钥证书失败：",err.Error())
	}

   cipherText4,err:=rsa.RSAEncrypt(priv.PublicKey,[]byte(data4))
	if err != nil {
		fmt.Println("rsa算法加密失败",err.Error())
		return
	}
	originalText4,err:=rsa.RSADecrypt(priv,cipherText4)
	if err != nil {
		fmt.Println("rsa算法解密失败",err.Error())
		return
	}
	fmt.Println("rsa解密成功，结果是",string(originalText4))
    //对源文件数据进行签名
    signText4,err:=rsa.RSASign(priv.PublicKey,[]byte(data4))
	if err != nil {
		fmt.Println("rsa算法签名失败：",err.Error())
		return
	}
	fmt.Println("",string(signText4))*/

	fmt.Println("==================椭圆曲线数字签名算法===================")
	//① 生成秘钥
	pri, err := ecc.GenerateECDSAKey()
	if err != nil {
		fmt.Println("生成ECDSA秘钥对失败:", err.Error())
		return
	}
	//② 准备数据
	data6 := "张华考上了北京大学,李萍进了软件职业技术学校,我在百货公司当售货员,我们都有光明的前途和未来"

	//③ 数字签名
	r, s, err := ecc.ECDSASign(pri, []byte(data6))
	fmt.Printf("%x\n", r)
	fmt.Printf("%x\n", s)
	fmt.Println(r)
	fmt.Println(s)
	//pem格式：密钥证书文件的格式
	//der格式：

	//④ 数字签名验证
	data6 = "他开着比克，你提了林肯，我在拼多多砍单自行车成功，我们都有的未来"
	verifyResult := ecc.ECDSAVerify(pri.PublicKey, []byte(data6), r, s)
	if verifyResult {
		fmt.Println("ECDSA数字签名验证成功")
	} else {
		fmt.Println("ECDSA数字签名验证失败")
	}
}


/*
明文数据尾部填充
*/
func PKCS5Padding(text []byte,blockSize int) []byte  {
    //计算要填充的块内容的大小
	paddingSize:=blockSize-len(text)%blockSize
	paddingText:=bytes.Repeat([]byte{byte(paddingSize)},paddingSize)
	return append(text,paddingText...)
}
