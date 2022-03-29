package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"yandex-food/utils"
)

// CreateIndex creating indexes
func CreateIndex(ctx context.Context, client *mongo.Client) {
	fmt.Println("creating indexes")

	collection := getCollection(client)

	_, err := collection.Indexes().CreateMany(
		ctx,
		[]mongo.IndexModel{
			{
				Keys: bson.D{
					{
						Key:   "email",
						Value: 1,
					},
				},
			},
			{
				Keys: bson.D{
					{
						Key:   "full_name",
						Value: 1,
					},
				},
			},
			{
				Keys: bson.D{
					{
						Key:   "phone_number",
						Value: 1,
					},
				},
			},
			{
				Keys: bson.D{
					{
						Key:   "address_city",
						Value: 1,
					},
				},
			},
			{
				Keys: bson.D{
					{
						Key:   "address_street",
						Value: 1,
					},
				},
			},
		},
	)
	utils.CheckError(err, "create index")
}
