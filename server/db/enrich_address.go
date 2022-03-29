package db

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"yandex-food/utils"
)

type Address struct {
	PostalCode     int64  `json:"postal_code"`
	Country        string `json:"country"`
	City           string `json:"city"`
	SettlementType string `json:"settlement_type_full"`
	Settlement     string `json:"settlement"`
	StreetType     string `json:"street_type_full"`
	Street         string `json:"street"`
	House          string `json:"house"`
	BlockType      string `json:"block_type"`
	Block          string `json:"block"`
}

type ResponseData struct {
	Value string  `json:"value"`
	Data  Address `json:"data"`
}

type Response struct {
	Suggestions []ResponseData `json:"suggestions"`
}

const token = "39ca673810e9bcd36b78b19a562a3941f3d00883"

func EnrichAddress(ctx context.Context, client *mongo.Client, city string) {
	recs, _ := getEmptyStreetRecords(ctx, client, city)
	//recs, _ := getLatinCity(ctx, client)

	for _, r := range recs {
		url := fmt.Sprintf(
			"https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address?lat=%f&lon=%f",
			r.LocationLatitude, r.LocationLongitude)

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))

		res, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(res.Body)

		enrichAddressFromRequest(r, body)
		res.Body.Close()
	}

	updateMany(ctx, client, recs)
}

func getEmptyStreetRecords(ctx context.Context, client *mongo.Client, city string) ([]*utils.Record, error) {
	collection := getCollection(client)

	filter := bson.D{
		{"address_street",
			bson.D{
				{"$eq", ""},
			},
		},
	}

	if city != "" {
		filter = and(
			filter,
			bson.D{{"address_city", bson.D{{"$eq", city}}}})
	}

	cursor, err := collection.Find(ctx, filter)
	utils.CheckError(err, "get cursor")

	return readAllRecords(ctx, cursor)
}

func getLatinCity(ctx context.Context, client *mongo.Client) ([]*utils.Record, error) {
	collection := getCollection(client)

	filter := bson.D{}
	city, err := collection.Distinct(ctx, "address_city", filter)
	utils.CheckError(err, "distinct get cursor")

	latinCity := make([]string, 0, 10)
	re := regexp.MustCompile("[А-Яа-я0-9-]+?")
	for _, c := range city {
		if isRussian := re.MatchString(c.(string)); !isRussian {
			latinCity = append(latinCity, c.(string))
		}
	}

	var allRecs []*utils.Record

	for _, c := range latinCity {
		filter := bson.D{
			{"address_city",
				bson.D{
					{"$eq", c},
				},
			},
		}
		cursor, err := collection.Find(ctx, filter)
		utils.CheckError(err, "get cursor")

		recs, _ := readAllRecords(ctx, cursor)
		allRecs = append(allRecs, recs...)
	}

	return allRecs, nil
}

func enrichAddressFromRequest(rec *utils.Record, body []byte) {
	var response Response
	json.Unmarshal(body, &response)

	if len(response.Suggestions) == 0 {
		return
	}

	respAddr := response.Suggestions[0].Data

	if rec.AddressCity == "" && respAddr.City != "" {
		//if respAddr.City != "" {
		rec.AddressCity = strings.ToLower(respAddr.City)
	}

	if rec.AddressStreet == "" {
		if respAddr.Street != "" {
			rec.AddressStreet = strings.ToLower(respAddr.StreetType + " " + respAddr.Street)
		} else if respAddr.Settlement != " " {
			rec.AddressStreet = strings.ToLower(respAddr.SettlementType + " " + respAddr.Settlement)
		}
	}

	if rec.AddressHouse == "" && respAddr.House != "" {
		house := strings.ToLower(respAddr.House)
		if respAddr.Block != "" {
			house = house + respAddr.BlockType + respAddr.Block
		}
		rec.AddressHouse = house
	}
}
