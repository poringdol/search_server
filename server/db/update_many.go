package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"yandex-food/utils"
)

func updateMany(ctx context.Context, client *mongo.Client, records []*utils.Record) {
	collection := getCollection(client)

	dbRecords := toBsonData(records)
	for _, dbRec := range dbRecords {
		_, err := collection.UpdateOne(
			ctx,
			bson.M{
				"_id": bson.M{"$eq": dbRec.ID},
			},
			bson.M{
				"$set": dbRec,
			},
			options.Update().SetUpsert(true),
		)
		if err != nil {
			fmt.Printf("upsert failed: %s", err)
		}
	}
}
