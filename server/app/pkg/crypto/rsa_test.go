package crypto

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FileLoad(filepath string) (pb []byte, err error) {
	privatefile, err := os.Open(filepath)
	defer privatefile.Close()
	if err != nil {
		return
	}
	privateKey := make([]byte, 2048)
	num, err := privatefile.Read(privateKey)
	pb = privateKey[:num]
	return
}

func TestGen(t *testing.T) {
	publickey, err := FileLoad("./mypublic.pem")
	if !assert.NoError(t, err) {
		return
	}
	privatekey, err := FileLoad("./myprivatekey.pem")
	if !assert.NoError(t, err) {
		return
	}
	data, err := RSAEncrypt([]byte("QQ77025077"), publickey)
	if !assert.NoError(t, err) {
		return
	}
	fmt.Println("加密：", base64.StdEncoding.EncodeToString(data))
	origData, err := RSADecrypt(data, privatekey) //解密
	if !assert.NoError(t, err) {
		return
	}
	fmt.Println("解密:", string(origData))
}
