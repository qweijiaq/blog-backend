package image_server

import (
	"backend/global"
	"backend/models"
	"backend/models/ctype"
	"backend/plugins/qiniu"
	"backend/utils"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

var ImageWhiteList = []string{
	"jpg",
	"jpeg",
	"tif",
	"ico",
	"tiff",
	"gif",
	"svg",
	"webp",
	"png",
	"bmp",
}

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

// ImageUploadService 上传多个图片，返回图片的 URL
// @Tags 图片管理
// @Summary 上传多个图片，返回图片的 URL
// @Description 上多单个图片，返回图片的 URL
// @Param token header string true "token"
// @Accept mutlipart/form-data
// @Param images formData file true "上传文件"
// @Router /api/images [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename
	res.FileName = fileName
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	// 文件白名单判断
	if !utils.InList(suffix, ImageWhiteList) {
		res.Msg = "上传的图片文件名后缀不在白名单中"
		return res
	}
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, file.Filename)
	// 判断大小
	size := float64(file.Size) / float64(1024*1024)
	var pre_size float64
	if global.Config.QiNiu.Enable {
		pre_size = global.Config.QiNiu.Size
	} else {
		pre_size = float64(global.Config.Upload.Size)
	}
	if size >= pre_size {
		res.Msg = fmt.Sprintf("图片大小超过设定的大小 %.0f MB，当前图片大小为 %.2f MB", pre_size, size)
		return res
	}
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	ImageHash := utils.Md5(byteData)
	// 去数据库中查询该文件是否已经存在
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", ImageHash).Error
	if err == nil {
		// 找到了
		res.FileName = bannerModel.Name
		res.Msg = "图片已存在"
		return res
	}
	fileType := ctype.Local
	res.Msg = "图片本地上传成功"
	res.IsSuccess = true
	filePath = "/" + filePath
	if global.Config.QiNiu.Enable {
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.Msg = "成功上传到七牛云"
		fileType = ctype.QiNiu
	}
	// 图片入库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      ImageHash,
		Name:      fileName,
		ImageType: fileType,
	})
	return
}
