package http

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strings"
	"yandex-food/server/db"
	"yandex-food/utils"
)

func findByAddress(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)

	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoPath))
	if err != nil {
		fmt.Printf("mongoDB connect: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	url := r.URL
	address := &utils.Record{
		AddressCity:   strings.ToLower(url.Query().Get("city")),
		AddressStreet: strings.ToLower(url.Query().Get("street")),
		AddressHouse:  strings.ToLower(url.Query().Get("house")),
		AddressOffice: strings.ToLower(url.Query().Get("office")),
	}

	if address.AddressCity == "" || address.AddressStreet == "" {
		fmt.Print("bad request, empty city or street")
		SetHTTPError(w, http.StatusBadRequest, "empty city or street")
		return
	}

	response := utils.HTTPGetByAddressResponse{}
	byAddress, err := db.GetByAddress(ctx, mongoClient, address)
	for _, recByAddr := range byAddress {
		response.Addresses = append(response.Addresses, recordToHTTPFullRecord(recByAddr))
	}

	response.Addresses = append(response.Addresses)
	response.Error = &utils.HHTPError{Code: http.StatusOK}
	responseRecs, _ := json.Marshal(response)
	w.Write(responseRecs)

	w.WriteHeader(http.StatusOK)
}
