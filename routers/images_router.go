package routers

import "backend/api"

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("images", imagesApi.ImageUploadView)
	router.POST("image", imagesApi.ImageUploadJustReturnAddrView)
	router.GET("images", imagesApi.ImageListView)
	router.DELETE("images", imagesApi.ImageRemoveView)
	router.PUT("images", imagesApi.ImageUpdateView)
	router.GET("image_names", imagesApi.ImageNameListView)
}
