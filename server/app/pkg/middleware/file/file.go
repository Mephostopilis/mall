package file

import (
	"context"
	"io"
	"os"
	"strconv"

	pb "edu/api/sys/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang/protobuf/ptypes"
)

// FILE 获取文件数据
func FILE(logger log.Logger) middleware.Middleware {
	log := log.NewHelper("middleware/file", logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			h, ok := transporthttp.FromServerContext(ctx)
			if ok {
				if h.Request.URL.Path == "" {
					ctx, err := uploadfile(ctx, log, h)
					if err != nil {
						return nil, err
					}
					return handler(ctx, req)
					// } else if strings.Contains(h.Request.URL.Path, "/admin/file/v1/cert/download/") {
					// 	ret, err := handler(ctx, req)
					// 	if err != nil {
					// 		return nil, err
					// 	}
					// 	return downloadCertFile(ctx, log, h, ret)
				}
			}
			return handler(ctx, req)
		}
	}
}

func uploadfile(ctx context.Context, log *log.Helper, h transporthttp.ServerInfo) (context.Context, error) {
	log.Info("http: [%s] %s", h.Request.Method, h.Request.URL.Path)
	// tag := h.Request.PostForm.Get("type")
	// // urlPerfix := fmt.Sprintf("http://%s/", c.Request.Host)
	// if tag == "" {
	// 	return handler(ctx, req)
	// } else {
	// 	switch tag {
	// 	case "1": // 单图
	// 		files, ok := h.Request.MultipartForm.File["file"]
	// 		if !ok || len(files) < 0 {
	// 			return nil, ecode.FileErr
	// 		}
	// 		// 上传文件至指定目录
	// 		guid := uuid.New().String()
	// 		singleFile := guid + tools.GetExt(files[0].Filename)

	// 		// _ = h.Request. SaveUploadedFile(files, singleFile)
	// 		// app.OK(c, urlPerfix+singleFile, "上传成功")
	// 		ctx := context.WithValue(ctx, "singleFile", singleFile)
	// 		return handler(ctx, req)
	// 	case "2": // 多图
	// 		// files := c.Request.MultipartForm.File["file"]
	// 		// multipartFile := make([]string, len(files))
	// 		// for _, f := range files {
	// 		// guid := uuid.New().String()
	// 		// multipartFileName := "static/uploadfile/" + guid + utils.GetExt(f.Filename)
	// 		// _ = c.SaveUploadedFile(f, multipartFileName)
	// 		// multipartFile = append(multipartFile, urlPerfix+multipartFileName)
	// 		// }
	// 		// app.OK(c, multipartFile, "上传成功")
	// 		return handler(ctx, req)
	// 	case "3": // base64
	// 		files := h.Request.Form.Get("file")
	// 		ddd, _ := base64.StdEncoding.DecodeString(files)
	// 		guid := uuid.New().String()
	// 		_ = ioutil.WriteFile("static/uploadfile/"+guid+".jpg", ddd, 0666)
	// 		// app.OK(c, urlPerfix+"static/uploadfile/"+guid+".jpg", "上传成功")
	// 	}
	// }
	return ctx, nil
}

func downloadCertFile(ctx context.Context, log *log.Helper, h transporthttp.ServerInfo, ret interface{}) (interface{}, error) {
	reply := ret.(*pb.ApiReply)
	if reply.Code == 0 {
		it := reply.Data[0]
		var d pb.DownloadCertReply
		if err := ptypes.UnmarshalAny(it, &d); err != nil {
			return ctx, err
		}
		file, err := os.Open(d.Path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		fileHeader := make([]byte, 512)
		file.Read(fileHeader)

		fileStat, _ := file.Stat()

		h.Response.Header().Set("Content-Disposition", "attachment;filename="+d.Crt)
		h.Response.Header().Set("Content-Type", "application/force-download")
		h.Response.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size(), 10))

		file.Seek(0, 0)
		io.Copy(h.Response, file)
	}
	return ret, nil
}
