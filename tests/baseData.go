package tests

import . "github.com/asrx/go-kinghood-api-wrapper/src/Entry"

func GetShipFrom() *Address {
	return &Address{
		PostalCode:  "75243",
		Name:        "Karpal Inc.",
		Company:     "Karpal Inc.",
		Address1:    "11516 Pagemill RD",
		State:       "TX",
		City:        "Dallas",
		Country:     "US",
		PhoneNumber: "682-360-2931",
		AddrType:    1,
	}
}

func GetShipTo() *Address {
	return &Address{
		PostalCode:  "92551",
		Name:        "Donovan",
		Address1:    "24300 Nandina Ave",
		State:       "CA",
		City:        "Moreno Valley",
		Country:     "US",
		PhoneNumber: "1347-448-3190",
		AddrType:    1,
	}
}

func GetParcels() []*Parcels {
	var pkgs = make([]*Parcels,0)
	pkgs = append(pkgs, &Parcels{
		Weight: 7.83,
		Length: 39.37,
		Width:  14.57,
		Height: 9.45,
	})
	return pkgs
}
