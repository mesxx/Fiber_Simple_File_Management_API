package usecases

import (
	"errors"
	"os"

	"github.com/mesxx/Fiber_Simple_File_Management_API/models"
	"github.com/mesxx/Fiber_Simple_File_Management_API/repositories"
)

type (
	FileUsecase interface {
		Create(file *models.File) (*models.File, error)
		GetAll() ([]models.File, error)
		GetByID(id uint) (*models.File, error)
		Delete(id uint) (*models.File, error)
	}

	fileUsecase struct {
		FileRepository repositories.FileRepository
	}
)

func NewFileUsecase(fr repositories.FileRepository) FileUsecase {
	return &fileUsecase{
		FileRepository: fr,
	}
}

func (fu fileUsecase) Create(file *models.File) (*models.File, error) {
	return fu.FileRepository.Create(file)
}

func (fu fileUsecase) GetAll() ([]models.File, error) {
	return fu.FileRepository.GetAll()
}

func (fu fileUsecase) GetByID(id uint) (*models.File, error) {
	file, err := fu.FileRepository.GetByID(id)
	if err != nil {
		return nil, err
	} else if file.ID == 0 {
		return nil, errors.New("file ID is invalid, please try again")
	}

	return file, nil
}

func (fu fileUsecase) Delete(id uint) (*models.File, error) {
	getFile, err := fu.FileRepository.GetByID(id)
	if err != nil {
		return nil, err
	} else if getFile.ID == 0 {
		return nil, errors.New("file ID is invalid, please try again")
	}

	//
	filename := getFile.Name
	destination := "./public/images/" + filename
	if err := os.Remove(destination); err != nil {
		return nil, err
	}
	//

	fileDeleted, err := fu.FileRepository.Delete(getFile, getFile.ID)
	if err != nil {
		return nil, err
	}

	return fileDeleted, nil
}
