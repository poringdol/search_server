package db

import (
	"context"
	"fmt"
	"gitlab.stageoffice.ru/UCS-COMMON/utils/errors"
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
	}

	if rec.AddressHouse != "" {
		filter = and(filter, bson.D{{"address_house",
			bson.D{
				{"$eq", rec.AddressHouse}}}})
	}

	if rec.AddressOffice != "" {
		filter = and(filter, bson.D{{"address_office",
			bson.D{
				{"$eq", rec.AddressOffice}}}})
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.New("get cursor")
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			fmt.Println("close cursor error")
		}
	}()

	return readAllRecords(ctx, cursor)
}
