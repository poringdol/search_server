package http

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"yandex-food/server/db"
	"yandex-food/utils"
)

func findByPhone(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)

	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoPath))
	if err != nil {
		fmt.Printf("mongoDB connect: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	url := r.URL
	phone, err := strconv.Atoi(url.Query().Get("phone"))
	if err != nil {
		SetHTTPError(w, http.StatusInternalServerError, "Invalid phone number format. Should be 79.........\n")
		return
	}

	byPhone, err := db.GetByPhone(ctx, mongoClient, int64(phone))
	if err != nil {
		fmt.Printf("get records by phone: %s", err)
		SetHTTPError(w, http.StatusInternalServerError, "")
		return
	}

	response := utils.HTTPGetByPhoneResponse{}
	if len(byPhone) == 0 {
		response.Error = &utils.HHTPError{Code: http.StatusNotFound, Message: "Записей по данному телефону не найдено"}
		responseRecs, _ := json.Marshal(response)
		w.Write(responseRecs)
		return
	}

	response.Person = recordToHTTPPerson(byPhone[0])

	for i, rec := range byPhone {
		byAddress, err := db.GetByAddress(ctx, mongoClient, rec)
		if err != nil {
			fmt.Printf("get record by address: %s", err)
			SetHTTPError(w, http.StatusInternalServerError, "")
			return
		}
		response.Addresses = append(response.Addresses, recordToHTTPAddress(rec))
		for _, recByAddr := range byAddress {
			response.Addresses[i].FullRecords = append(response.Addresses[i].FullRecords, recordToHTTPFullRecord(recByAddr))
		}
	}
	response.Error = &utils.HHTPError{Code: http.StatusOK}
	responseRecs, _ := json.Marshal(response)
	w.Write(responseRecs)

	w.WriteHeader(http.StatusOK)
}
