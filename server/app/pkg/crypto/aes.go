package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"

	"edu/pkg/ecode"
)

func EncryptWithAES(key, message string) (string, error) {
	hash := md5.New()
	hash.Write([]byte(key))
	keyData := hash.Sum(nil)

	block, err := aes.NewCipher(keyData)
	if err != nil {
		return "", err
	}

	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

	enc := cipher.NewCBCEncrypter(block, iv)
	content := PKCS5Padding([]byte(message), block.BlockSize())
	crypted := make([]byte, len(content))
	enc.CryptBlocks(crypted, content)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func EncryptPrivWithAES(key, message string) (string, error) {
	hash := md5.New()
	hash.Write([]byte(key + "b^da6-d3*05fda!9"))
	keyData := hash.Sum(nil)

	block, err := aes.NewCipher(keyData)
	if err != nil {
		return "", err
	}

	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

	enc := cipher.NewCBCEncrypter(block, iv)
	content := PKCS5Padding([]byte(message), block.BlockSize())
	crypted := make([]byte, len(content))
	enc.CryptBlocks(crypted, content)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func DecryptWithAES(key, message string) (string, error) {
	hash := md5.New()
	hash.Write([]byte(key))
	keyData := hash.Sum(nil)

	block, err := aes.NewCipher(keyData)
	if err != nil {
		return "", err
	}

	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

	messageData, _ := base64.StdEncoding.DecodeString(message)
	dec := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(messageData))
	dec.CryptBlocks(decrypted, messageData)
	return string(PKCS5Unpadding(decrypted)), nil
}

// DecryptPrivWithAES 解压
func DecryptPrivWithAES(key, message string) (encry string, err error) {
	hash := md5.New()
	// code, err := GetCode()
	// if err != nil {
	// 	return
	// }
	// hash.Write([]byte(key + "b^da6-d3*05fda!9" + code))
	//hash.Write([]byte(key+"b^da6-d3*05fda!9"+"aaaaaaaa"))
	keyData := hash.Sum(nil)

	block, err := aes.NewCipher(keyData)
	if err != nil {
		err = ecode.WrapError(err)
		return
	}

	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

	messageData, _ := base64.StdEncoding.DecodeString(message)
	dec := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(messageData))
	dec.CryptBlocks(decrypted, messageData)
	encry = string(PKCS5Unpadding(decrypted))
	return
}

func DecryptPrivWithAES1(key, message string) (d string, err error) {
	hash := md5.New()
	hash.Write([]byte(key + "b^da6-d3*05fda!9"))
	keyData := hash.Sum(nil)

	block, err := aes.NewCipher(keyData)
	if err != nil {
		err = ecode.WrapError(err)
		return
	}

	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

	messageData, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return
	}
	dec := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(messageData))
	dec.CryptBlocks(decrypted, messageData)
	d = string(PKCS5Unpadding(decrypted))
	return
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Unpadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
