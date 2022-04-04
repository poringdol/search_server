package utils

import "gitlab.stageoffice.ru/UCS-COMMON/utils/uuid"

type Record struct {
	ID                string  `bson:"id" csv:"id"`
	FirstName         string  `bson:"first_name" csv:"first_name"`
	FullName          string  `bson:"full_name" csv:"full_name"`
	Email             string  `bson:"email" csv:"email"`
	PhoneNumber       int64   `bson:"phone_number" csv:"phone_number"`
	AddressCity       string  `bson:"address_city" csv:"address_city"`
	AddressStreet     string  `bson:"address_street" csv:"address_street"`
	AddressHouse      string  `bson:"address_house" csv:"address_house"`
	AddressEntrance   string  `bson:"address_entrance" csv:"address_entrance"`
	AddressFloor      string  `bson:"address_floor" csv:"address_floor"`
	AddressOffice     string  `bson:"address_office" csv:"address_office"`
	AddressComment    string  `bson:"address_comment" csv:"address_comment"`
	LocationLatitude  float64 `bson:"location_latitude" csv:"location_latitude"`
	LocationLongitude float64 `bson:"location_longitude" csv:"location_longitude"`
	AmountCharge      int64   `bson:"amount_charged" csv:"amount_charged"`
	UserID            string  `bson:"user_id" csv:"user_id"`
	UserAgent         string  `bson:"user_agent" csv:"user_agent"`
	CreatedAt         string  `bson:"created_at" csv:"created_at"`
	AddressDoorCode   string  `bson:"address_doorcode" csv:"address_doorcode"`
}

type DBRecord struct {
	ID                *uuid.UUID `bson:"_id,omitempty"`
	FirstName         *string    `bson:"first_name,omitempty"`
	FullName          *string    `bson:"full_name,omitempty"`
	Email             *string    `bson:"email,omitempty"`
	PhoneNumber       *int64     `bson:"phone_number,omitempty"`
	AddressCity       *string    `bson:"address_city,omitempty"`
	AddressStreet     *string    `bson:"address_street,omitempty"`
	AddressHouse      *string    `bson:"address_house,omitempty"`
	AddressEntrance   *string    `bson:"address_entrance,omitempty"`
	AddressFloor      *string    `bson:"address_floor,omitempty"`
	AddressOffice     *string    `bson:"address_office,omitempty"`
	AddressComment    *string    `bson:"address_comment,omitempty"`
	LocationLatitude  *float64   `bson:"location_latitude,omitempty"`
	LocationLongitude *float64   `bson:"location_longitude,omitempty"`
	AmountCharge      *int64     `bson:"amount_charged,omitempty"`
	UserID            *string    `bson:"user_id,omitempty"`
	UserAgent         *string    `bson:"user_agent,omitempty"`
	CreatedAt         *string    `bson:"created_at,omitempty"`
	AddressDoorCode   *string    `bson:"address_doorcode,omitempty"`
}

type FullName struct {
	Name       string
	Surname    string
	Patronymic string
}

type HHTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HTTPResponse struct {
	Error *HHTPError `json:"error"`
}

type HTTPGetByPhoneResponse struct {
	Error     *HHTPError     `json:"error"`
	Person    *HHTPerson     `json:"person"`
	Addresses []*HTTPAddress `json:"addresses"`
}

type HTTPGetByAddressResponse struct {
	Error     *HHTPError        `json:"error"`
	Addresses []*HTTPFullRecord `json:"addresses"`
}

type HTTPGetByNameResponse struct {
	Error   *HHTPError   `json:"error"`
	Persons []*HHTPerson `json:"persons"`
}

type HTTPGetCitiesResponse struct {
	Error  *HHTPError `json:"error"`
	Cities []string   `json:"cities"`
}

type HHTPerson struct {
	FullName string       `json:"full_name"`
	Phone    int64        `json:"phone"`
	Email    string       `json:"email"`
	Address  *HTTPAddress `json:"address"`
}

type HTTPAddress struct {
	ID          string            `json:"id"`
	City        string            `json:"city"`
	Street      string            `json:"street"`
	House       string            `json:"house"`
	Entrance    string            `json:"entrance"`
	Floor       string            `json:"floor"`
	Office      string            `json:"office"`
	DoorCode    string            `json:"door_code"`
	FullRecords []*HTTPFullRecord `json:"full_records"`
}

type HTTPLocation struct {
	Latitude  float64 `json:"location_latitude"`
	Longitude float64 `json:"location_longitude"`
}

type HTTPFullRecord struct {
	Person         *HHTPerson   `json:"person"`
	Address        *HTTPAddress `json:"address"`
	AddressComment string       `json:"address_comment"`
	AmountCharged  int64        `json:"amount_charged"`
	CreatedAt      string       `json:"created_at"`
}
