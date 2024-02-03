package repositories

import (
	"github.com/mesxx/Fiber_Simple_File_Management_API/models"

	"gorm.io/gorm"
)

type (
	FileRepository interface {
		Create(file *models.File) (*models.File, error)
		GetAll() ([]models.File, error)
		GetByID(id uint) (*models.File, error)
		Delete(file *models.File, id uint) (*models.File, error)
	}

	fileRepository struct {
		DB *gorm.DB
	}
)

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepository{
		DB: db,
	}
}

func (fr fileRepository) Create(file *models.File) (*models.File, error) {
	if err := fr.DB.Create(&file).Error; err != nil {
		return nil, err
	}
	return file, nil
}

func (fr fileRepository) GetAll() ([]models.File, error) {
	var files []models.File
	if err := fr.DB.Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (fr fileRepository) GetByID(id uint) (*models.File, error) {
	var file models.File
	if err := fr.DB.Where("ID = ?", id).Find(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (fr fileRepository) Delete(file *models.File, id uint) (*models.File, error) {
	if err := fr.DB.Where("ID = ?", id).Delete(&file).Error; err != nil {
		return nil, err
	}
	return file, nil
}
