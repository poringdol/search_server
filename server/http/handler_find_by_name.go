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

func findByName(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)

	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoPath))
	if err != nil {
		fmt.Printf("mongoDB connect: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	url := r.URL
	fullName := &utils.FullName{
		Name:       strings.ToLower(url.Query().Get("name")),
		Surname:    strings.ToLower(url.Query().Get("surname")),
		Patronymic: strings.ToLower(url.Query().Get("patronymic")),
	}

	if fullName.Surname == "" || (fullName.Name == "" && fullName.Patronymic != "") {
		fmt.Println("bad request")
		SetHTTPError(w, http.StatusBadRequest, "")
		return
	}

	people, err := db.GetByName(ctx, mongoClient, fullName)

	response := utils.HTTPGetByNameResponse{}
	if len(people) == 0 {
		response.Error = &utils.HHTPError{Code: http.StatusNotFound, Message: "Записей по данному имени не найдено"}
		responseRecs, _ := json.Marshal(response)
		w.Write(responseRecs)
		return
	}

	uniquePeople := make(map[int64]*utils.Record)
	for _, p := range people {
		uniquePeople[p.PhoneNumber] = p
	}
	for _, p := range uniquePeople {
		response.Persons = append(response.Persons, recordToHTTPPerson(p))
	}
	response.Error = &utils.HHTPError{Code: http.StatusOK}
	responseRecs, _ := json.Marshal(response)
	w.Write(responseRecs)

	w.WriteHeader(http.StatusOK)
}
