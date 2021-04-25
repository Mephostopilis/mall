package fastdfs

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego/httplib"
)

const (
	CONST_SMALL_FILE_NAME          = "small.txt"
	CONST_BIG_FILE_NAME            = "big.txt"
	CONST_DOWNLOAD_BIG_FILE_NAME   = "big_dowload.txt"
	CONST_DOWNLOAD_SMALL_FILE_NAME = "small_dowload.txt"
)

func TestUpload(t *testing.T) {
	obj, err := Upload("192.168.21.95", "C:\\Users\\Admin\\Pictures\\1.jpg", "1.png")
	if err != nil {
		t.Errorf("err = %v", err)
	}
	t.Logf("url = %v", obj.Url)
}

func TestUploadSmall(t *testing.T) {
	var obj FileResult
	req := httplib.Post("http://192.168.21.95:8080/group1" + "/upload")
	req.PostFile("file", "logo.png")
	req.Param("output", "json")
	req.Param("scene", "")
	req.Param("path", "")
	req.ToJSON(&obj)
	fmt.Println(obj.Url)
	// if obj.Md5 != testSmallFileMd5 {
	// 	t.Error("file not equal")
	// } else {
	// 	req = httplib.Get(obj.Url)
	// 	req.ToFile(CONST_DOWNLOAD_SMALL_FILE_NAME)
	// 	if md5sum, err := testUtil.GetFileSumByName(CONST_DOWNLOAD_SMALL_FILE_NAME, ""); md5sum != testSmallFileMd5 {
	// 		t.Error("small file not equal", err)
	// 	}
	// }
}
