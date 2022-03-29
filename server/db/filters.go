package db

import "go.mongodb.org/mongo-driver/bson"

func and(conditions ...interface{}) bson.D {
	return bson.D{{
		"$and",
		bson.A(conditions),
	}}
}

func or(conditions ...interface{}) bson.D {
	return bson.D{{
		"$or",
		bson.A(conditions),
	}}
}
