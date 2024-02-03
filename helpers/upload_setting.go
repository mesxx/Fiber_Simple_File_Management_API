package helpers

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

func UploadSettingType(fileType string) error {
	allowedTypes := []string{"image/jpeg", "image/jpg", "image/png"}

	for i := 0; i < len(allowedTypes); i++ {
		el := allowedTypes[i]
		if el == fileType {
			return nil
		}
	}

	return errors.New("unsupported file type")
}

func UploadSettingName(originalName string) (string, error) {
	removeExt := strings.Split(originalName, ".")[0]
	fileExt := strings.Split(originalName, ".")[1]
	lowerCase := strings.ToLower(removeExt)
	removeSpacing := strings.ReplaceAll(lowerCase, " ", "")
	fileName := uuid.New().String() + "-" + removeSpacing + "." + fileExt

	return fileName, nil
}
