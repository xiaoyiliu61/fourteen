package base

import "encoding/base64"

func Base64Encode(data []byte) string {
	encoding:=base64.StdEncoding
	dst:=make([]byte,0)

	encoding.Encode(dst,data)
	return string(dst)
}

func Base64Decode(data string) string {
	encoding:=base64.StdEncoding
	dst:=make([]byte,0)
	encoding.Decode(dst,[]byte(data))
	return string(dst)
}