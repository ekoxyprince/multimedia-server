package models

import (
	"gorm.io/gorm"
)

type Upload struct{
	gorm.Model
	UserId string `json:"userId"`
	UploadUrl string `json:"uploadUrl"`
	UploadSize int64 `json:"uploadSize"`
}

type UploadRepository interface{
	CreateUpload(upload *Upload)error
	CreateMultipleUploads(uploads *[]Upload)error
	GetUploads()(*[]Upload,error)
	GetUploadById(id int64)(*Upload,error)
}