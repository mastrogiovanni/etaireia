package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SubscriptionStatus int8

const (
	Requested SubscriptionStatus = iota
	Rejected
	Approved
)

type Subscription struct {
	PublicKey        string             `bson:"publicKey,omitempty"`
	IdentityDocument []byte             `bson:"identityDocument,omitempty"`
	Name             string             `bson:"name,omitempty"`
	Surname          string             `bson:"surname,omitempty"`
	Status           SubscriptionStatus `bson:"status"`
	Approver         *string            `bson:"approver"`
}

type ApproveSubscriptionDto struct {
	ApproverPublicKey     string `bson:"approverPublicKey,omitempty"`
	SubscriptionPublicKey string `bson:"subscriptionPublicKey,omitempty"`
}

type FindByPublicKeyQuery struct {
	PublicKey string `bson:"publicKey,omitempty"`
}

func CreateSubscription(subscription *Subscription) error {

	log.Println(os.Getenv("MONGO_USERNAME"))
	log.Println(os.Getenv("MONGO_PASSWORD"))
	log.Println(os.Getenv("MONGO_URL"))
	log.Println(os.Getenv("DATABASE_NAME"))

	subscription.Status = Requested
	subscription.Approver = nil

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

	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection("subscription")

	_, err = collection.InsertOne(ctx, subscription)
	return err

}

func ApproveSubscription(approveSubscriptionRequest *ApproveSubscriptionDto) (*Subscription, error) {

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

	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection("subscription")

	filter := bson.M{
		"status":    Requested,
		"publicKey": approveSubscriptionRequest.SubscriptionPublicKey,
	}

	update := bson.M{
		"$set": bson.M{
			"status":   Approved,
			"approver": approveSubscriptionRequest.ApproverPublicKey,
		},
	}

	upsert := false
	after := options.After

	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := collection.FindOneAndUpdate(ctx, filter, update, &opt)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var doc Subscription
	decodeErr := result.Decode(&doc)

	return &doc, decodeErr
}

func FindSubscriptions() ([]Subscription, error) {

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

	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection("subscription")

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []Subscription
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil

}

func FindSubscriptionByPublicKey(publicKey string) (*Subscription, error) {

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

	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection("subscription")

	result := collection.FindOne(ctx, FindByPublicKeyQuery{PublicKey: publicKey})

	var subscription Subscription
	err = result.Decode(&subscription)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &subscription, nil

}
