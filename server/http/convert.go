package http

import "yandex-food/utils"

func recordToHTTPPerson(rec *utils.Record) *utils.HHTPerson {
	return &utils.HHTPerson{
		FullName: rec.FullName,
		Phone:    rec.PhoneNumber,
		Email:    rec.Email,
		Address:  recordToHTTPAddress(rec),
	}
}
func recordToHTTPAddress(rec *utils.Record) *utils.HTTPAddress {
	return &utils.HTTPAddress{
		ID:       rec.ID,
		City:     rec.AddressCity,
		Street:   rec.AddressStreet,
		House:    rec.AddressHouse,
		Entrance: rec.AddressEntrance,
		Floor:    rec.AddressFloor,
		Office:   rec.AddressOffice,
		DoorCode: rec.AddressDoorCode,
	}
}

func recordToHTTPFullRecord(rec *utils.Record) *utils.HTTPFullRecord {
	return &utils.HTTPFullRecord{
		Person:         recordToHTTPPerson(rec),
		Address:        recordToHTTPAddress(rec),
		AddressComment: rec.AddressComment,
		AmountCharged:  rec.AmountCharge,
		CreatedAt:      rec.CreatedAt,
	}
}
