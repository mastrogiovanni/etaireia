package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Signed struct {
	Document  []byte    `bson:"document,omitempty"`
	PublicKey string    `bson:"publicKey,omitempty"`
	Signature string    `bson:"signature,omitempty"`
	Timestamp time.Time `bson:"timestamp,omitempty"`
}

func CreateSignedDocument(signedDocument *Signed) error {

	signedDocument.Timestamp = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credentials := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		Username:      os.Getenv("MONGO_USERNAME"),
		Password:      os.Getenv("MONGO_PASSWORD"),
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL")).SetAuth(credentials)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Println(err)
		}
	}()

	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection("signed")

	_, err = collection.InsertOne(ctx, signedDocument)
	return err

}
