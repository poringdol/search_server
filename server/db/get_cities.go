package db

import (
	"context"
	"gitlab.stageoffice.ru/UCS-COMMON/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCities(ctx context.Context, client *mongo.Client) ([]string, error) {
	collection := getCollection(client)

	filter := bson.D{}
	dbCities, err := collection.Distinct(ctx, "address_city", filter)
	if err != nil {
		return nil, errors.New("get cursor")
	}

	cities := make([]string, 0, len(dbCities))
	for _, c := range dbCities {
		cities = append(cities, c.(string))
	}

	return cities, nil
}
