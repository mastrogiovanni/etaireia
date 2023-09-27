package models

import "mime/multipart"

type SignRequest struct {
	Document  *multipart.FileHeader `form:"document" binding:"required"`
	Signature string                `form:"signature" binding:"required"`
	PublicKey string                `form:"publicKey" binding:"required"`
}
