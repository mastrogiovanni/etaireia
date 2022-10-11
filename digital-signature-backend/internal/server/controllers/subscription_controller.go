package controllers

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/mastrogiovanni/digital-signature-backend/internal/database"
	"github.com/mastrogiovanni/digital-signature-backend/internal/server/models"
)

func UploadController(c *gin.Context) {

	var uploadRequest models.UploadRequest
	if err := c.ShouldBind(&uploadRequest); err != nil {
		models.ResponseError(c, err)
		return
	}

	fmt.Printf("Saving %s (size: %d)\n", uploadRequest.Document.Filename, uploadRequest.Document.Size)

	file, err := uploadRequest.Document.Open()
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

	fmt.Printf("Readed %d\n", len(buffer))

	// msg := string(buffer)

	fmt.Printf("Name: '%s', Surname: '%s', Public Key: '%s', File Size: %d\n", uploadRequest.Name, uploadRequest.Surname, uploadRequest.PublicKey, len(buffer))

	/*
		accounts[uploadRequest.PublicKey] = &UserData{
			Name:    uploadRequest.Name,
			Surname: uploadRequest.Surname,
		}
	*/

	err = database.InsertIdentity(&database.Identity{
		PublicKey:        uploadRequest.PublicKey,
		IdentityDocument: buffer,
		Name:             uploadRequest.Name,
		Surname:          uploadRequest.Surname,
	})
	if err != nil {
		models.ResponseError(c, err)
		return
	}

	models.ResponseSuccess(c)

}
