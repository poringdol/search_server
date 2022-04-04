package db

import (
	"context"
	"fmt"
	"gitlab.stageoffice.ru/UCS-COMMON/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"yandex-food/utils"
)

func GetByName(ctx context.Context, client *mongo.Client, fullName *utils.FullName) ([]*utils.Record, error) {
	collection := getCollection(client)

	var cursor *mongo.Cursor
	var err error

	if fullName.Name == "" {
		// есть только фамилия
		cursor, err = collection.Find(ctx,
			bson.D{{"full_name", bson.D{{"$regex", fullName.Surname}}}})
	} else if fullName.Patronymic == "" {
		// есть имя и фамилия
		filter := and(bson.D{{"full_name", bson.D{{"$regex", fullName.Name}}}},
			bson.D{{"full_name", bson.D{{"$regex", fullName.Surname}}}})
		cursor, err = collection.Find(ctx, filter)
	} else {
		// есть, имя фамилия и отчество
		filter := or(bson.D{{"full_name", bson.D{{"$eq", strings.Join([]string{fullName.Name, fullName.Patronymic, fullName.Surname}, " ")}}}},
			bson.D{{"full_name", bson.D{{"$eq", strings.Join([]string{fullName.Surname, fullName.Name, fullName.Patronymic}, " ")}}}})
		cursor, err = collection.Find(ctx, filter)
	}
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
