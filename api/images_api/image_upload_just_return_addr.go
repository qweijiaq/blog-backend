package images_api

import (
	"backend/global"
	"backend/models/res"
	"backend/service/image_server"
	"backend/utils"
	"fmt"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

// ImageUploadJustReturnAddrView 上传单个图片，返回图片的 URL
// @Tags 图片管理
// @Summary 上传单个图片，返回图片的 URL
// @Description 上传单个图片，返回图片的 URL
// @Param token header string true "token"
// @Accept mutlipart/form-data
// @Param images formData file true "上传文件"
// @Router /api/images [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (ImagesApi) ImageUploadJustReturnAddrView(c *gin.Context) {
	file, err := c.FormFile("images")
	if err != nil {
		res.FailWithMessage("参数校验失败", c)
		return
	}

	fileName := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, fileName)

	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	// 文件白名单判断
	if !utils.InList(suffix, image_server.ImageWhiteList) {
		res.FailWithMessage("上传的图片文件后缀不在白名单中", c)
		return
	}

	// 判断大小
	size := float64(file.Size) / float64(1024*1024)
	var pre_size float64
	if global.Config.QiNiu.Enable {
		pre_size = global.Config.QiNiu.Size
	} else {
		pre_size = float64(global.Config.Upload.Size)
	}
	if size >= pre_size {
		res.FailWithMessage(fmt.Sprintf("图片大小超过设定的大小 %.0f MB，当前图片大小为 %.2f MB", pre_size, size), c)
		return
	}

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData("/"+filePath, c)
	return
}
