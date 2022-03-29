package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"yandex-food/utils"
)

func GetByPhone(ctx context.Context, client *mongo.Client, phone int64) ([]*utils.Record, error) {
	collection := getCollection(client)

	filter := bson.D{{"phone_number", bson.D{{"$eq", phone}}}}

	cursor, err := collection.Find(ctx, filter)
	utils.CheckError(err, "get cursor")

	return readAllRecords(ctx, cursor)
}
