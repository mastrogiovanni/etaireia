package models

import "mime/multipart"

type CreateSubscriptionRequest struct {
	Document  *multipart.FileHeader `form:"document" binding:"required"`
	Name      string                `form:"name" binding:"required"`
	Surname   string                `form:"surname" binding:"required"`
	PublicKey string                `form:"publicKey" binding:"required"`
}

type ApproveSubscriptionRequest struct {
	ApproverPublicKey     string `json:"approverPublicKey" binding:"required"`
	SubscriptionPublicKey string `json:"subscriptionPublicKey" binding:"required"`
}
