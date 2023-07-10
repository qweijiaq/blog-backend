package service

import (
	"backend/service/image_server"
	"backend/service/user_server"
)

type ServiceGroup struct {
	ImageService image_server.ImageService
	UserService  user_server.UserService
}

var ServiceApp = new(ServiceGroup)
