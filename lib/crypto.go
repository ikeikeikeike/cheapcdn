package lib

import (
	"crypto/cipher"
	"encoding/hex"
)

var (
	aesBlock    cipher.Block
	aesCommonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
)

// EncryptAexHex encodes string to decoded.
func EncryptAexHex(src []byte) []byte {
	edst := make([]byte, len(src))
	cipher.NewCFBEncrypter(aesBlock, aesCommonIV).XORKeyStream(edst, src)

	hdst := make([]byte, hex.EncodedLen(len(edst)))
	hex.Encode(hdst, src)
	return hdst
}

// EncryptStringAexHex encodes string to decoded.
func EncryptStringAexHex(data string) string {
	src := []byte(data)
	dst := make([]byte, len(src))

	cipher.NewCFBEncrypter(aesBlock, aesCommonIV).XORKeyStream(dst, src)
	return hex.EncodeToString(dst)
}

// DecryptStringAexHex decodes decoded string
func DecryptStringAexHex(data string) string {
	dec, _ := hex.DecodeString(data)

	src := []byte(dec)
	dst := make([]byte, len(src))

	cipher.NewCFBDecrypter(aesBlock, aesCommonIV).XORKeyStream(dst, src)
	return string(dst)
}
