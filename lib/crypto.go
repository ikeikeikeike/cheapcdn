package lib

import (
	"crypto/cipher"
)

var (
	aesBlock    cipher.Block
	aesCommonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
)

// EncryptAES encodes enctypt string
func EncryptAES(data string) []byte {
	src := []byte(data)
	dst := make([]byte, len(src))

	cipher.NewCFBEncrypter(aesBlock, aesCommonIV).XORKeyStream(dst, src)
	return dst
}

// DecryptAES decodes desript string
func DecryptAES(data string) []byte {
	src := []byte(data)
	dst := make([]byte, len(src))

	cipher.NewCFBDecrypter(aesBlock, aesCommonIV).XORKeyStream(dst, src)
	return dst
}
