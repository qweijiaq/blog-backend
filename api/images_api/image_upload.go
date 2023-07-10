package images_api

import (
	"backend/global"
	"backend/models/res"
	"backend/service"
	"backend/service/image_server"
	"io/fs"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

//type FileUploadResponse struct {
//	FileName  string `json:"file_name"`  // 文件名
//	IsSuccess bool   `json:"is_success"` // 是否上传成功
//	Msg       string `json:"msg"`        // 消息
//}

// ImageUploadView 上传多个图片，返回图片的 URL
// @Tags 图片管理
// @Summary 上传多个图片，返回图片的 URL
// @Description 上传多个图片，返回图片的 URL
// @Param token header string true "token"
// @Accept mutlipart/form-data
// @Param images formData file true "上传文件"
// @Router /api/images [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, _ := form.File["images"]
	if len(fileList) == 0 {
		res.FailWithMessage("图片不存在", c)
		return
	}
	// 判断路径是否存在 -- 不存在就创建
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		// 递归创建
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	var resList []image_server.FileUploadResponse

	for _, file := range fileList {
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		if !global.Config.QiNiu.Enable {
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.FileName = file.Filename
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
			filePath := path.Join(basePath, file.Filename)
			err = c.SaveUploadedFile(file, filePath)
			if err != nil {
				global.Log.Error(err)
				serviceRes.IsSuccess = false
				serviceRes.FileName = file.Filename
				serviceRes.Msg = err.Error()
				return
			}
		}
		resList = append(resList, serviceRes)
	}
	res.OkWithData(resList, c)
}
