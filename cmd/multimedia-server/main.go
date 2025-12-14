package main

import (
	"log"
	"os"

	"astrodev.online/multimedia-server/internal/database"
	"astrodev.online/multimedia-server/internal/handlers"
	"astrodev.online/multimedia-server/internal/repository"
	"astrodev.online/multimedia-server/internal/routes"
	"astrodev.online/multimedia-server/internal/services"
	"astrodev.online/multimedia-server/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
  err := godotenv.Load()	
  if err != nil{
	log.Fatal(err)
  }
  url := os.Getenv("DB_URL")
db,err := database.New(url)
   if err != nil{
	log.Fatal(err)
  }
 uploadRepo := repository.New(db)
 uploadService := services.NewUploadService(uploadRepo,utils.ImageUtilityHelper{})
 uploadHandler := handlers.NewUploadHandler(uploadService)
 
 server := gin.Default()
 routes.RegisterUploadRoutes(server,uploadHandler)
  server.Static("/uploads","./uploads")
  server.Run(":8080")
}