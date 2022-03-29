package db

import (
	"context"
	"gitlab.stageoffice.ru/UCS-COMMON/gaben"
	"go.mongodb.org/mongo-driver/mongo"
	"yandex-food/utils"
)

func readAllRecords(ctx context.Context, cursor *mongo.Cursor) ([]*utils.Record, error) {
	var records []*utils.DBRecord

	defer func() {
		if err := cursor.Close(ctx); err != nil {
			gaben.Ctx(ctx).Error("close read user records cursor", gaben.Error(err))
		}
	}()

	for cursor.Next(ctx) {
		if err := cursor.Err(); err != nil {
			return nil, err
		}

		var record *utils.DBRecord
		if err := cursor.Decode(&record); err != nil {
			return nil, err
		}

		records = append(records, record)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return fromBsonData(records), nil
}
