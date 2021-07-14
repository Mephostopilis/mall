package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"edu/pkg/ecode"
	"encoding/pem"
)

func RSAEncrypt(orgidata []byte, publickey []byte) (pb []byte, err error) {
	block, _ := pem.Decode(publickey)
	if block == nil {
		err = ecode.ErrBadPublicKey("public key [%s] is null", string(publickey))
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
		err = ecode.ErrBadPublicKey("public key [%s] is null", string(privatekey))
		return
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	pb, err = rsa.DecryptPKCS1v15(rand.Reader, priv, cipertext)
	return
}
