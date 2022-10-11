package controllers

import (
	"errors"
	"fmt"
	"io/ioutil"

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

	fmt.Printf("Saving %s (size: %d)\n", signRequest.Document.Filename, signRequest.Document.Size)

	file, err := signRequest.Document.Open()
	if err != nil {
		models.ResponseError(c, err)
		return
	}

	fmt.Println("Read fully")

	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		models.ResponseError(c, err)
		return
	}

	fmt.Printf("Public Key: '%s', Signature: '%s', File Size: %d\n", signRequest.PublicKey, signRequest.Signature, len(buffer))

	user, err := database.FindByPublicKey(signRequest.PublicKey)
	if err != nil {
		models.ResponseError(c, err)
		return
	}
	if user == nil {
		models.ResponseError(c, errors.New("bad publicKey for this document"))
		return
	}

	fmt.Printf("Signed By: '%s %s',\n", user.Name, user.Surname)

	err = database.InsertSignedDocument(&database.Signed{
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
