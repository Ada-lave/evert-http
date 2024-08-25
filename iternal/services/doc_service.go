package services

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/Ada-lave/evert-core"
)

type DocService struct{}

func (DC *DocService) FormatDocument(file multipart.FileHeader, params evert.FormatterParams) (string, error) {
	openedFile, err := file.Open()
	if err != nil {
		return "", err
	}
	evertDoc, err := evert.New(openedFile, file.Size)

	if err != nil {
		return "", err
	}

	evertFormatter := evert.NewFormatter(evertDoc)
	evertFormatter.Format(params)

	tempDir := os.TempDir()
	filePath := filepath.Join(tempDir, file.Filename)


	err = evertDoc.SaveFormattedDoc(filePath)

	if err != nil {
		return "", err
	}
	return filePath, nil
}