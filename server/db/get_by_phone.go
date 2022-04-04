package db

import (
	"context"
	"fmt"
	"gitlab.stageoffice.ru/UCS-COMMON/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"yandex-food/utils"
)

func GetByPhone(ctx context.Context, client *mongo.Client, phone int64) ([]*utils.Record, error) {
	collection := getCollection(client)

	filter := bson.D{{"phone_number", bson.D{{"$eq", phone}}}}

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
