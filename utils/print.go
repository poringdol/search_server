package utils

import (
	"fmt"
	"io"
)

func PrintRecord(w io.Writer, rec *Record) {
	fmt.Fprintf(
		w, "%s %s %s %d %s %s %s %s %s %s %s %d %s %s\n",
		rec.FirstName,
		rec.FullName,
		rec.Email,
		rec.PhoneNumber,
		rec.AddressCity,
		rec.AddressStreet,
		rec.AddressHouse,
		rec.AddressEntrance,
		rec.AddressFloor,
		rec.AddressOffice,
		rec.AddressComment,
		rec.AmountCharge,
		rec.CreatedAt,
		rec.AddressDoorCode,
	)
}
