package models

import "mime/multipart"

type UploadRequest struct {
	Document  *multipart.FileHeader `form:"document" binding:"required"`
	Name      string                `form:"name" binding:"required"`
	Surname   string                `form:"surname" binding:"required"`
	PublicKey string                `form:"publicKey" binding:"required"`
}
