package services

import (
	"mime/multipart"

	"github.com/Ada-lave/evert-core"
)

type DocService struct{}

func (DC *DocService) FormatDocument(file multipart.FileHeader, params evert.FormatterParams) error {
	openedFile, err := file.Open()
	if err != nil {
		return err
	}
	evertDoc, err := evert.New(openedFile, file.Size)

	if err != nil {
		return err
	}

	evertFormatter := evert.NewFormatter(evertDoc)
	evertFormatter.Format(params)
	return nil
}