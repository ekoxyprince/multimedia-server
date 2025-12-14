package services

import (
	"fmt"
	"image"

	"astrodev.online/multimedia-server/internal/database/models"
	"astrodev.online/multimedia-server/pkg/utils"
)

type UploadService struct {
	upload_repo models.UploadRepository
	image_utils utils.ImageUtility
}
func NewUploadService(upload_repo models.UploadRepository, image_utils utils.ImageUtility) *UploadService{
	return &UploadService{upload_repo: upload_repo,image_utils: image_utils}
}
func (u *UploadService) UploadSingleImage(src image.Image)error{
  imageName,err :=	u.image_utils.ResizeImage(src)
  if err != nil{
	return err
  }
  err = u.upload_repo.CreateUpload(&models.Upload{
	UploadUrl: fmt.Sprintf("http://localhost:8080/%s",imageName),
	UserId: "randomuser_id",
	UploadSize: 1024,
  })
    if err != nil{
	return err
  }
  return nil
}