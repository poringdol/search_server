package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"yandex-food/utils"
)

func GetByAddress(ctx context.Context, client *mongo.Client, rec *utils.Record) ([]*utils.Record, error) {
	collection := getCollection(client)

	filter := bson.D{
		{"address_city",
			bson.D{
				{"$eq", rec.AddressCity},
			},
		},
		{"address_street",
			bson.D{
				{"$eq", rec.AddressStreet},
			},
		},
		{"address_house",
			bson.D{
				{"$eq", rec.AddressHouse},
			},
		},
		{"address_office",
			bson.D{
				{"$eq", rec.AddressOffice},
			},
		},
	}

	cursor, err := collection.Find(ctx, filter)
	utils.CheckError(err, "get cursor")

	return readAllRecords(ctx, cursor)
}
