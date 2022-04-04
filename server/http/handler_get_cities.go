package http

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"yandex-food/server/db"
	"yandex-food/utils"
)

func getCities(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)

	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoPath))
	if err != nil {
		fmt.Printf("mongoDB connect: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := utils.HTTPGetCitiesResponse{}
	cities, err := db.GetCities(ctx, mongoClient)
	if err != nil {
		fmt.Println("get cities")
		SetHTTPError(w, http.StatusInternalServerError, "")
		return
	}

	response.Cities = cities
	response.Error = &utils.HHTPError{Code: http.StatusOK}
	responseRecs, _ := json.Marshal(response)
	w.Write(responseRecs)

	w.WriteHeader(http.StatusOK)
}
