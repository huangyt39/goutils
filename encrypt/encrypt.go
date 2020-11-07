package encrypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func DesEncryption(key, plainText []byte) ([]byte, error) {
	if len(key) > 8 {
		key = key[:8]
	} else if len(key) < 8 {
		key = PKCS5Padding(key, 8)
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData := PKCS5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key)
	encryptBytes := make([]byte, len(origData))
	blockMode.CryptBlocks(encryptBytes, origData)
	return encryptBytes, nil
}

func DesDecryption(key, cipherText []byte) ([]byte, error) {
	if len(key) > 8 {
		key = key[:8]
	} else if len(key) < 8 {
		key = PKCS5Padding(key, 8)
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// 用余数填充
func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padBytes := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padBytes...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
}
