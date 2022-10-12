package controllers

import (
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mastrogiovanni/digital-signature-backend/internal/database"
	"github.com/mastrogiovanni/digital-signature-backend/internal/server/models"
)

func CreateSubscription(c *gin.Context) {

	var createSubscriptionRequest models.CreateSubscriptionRequest
	if err := c.ShouldBind(&createSubscriptionRequest); err != nil {
		models.ResponseError(c, err)
		return
	}

	log.Printf("Saving %s (size: %d)\n", createSubscriptionRequest.Document.Filename, createSubscriptionRequest.Document.Size)

	file, err := createSubscriptionRequest.Document.Open()
	if err != nil {
		models.ResponseError(c, err)
		return
	}

	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		models.ResponseError(c, err)
		return
	}

	log.Printf("Name: '%s', Surname: '%s', Public Key: '%s', File Size: %d\n", createSubscriptionRequest.Name, createSubscriptionRequest.Surname, createSubscriptionRequest.PublicKey, len(buffer))

	err = database.CreateSubscription(&database.Subscription{
		PublicKey:        createSubscriptionRequest.PublicKey,
		IdentityDocument: buffer,
		Name:             createSubscriptionRequest.Name,
		Surname:          createSubscriptionRequest.Surname,
	})
	if err != nil {
		models.ResponseError(c, err)
		return
	}

	models.ResponseSuccess(c)

}

func ListSubscription(c *gin.Context) {

	subscriptions, err := database.FindSubscriptions()
	if err != nil {
		models.ResponseError(c, err)
		return
	}

	models.ResponseSuccessWithData(c, subscriptions)

}

func SubscriptionApprove(c *gin.Context) {

	var approveSubscriptionRequest models.ApproveSubscriptionRequest
	if err := c.ShouldBind(&approveSubscriptionRequest); err != nil {
		models.ResponseError(c, err)
		return
	}

	log.Printf("%s is approving %s\n", approveSubscriptionRequest.ApproverPublicKey, approveSubscriptionRequest.SubscriptionPublicKey)

	subscription, err := database.ApproveSubscription(&database.ApproveSubscriptionDto{
		ApproverPublicKey:     approveSubscriptionRequest.ApproverPublicKey,
		SubscriptionPublicKey: approveSubscriptionRequest.SubscriptionPublicKey,
	})
	if err != nil {
		models.ResponseError(c, err)
		return
	}

	models.ResponseSuccessWithData(c, subscription)

}
