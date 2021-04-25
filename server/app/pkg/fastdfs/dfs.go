package fastdfs

import (
	"fmt"

	"github.com/astaxie/beego/httplib"
)

type FileResult struct {
	Url     string `json:"url"`
	Md5     string `json:"md5"`
	Path    string `json:"path"`
	Domain  string `json:"domain"`
	Scene   string `json:"scene"`
	Size    int64  `json:"size"`
	ModTime int64  `json:"mtime"`
	//Just for Compatibility
	Scenes  string `json:"scenes"`
	Retmsg  string `json:"retmsg"`
	Retcode int    `json:"retcode"`
	Src     string `json:"src"`
}

func Upload(ip string, file string, args ...string) (obj FileResult, err error) {
	url := fmt.Sprintf("http://%s:8080/group1/upload", ip)
	req := httplib.Post(url)
	req.PostFile("file", file)
	req.Param("output", "json")
	req.Param("scene", "")
	req.Param("path", "")
	if err = req.ToJSON(&obj); err != nil {
		return
	}
	return
}
