package http_evert

import (
	"mime/multipart"
)

type FileParams struct {
	AddSpacesBeetweenImageText bool `form:"add_spaces_beetween_image_text" binding:"required"`
	FormatImageDescription bool `form:"format_image_description,omitempty"`
}

type DocRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}