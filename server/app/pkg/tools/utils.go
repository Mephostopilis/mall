package tools

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// CompareHashAndPassword 比较hash
func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetGID 获取当前协程id
func GetGID() (n uint64, err error) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return
	}
	return
}

// GetLocalNetDeviceNames 获取网卡名
func GetLocalNetDeviceNames() (names []string, err error) {
	baseNicPath := "/sys/class/net/"
	cmd := exec.Command("ls", baseNicPath)
	buf, err := cmd.Output()
	if err != nil {
		return
	}
	output := string(buf)
	str := ""
	for _, device := range strings.Split(output, "\n") {
		if len(device) > 1 {
			if device != "lo" {
				str += device + "|"
			}
		}
	}
	names = strings.Split(str[:len(str)-1], "|")
	return
}

func Hmac(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func IsStringEmpty(str string) bool {
	return strings.Trim(str, " ") == ""
}

func GetUUID() string {
	u := uuid.New()
	return strings.ReplaceAll(u.String(), "-", "")
}

// func PathExists(path string) bool {
// 	_, err := os.Stat(path)
// 	if err == nil {
// 		return true
// 	}

// 	if os.IsNotExist(err) {
// 		return false
// 	}

// 	return false
// }

func Base64ToImage(imageBase64 string) ([]byte, error) {
	image, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func GetDirFiles(dir string) ([]string, error) {
	dirList, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	filesRet := make([]string, 0)
	for _, file := range dirList {
		if file.IsDir() {
			files, err := GetDirFiles(dir + string(os.PathSeparator) + file.Name())
			if err != nil {
				return nil, err
			}

			filesRet = append(filesRet, files...)
		} else {
			filesRet = append(filesRet, dir+string(os.PathSeparator)+file.Name())
		}
	}
	return filesRet, nil
}

func GetCurrentTimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}

//slice去重
func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

func IdsStrToIdsIntGroup(keys string) []int {
	IDS := make([]int, 0)
	ids := strings.Split(keys, ",")
	for i := 0; i < len(ids); i++ {
		ID, _ := StringToInt(ids[i])
		IDS = append(IDS, ID)
	}
	return IDS
}
