package db

import (
	"gitlab.stageoffice.ru/UCS-COMMON/utils/uuid"
	"reflect"
	"yandex-food/utils"
)

func toBsonData(data []*utils.Record) []*utils.DBRecord {
	var bsonData []*utils.DBRecord

	for _, d := range data {
		var bsonUID uuid.UUID
		bsonUID, err := uuid.FromString(d.ID)
		if err != nil {
			bsonUID = uuid.NewV4()
		}

		bsonData = append(bsonData, &utils.DBRecord{
			ID:                &bsonUID,
			FirstName:         &d.FirstName,
			FullName:          &d.FullName,
			Email:             &d.Email,
			PhoneNumber:       &d.PhoneNumber,
			AddressCity:       &d.AddressCity,
			AddressStreet:     &d.AddressStreet,
			AddressHouse:      &d.AddressHouse,
			AddressEntrance:   &d.AddressEntrance,
			AddressFloor:      &d.AddressFloor,
			AddressOffice:     &d.AddressOffice,
			AddressComment:    &d.AddressComment,
			LocationLatitude:  &d.LocationLatitude,
			LocationLongitude: &d.LocationLongitude,
			AmountCharge:      &d.AmountCharge,
			UserID:            &d.UserID,
			UserAgent:         &d.UserAgent,
			CreatedAt:         &d.CreatedAt,
			AddressDoorCode:   &d.AddressDoorCode,
		})
	}

	return bsonData
}

func fromBsonData(bsonData []*utils.DBRecord) []*utils.Record {
	var data []*utils.Record

	for _, d := range bsonData {
		var err error
		strUID := uuid.Must(*d.ID, err)
		utils.CheckError(err, "get uuid from binary")

		data = append(data, &utils.Record{
			ID:                strUID.String(),
			FirstName:         *d.FirstName,
			FullName:          *d.FullName,
			Email:             *d.Email,
			PhoneNumber:       *d.PhoneNumber,
			AddressCity:       *d.AddressCity,
			AddressStreet:     *d.AddressStreet,
			AddressHouse:      *d.AddressHouse,
			AddressEntrance:   *d.AddressEntrance,
			AddressFloor:      *d.AddressFloor,
			AddressOffice:     *d.AddressOffice,
			AddressComment:    *d.AddressComment,
			LocationLatitude:  *d.LocationLatitude,
			LocationLongitude: *d.LocationLongitude,
			AmountCharge:      *d.AmountCharge,
			UserID:            *d.UserID,
			UserAgent:         *d.UserAgent,
			CreatedAt:         *d.CreatedAt,
			AddressDoorCode:   *d.AddressDoorCode,
		})
	}

	return data
}

func prepareSlice(slice interface{}) []interface{} {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		panic("not a slice")
	}
	if val.IsNil() {
		return nil
	}
	result := make([]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		result[i] = val.Index(i).Interface()
	}
	return result
}
