package controllers

import (
	"encoding/hex"
	"errors"
	"io/ioutil"
	"log"

	"crypto/ed25519"

	"github.com/gin-gonic/gin"
	"github.com/mastrogiovanni/digital-signature-backend/internal/database"
	"github.com/mastrogiovanni/digital-signature-backend/internal/server/models"
)

func SignController(c *gin.Context) {

	var signRequest models.SignRequest
	if err := c.ShouldBind(&signRequest); err != nil {
		models.ResponseError(c, err)
		return
	}

	log.Printf("Saving %s (size: %d)\n", signRequest.Document.Filename, signRequest.Document.Size)

	file, err := signRequest.Document.Open()
	if err != nil {
		models.ResponseError(c, err)
		return
	}

	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		models.ResponseError(c, err)
		return
	}

	log.Printf("Public Key: '%s', Signature: '%s', File Size: %d\n", signRequest.PublicKey, signRequest.Signature, len(buffer))

	user, err := database.FindSubscriptionByPublicKey(signRequest.PublicKey)
	if err != nil {
		models.ResponseError(c, err)
		return
	}
	if user == nil {
		models.ResponseError(c, errors.New("bad publicKey for this document"))
		return
	}

	signatureBytes, err := hex.DecodeString(signRequest.Signature)
	if err != nil {
		models.ResponseError(c, err)
		return
	}

	publicKeyBytes, err := hex.DecodeString(signRequest.PublicKey)
	if err != nil {
		models.ResponseError(c, err)
		return
	}

	verified := ed25519.Verify(publicKeyBytes, buffer, signatureBytes)
	if !verified {
		models.ResponseError(c, errors.New("signature is not valid"))
		return
	}

	log.Printf("Signed By: '%s %s',\n", user.Name, user.Surname)

	err = database.CreateSignedDocument(&database.Signed{
		Document:  buffer,
		PublicKey: signRequest.PublicKey,
		Signature: signRequest.Signature,
	})
	if err != nil {
		models.ResponseError(c, err)
		return
	}

	models.ResponseSuccess(c)

}

func ListSigned(c *gin.Context) {

	publicKey := c.Param("publicKey")

	log.Println("Search for", publicKey)

	signeds, err := database.FindSignedByPublicKey(publicKey)
	if err != nil {
		models.ResponseError(c, err)
		return
	}

	models.ResponseSuccessWithData(c, signeds)

}
