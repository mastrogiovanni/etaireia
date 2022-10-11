package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Identity struct {
	PublicKey        string `bson:"publicKey,omitempty"`
	IdentityDocument []byte `bson:"identityDocument,omitempty"`
	Name             string `bson:"name,omitempty"`
	Surname          string `bson:"surname,omitempty"`
}

type FindByPublicKeyQuery struct {
	PublicKey string `bson:"publicKey,omitempty"`
}

func InsertIdentity(identity *Identity) error {

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

	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection("identity")

	_, err = collection.InsertOne(ctx, identity)
	return err

}

func FindByPublicKey(publicKey string) (*Identity, error) {

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
		return nil, err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Println(err)
		}
	}()

	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection("identity")

	result := collection.FindOne(ctx, FindByPublicKeyQuery{PublicKey: publicKey})

	var identity Identity
	err = result.Decode(&identity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &identity, nil

}
