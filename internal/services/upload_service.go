package services

import (
	"fmt"
	"image"
	"mime/multipart"
	"strconv"

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
func (u *UploadService) UploadSingleImage(src image.Image,width,height,mode,filename string)error{
  intWidth,err :=strconv.ParseInt(width,10,64)
  if err != nil{
	return err
  }
  intHeight,err := strconv.ParseInt(height,10,64)
  if err != nil{
	return err
  }
  imageName,err :=	u.image_utils.ResizeImage(src,intWidth,intHeight,mode,filename)
  if err != nil{
	return err
  }
  err = u.upload_repo.CreateUpload(&models.Upload{
	UploadUrl: fmt.Sprintf("http://localhost:8080/uploads/%s",imageName),
	UserId: "randomuser_id",
	UploadSize: 1024,
  })
    if err != nil{
	return err
  }
  return nil
}

func (u *UploadService) UploadMultipleImage(fileHeaders []*multipart.FileHeader,width,height,mode string)error{
 uploads := make([]models.Upload,0)
  intWidth,err :=strconv.ParseInt(width,10,64)
  if err != nil{
	return err
  }
  intHeight,err := strconv.ParseInt(height,10,64)
  if err != nil{
	return err
  }
  for _,fileHeader := range fileHeaders{
   file,err := fileHeader.Open()
   if err != nil {
	return err
   }
   defer file.Close()
   srcImage,_,err := image.Decode(file)
   if err != nil{
	return err
   }
   imageName,err := u.image_utils.ResizeImage(srcImage,intWidth,intHeight,mode,fileHeader.Filename)
   if err != nil{
	return err
   }
    upload := models.Upload{
	UploadUrl: fmt.Sprintf("http://localhost:8080/uploads/%s",imageName),
	UserId: "randomuser_id",
	UploadSize: 1024,
  }
   uploads = append(uploads,upload)
  }
  err = u.upload_repo.CreateMultipleUploads(&uploads)
  if err != nil{
	return err
  }
  return nil
}