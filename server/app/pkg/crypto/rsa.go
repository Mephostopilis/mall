package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	ErrBadPublicKey = errors.ResourceExhausted("res", "public key is bad")
)

func RSAEncrypt(orgidata []byte, publickey []byte) (pb []byte, err error) {
	block, _ := pem.Decode(publickey)
	if block == nil {
		err = ErrBadPublicKey
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	pub := pubInterface.(*rsa.PublicKey)
	pb, err = rsa.EncryptPKCS1v15(rand.Reader, pub, orgidata) //加密
	return
}

func RSADecrypt(cipertext, privatekey []byte) (pb []byte, err error) {
	block, _ := pem.Decode(privatekey)
	if block == nil {
		err = ErrBadPublicKey
		return
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	pb, err = rsa.DecryptPKCS1v15(rand.Reader, priv, cipertext)
	return
}
