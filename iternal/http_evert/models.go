package http_evert

import (
	"mime/multipart"

	"github.com/Ada-lave/evert-core"
)

type DocRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
	Params evert.FormatterParams `json:"params" binding:"required"`
}