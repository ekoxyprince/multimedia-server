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
	ResizeImage(srcImage image.Image,width,height int64,mode,filename string)(string,error)
}

type ImageUtilityHelper struct{

}

func (i ImageUtilityHelper) ResizeImage(srcImage image.Image,width,height int64,mode,filename string)(string,error){
 var dstImage *image.NRGBA
 switch mode {
 case "resize":
	 dstImage =  imaging.Resize(srcImage,int(width),int(height),imaging.Lanczos)
 case "scale":
	 dstImage = imaging.Fit(srcImage,int(width),int(height),imaging.Lanczos)
 case "crop":
	  dstImage = imaging.Fill(srcImage,int(width),int(height),imaging.Center,imaging.Lanczos)
 }
 imageName := fmt.Sprintf("%d--%s",time.Now().UnixNano(),filename)
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