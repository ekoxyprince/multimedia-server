package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"time"

	"github.com/disintegration/imaging"
)


type ImageUtility interface {
	ResizeImage(srcImage image.Image)(string,error)
}


type ImageUtilityHelper struct{

}

func (i ImageUtilityHelper) ResizeImage(srcImage image.Image)(string,error){
 dstImage :=  imaging.Resize(srcImage,250,250,imaging.Lanczos)
 imageName := fmt.Sprintf("%d-image.jpg",time.Now().UnixNano())
 writer,err := os.Create("uploads/"+imageName)
 if err != nil{
	return "",err
 }
 defer writer.Close()
 err = jpeg.Encode(writer,dstImage,&jpeg.Options{Quality: 100})
  if err != nil{
	return "",err
 }
 return imageName,nil
}