package handlers

import (
	"image"
	"log"
	"net/http"

	"astrodev.online/multimedia-server/internal/services"
	"github.com/gin-gonic/gin"
)

type UploadHandler struct{
	uploadService *services.UploadService
}

func NewUploadHandler(upload_service *services.UploadService)*UploadHandler{
	return &UploadHandler{uploadService: upload_service}
}

func (handler *UploadHandler) UploadSingleImage(c *gin.Context){
 fileHeader,err := c.FormFile("image")
 if err != nil{
  log.Print(err)
  c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":"An error occured while parsing form"})
  return 
 }

 file, err := fileHeader.Open()
  if err != nil{
  log.Print(err)
  c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":"An error occured while opening file header"})
  return 
 }
 defer file.Close()
 srcImage,format,err := image.Decode(file)
  log.Print(format)
  if err != nil{
  log.Print(err)
  c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":"An error occured while decoding file"})
  return 
 }
 err = handler.uploadService.UploadSingleImage(srcImage)
 if err != nil{
  log.Print(err)
  c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":"An error occured while uploading file"})
  return 
 }
 c.JSON(http.StatusOK,gin.H{"success":true,"message":"file uploaded"})
}