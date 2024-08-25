package services

import (
	"mime/multipart"

	"github.com/Ada-lave/evert-core"
)

type DocService struct{}

func (DC *DocService) FormatDocument(file multipart.FileHeader, params evert.FormatterParams) ([]byte, error) {
	openedFile, err := file.Open()
	if err != nil {
		return []byte{}, err
	}
	evertDoc, err := evert.New(openedFile, file.Size)

	if err != nil {
		return []byte{}, err
	}

	evertFormatter := evert.NewFormatter(evertDoc)
	evertFormatter.Format(params)

	res, err := evertDoc.GetBytes()

	if err != nil {
		return []byte{}, err
	}
	return res, nil
}