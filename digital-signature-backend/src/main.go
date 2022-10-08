package main

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadRequest struct {
	Document  *multipart.FileHeader `form:"document" binding:"required"`
	Name      string                `form:"name" binding:"required"`
	Surname   string                `form:"surname" binding:"required"`
	PublicKey string                `form:"publicKey" binding:"required"`
}

type SignRequest struct {
	Document  *multipart.FileHeader `form:"document" binding:"required"`
	Signature string                `form:"signature" binding:"required"`
	PublicKey string                `form:"publicKey" binding:"required"`
}

type UserData struct {
	Name    string
	Surname string
}

func main() {

	accounts := make(map[string]*UserData)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/v1/sign", func(c *gin.Context) {
		var signRequest SignRequest
		if err := c.ShouldBind(&signRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Saving %s (size: %d)\n", signRequest.Document.Filename, signRequest.Document.Size)

		file, err := signRequest.Document.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("Read fully")

		buffer, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// msg := string(buffer)

		fmt.Printf("Public Key: '%s', Signature: '%s', File Size: %d\n", signRequest.PublicKey, signRequest.Signature, len(buffer))

		user := accounts[signRequest.PublicKey]

		fmt.Printf("Signed By: '%s %s',\n", user.Name, user.Surname)

		/*
			err := c.SaveUploadedFile(uploadRequest.Document, uploadRequest.Document.Filename)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		*/

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	r.POST("/api/v1/upload", func(c *gin.Context) {
		var uploadRequest UploadRequest
		if err := c.ShouldBind(&uploadRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Saving %s (size: %d)\n", uploadRequest.Document.Filename, uploadRequest.Document.Size)

		file, err := uploadRequest.Document.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("Read fully")

		buffer, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Readed %d\n", len(buffer))

		// msg := string(buffer)

		fmt.Printf("Name: '%s', Surname: '%s', Public Key: '%s', File Size: %d\n", uploadRequest.Name, uploadRequest.Surname, uploadRequest.PublicKey, len(buffer))

		accounts[uploadRequest.PublicKey] = &UserData{
			Name:    uploadRequest.Name,
			Surname: uploadRequest.Surname,
		}

		/*
			err := c.SaveUploadedFile(uploadRequest.Document, uploadRequest.Document.Filename)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		*/

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
