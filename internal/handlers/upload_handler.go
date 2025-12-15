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
type ImageForm struct{
 Width string `form:"width" binding:"required"`
 Height string `form:"height" binding:"required"`
 Mode string `form:"mode" binding:"required"`
}

func NewUploadHandler(upload_service *services.UploadService)*UploadHandler{
	return &UploadHandler{uploadService: upload_service}
}

func (handler *UploadHandler) UploadSingleImage(c *gin.Context){
 var form ImageForm;	
 err := c.ShouldBind(&form)
  if err != nil{
  log.Print(err)
  c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":err.Error()})
  return 
 }
 fileHeader,err := c.FormFile("image")
 if err != nil{
  log.Print(err)
  c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":err.Error()})
  return 
 }

 file, err := fileHeader.Open()
  if err != nil{
  log.Print(err)
  c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":err.Error()})
  return 
 }
 defer file.Close()
 srcImage,format,err := image.Decode(file)
  log.Print(format)
  if err != nil{
  log.Print(err)
  c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":err.Error()})
  return 
 }
 err = handler.uploadService.UploadSingleImage(srcImage,form.Width,form.Height,form.Mode,fileHeader.Filename)
 if err != nil{
  log.Print(err)
  c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":"An error occured while uploading file"})
  return 
 }
 c.JSON(http.StatusOK,gin.H{"success":true,"message":"file uploaded"})
}
func (handler *UploadHandler) UploadMultipleImage(c *gin.Context){
  var formItem ImageForm
  err := c.ShouldBind(&formItem)
  if err != nil{
  log.Print(err)
  c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":err.Error()})
  return 
 }
  form,err := c.MultipartForm()
  if err != nil{
  log.Print(err)
  c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":err.Error()})
  return 
 }
 fileHeaders := form.File["image"]
 err = handler.uploadService.UploadMultipleImage(fileHeaders,formItem.Width,formItem.Height,formItem.Mode)
    if err != nil{
  log.Print(err)
  c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":err.Error()})
  return 
 }
 c.JSON(http.StatusCreated,gin.H{"success":true,"message":"Images uploaded"})
}