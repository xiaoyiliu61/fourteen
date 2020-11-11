package utils

import "bytes"

func PKCS5EndPadding(text []byte, size int) []byte {
	paddingSize:=size-len(text)/size
	paddingText:=bytes.Repeat([]byte{(byte(paddingSize))},paddingSize)
	return append(text,paddingText...)
}
func ZeroEndPadding(text []byte, size int) []byte {
	paddingSize:=size-len(text)/size
	paddingText:=bytes.Repeat([]byte{(byte(0))},paddingSize)
	return append(text,paddingText...)
}

