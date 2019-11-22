package main

import (
	"LogicFun"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"go-common/library/log"
	"strings"
)

var default_aes_key string = "wps_openplatform"

type MyError struct {
	msg string
}

func (this *MyError) Error() string {
	defer println("aaa")
	return fmt.Sprintf("MyError %s", this.msg)
}

func getError() *MyError {
	return nil
	// var a *MyError
	// return a
}

func Decrypt(source_str string) string {
	source_str = strings.Trim(source_str, " ")
	if len(source_str) == 0 {
		return ""
	}
	aes_key := default_aes_key
	if len(default_aes_key) > 0 {
		aes_key = default_aes_key
	}
	source_bytes, e := base64.RawStdEncoding.DecodeString(source_str)
	if e != nil {
		log.Info("base64.RawStdEncoding.Decode err: %s\n", e)
		log.Error("base64.RawStdEncoding.Decode err: %s", e)
		return source_str
	}
	dest_bytes, err := AesDecrypt(source_bytes, []byte(aes_key))
	if err != nil {
		log.Info("Security.Decrypt err: %s\n", err)
		log.Error("Security.Decrypt err: %s", err)
		return source_str
	}
	return string(dest_bytes)
}

func AesDecrypt(crypted []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	block_size := block.BlockSize()
	cnt := len(crypted)
	if cnt%block_size != 0 {
		return crypted, nil
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:block_size])
	origin_data := make([]byte, cnt)
	// origin_data := crypted
	blockMode.CryptBlocks(origin_data, crypted)
	origin_data = PKCS5UnPadding(origin_data)
	// origin_data = ZeroUnPadding(origin_data)
	return origin_data, nil
}

func PKCS5UnPadding(origin_data []byte) []byte {
	length := len(origin_data)
	if length < 1 {
		return origin_data
	}
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origin_data[length-1])
	return origin_data[:(length - unpadding)]
}

func main() {
	// str := "UA1+Uq5wwwKlsrok3LlrvZcI1c2WeSXb+Iy+SCxSd03wNK2LnpqVkFdRMvF1FIAk"
	// println(Decrypt(str))
	LogicFun.MaximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}})
}

/*
([][]byte{
		{'0', '1', '1', '0', '0', '1', '0', '1', '0', '1'},
		{'0', '0', '1', '0', '1', '0', '1', '0', '1', '0'},
		{'1', '0', '0', '0', '0', '1', '0', '1', '1', '0'},
		{'0', '1', '1', '1', '1', '1', '1', '0', '1', '0'},
		{'0', '0', '1', '1', '1', '1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0', '1', '1', '1', '1', '0'},
		{'0', '0', '0', '1', '1', '0', '0', '0', '1', '0'},
		{'1', '1', '0', '1', '1', '0', '0', '1', '1', '1'}})
*/
