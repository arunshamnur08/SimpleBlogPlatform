package mongo_client

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Close(client *mongo.Client) {

	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func Connect(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://user:pass@localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}

func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	// select database and collection ith Client.Database method
	// and Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertOne accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertOne(context.Background(), doc)

	return result, err
}

func Insert(ctx context.Context, client *mongo.Client, document interface{}) error {
	insertOneResult, err := insertOne(client, ctx, "sample_training",
		"posts", document)
	if err != nil {
		fmt.Println("failed to insert document", insertOneResult)
		return err
	}
	return nil

}

func QueryAll(client *mongo.Client, dataBase, col string, filter bson.D) *mongo.Cursor {
	collection := client.Database(dataBase).Collection(col)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	return cursor
}

func Delete(client *mongo.Client, dataBase, col string, filter bson.D) (*mongo.DeleteResult, error) {
	opts := options.Delete().SetHint(bson.D{{"_id", 1}})
	collection := client.Database(dataBase).Collection(col)

	result, err := collection.DeleteMany(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Update(client *mongo.Client, dataBase, col string, filter bson.D, update bson.D) (*mongo.UpdateResult, error) {
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}
