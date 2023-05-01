package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"myblog_server/global"
	"myblog_server/models/res"
	"os"
	"path"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 反馈信息
}

func (ImagesApi) ImageUploadView(c *gin.Context) {
	// 上传多个文件
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("文件不存在", c)
		return
	}

	basePath := global.Config.Upload.Path
	// 判断路径是否存在
	_, err = os.ReadDir(basePath)
	if err != nil {
		// 不存在则创建文件夹，递归创建文件夹
		err = os.MkdirAll(basePath, fs.ModePerm)
		// 创建文件夹失败则抛出异常
		if err != nil {
			global.Log.Error(err)
		}
	}

	// 存放结果的图像列表
	var resultList []FileUploadResponse

	// 遍历图像列表
	for _, file := range fileList {
		// 获取单个图像文件的完整路径
		filePath := path.Join(basePath, file.Filename)
		// 判断图片大小
		//fmt.Println(filePath, float64(file.Size)/float64(1024*1024))
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			// 向resultList列表中压入图像信息（文件名字， 是否上传成功， 反馈信息）
			resultList = append(resultList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过设定的%.2fMB,当前大小为%.2fMB", float64(global.Config.Upload.Size), size),
			})
			continue
		}

		// 保存上传的图像文件
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			// 保存失败也属于图像上传的不成功
			global.Log.Error(err)
			resultList = append(resultList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
		}
		continue

		resultList = append(resultList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})

		//fmt.Println(file.Header, file.Size, file.Filename)
	}

	res.OkWithData(resultList, c)
}
