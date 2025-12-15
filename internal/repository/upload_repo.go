package repository

import (
	"context"

	"astrodev.online/multimedia-server/internal/database/models"
	"gorm.io/gorm"
)

type UploadRepository struct{
	db *gorm.DB
}
func New (db *gorm.DB) models.UploadRepository{
	return &UploadRepository{db:db}
}
func (u *UploadRepository) CreateUpload(data *models.Upload)error{
	ctx := context.Background()
	err := gorm.G[models.Upload](u.db).Create(ctx,data)
	if err != nil {
		return err
	}
	return nil
}
func (u *UploadRepository) CreateMultipleUploads(data *[]models.Upload) error{
	ctx := context.Background()
	err := gorm.G[[]models.Upload](u.db).Create(ctx,data)
	if err != nil {
		return err
	}
	return nil
}
func (u *UploadRepository) GetUploads()(*[]models.Upload,error){
	ctx := context.Background()
	uploads,err := gorm.G[models.Upload](u.db).Find(ctx)
	if err != nil{
		return nil,err
	}
	return &uploads,nil
}

func (u *UploadRepository) GetUploadById(id int64)(*models.Upload,error){
	ctx := context.Background()
	upload,err := gorm.G[models.Upload](u.db).Where("id = ?",id).First(ctx)
	if err != nil {
		return nil,err
	}
	return &upload,nil
}
