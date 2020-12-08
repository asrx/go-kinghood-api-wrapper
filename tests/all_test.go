package tests

import (
	"encoding/json"
	"fmt"
	"github.com/asrx/go-kinghood-api-wrapper/src"
	"github.com/asrx/go-kinghood-api-wrapper/src/Entry"
	"log"
	"testing"
)

var service *src.KinghoodService

const KEY = ""

func init() {
	service = src.NewKinghoodService(KEY)
}

func Test_balance(t *testing.T) {

	money,err := service.GetBalance()
	if err != nil {
		t.Fatal(err)
	}
	log.Println("Balance:", money)
}

func Test_carriers(t *testing.T) {
	content, err := service.GetCarriers()
	if err != nil {
		t.Fatal(err)
	}

	log.Println(content)
}

func Test_rate(t *testing.T) {
	shipment := Entry.Shipment{
		Carrier:             "UPS Ground 6",
		ShipFrom:            GetShipFrom(),
		ShipTo:              GetShipTo(),
		Parcels:             GetParcels(),
		//FreightCalss:        nil,
		//Signature:           nil,
		//WeightUnitType:      nil,
		CarrierServiceLevel: "UPS Ground",
		//LabelSize:           nil,
		//Remark1:             nil,
		//Remark2:             nil,
		//EcOrder:             nil,
	}
	params,err := json.Marshal(shipment)
	if err != nil {
		log.Fatal("Shipment Marshal Json Error:", err)
	}
	//fmt.Println(string(params))
	rateReply, err := service.GetRate(string(params))
	if err != nil {
		log.Fatal("Rate Error:",err)
	}
	fmt.Println(rateReply)
}


func Test_ship(t *testing.T) {
	shipment := Entry.Shipment{
		Carrier:             "UPS Ground 6",
		ShipFrom:            GetShipFrom(),
		ShipTo:              GetShipTo(),
		Parcels:             GetParcels(),
		CarrierServiceLevel: "UPS Ground",
	}
	params,err := json.Marshal(shipment)
	if err != nil {
		log.Fatal("Shipment Marshal Json Error:", err)
	}
	res, _ := service.GetShip(string(params))

	fmt.Printf("%+v\n", res)
}

func Test_Cancel(t *testing.T) {
	orderNumb := "EX145735965"
	res := service.GetCancel(orderNumb)
	log.Println("Cancel is:",res)
}